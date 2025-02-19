package session

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/session/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/session/dto"
	model "github.com/philipnathan/pijar-backend/internal/session/model"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
	"github.com/philipnathan/pijar-backend/middleware"
)

type SessionHandlerInterface interface {
	GetLearnerHistorySession(c *gin.Context)
	GetUpcommingSessionsLandingPage(c *gin.Context)
	GetAllSessionsWithFilter(c *gin.Context)
	GetSessionDetailById(c *gin.Context)
}

type SessionHandler struct {
	service service.SessionService
}

func NewSessionHandler(service service.SessionService) SessionHandlerInterface {
	return &SessionHandler{service: service}
}

// @Summary		Get learner history session
// @Description	Get learner history session
// @Tags			Learner
// @Produce		json
// @Success		200	{object}	GetUserHistorySessionResponseDto
// @Failure		401	{object}	Error	"Unauthorized"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/sessions/histories [get]
func (h *SessionHandler) GetLearnerHistorySession(c *gin.Context) {
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	histories, err := h.service.GetLearnerHistorySession(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	var response dto.GetUserHistorySessionResponseDto
	var mentorImageURL string

	for _, history := range *histories {
		if history.MentorSession.Schedule.After(time.Now()) {
			continue
		}
		if history.MentorSession.User.ImageURL != nil {
			mentorImageURL = *history.MentorSession.User.ImageURL
		} else {
			mentorImageURL = ""
		}

		response.Histories = append(response.Histories, dto.History{
			MentorSessionTitle: history.MentorSession.Title,
			ShortDescription:   history.MentorSession.ShortDescription,
			Schedule:           history.MentorSession.Schedule,
			Status:             string(history.Status),
			MentorDetails: dto.MentorDetails{
				Id:       history.MentorSession.User.ID,
				Fullname: history.MentorSession.User.Fullname,
				ImageURL: mentorImageURL,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

// @Summary	Get upcoming sessions
// @Schemes
// @Description	Get upcoming sessions
// @Tags			Session
// @Produce		json
// @Param			page		query		int	false	"page"
// @Param			pagesize	query		int	false	"pagesize"
// @Param			categoryid	query		int	false	"categoryid"
// @Success		200			{object}	GetAllSessionsResponse
// @Failure		400			{object}	Error	"Invalid user ID"
// @Failure		500			{object}	Error	"Internal server error"
// @Router			/sessions/upcoming [get]
func (h *SessionHandler) GetUpcommingSessionsLandingPage(c *gin.Context) {
	// check if user is authenticated
	var isAuthenticated bool
	authHeader, err := c.Cookie("access_token")
	if err != nil {
		isAuthenticated = false
	}
	if authHeader != "" {
		middleware.AuthMiddleware()(c)

		if !c.IsAborted() {
			isAuthenticated = true
		}
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	categoryIDint, _ := strconv.Atoi(c.DefaultQuery("categoryid", "0"))

	// if category_id is negative (invalid)
	if categoryIDint < 0 {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "category_id is invalid"})
		return
	}

	var sessions *[]model.MentorSession
	var total int

	// if category_id is available (categoryID is the priority because it is an input from guest/user)
	if categoryIDint > 0 {
		var categoryIDsuint []uint
		categoryIDsuint = append(categoryIDsuint, uint(categoryIDint))
		sessions, total, err = h.service.GetUpcommingSessionsByCategory(categoryIDsuint, page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}
	}

	// if user is not authenticated (without JWT) and category_id is not available
	if !isAuthenticated && categoryIDint == 0 {
		sessions, total, err = h.service.GetUpcomingSessions(page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}
	}

	// if user is authenticated (with JWT) - get user_id from JWT and get mentor by user interests
	if isAuthenticated && categoryIDint == 0 {
		UserID, exist := c.Get("user_id")
		if !exist {
			c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
			return
		}

		id, ok := UserID.(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
			return
		}

		sessions, total, err = h.service.GetSessionByLearnerInterests(uint(id), page, pageSize)

		if err != nil {
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}
	}

	var sessionDetails []dto.Session
	var sessionImage string
	var mentorImage string

	for _, session := range *sessions {
		if session.ImageURL == "" {
			sessionImage = ""
		} else {
			sessionImage = session.ImageURL
		}

		if session.User.ImageURL == nil {
			mentorImage = ""
		} else {
			mentorImage = *session.User.ImageURL
		}

		sessionDetails = append(sessionDetails, dto.Session{
			ID:               session.ID,
			MentorDetails:    dto.MentorDetails{Id: session.User.ID, Fullname: session.User.Fullname, ImageURL: mentorImage},
			Category:         session.Category.Category_name,
			Title:            session.Title,
			ShortDescription: session.ShortDescription,
			Detail:           session.Detail,
			Schedule:         session.Schedule.Format(time.RFC3339),
			Duration:         session.EstimateDuration,
			ImageURL:         sessionImage,
			Link:             session.Link,
			Day:              session.Schedule.Weekday().String(),
			Time:             session.Schedule.Format("03:04 PM"),
			AverageRating:    0,
		})
	}

	response := dto.GetAllSessionsResponse{
		Sessions: sessionDetails,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary		Get all sessions
// @Description	Get all sessions and can be filtered by categoryid and mentorid
// @Tags			Session
// @Produce		json
// @Param			categoryid	query		int		false	"Category ID"
// @Param			mentorid	query		int		false	"Mentor ID"
// @Param			page		query		int		false	"Page number"
// @Param			pagesize	query		int		false	"Page size"
// @Param			rating		query		string	false	"Rating"	Enums(highest, lowest)
// @Param			schedule	query		string	false	"Schedule"	Enums(newest, oldest)
// @Success		200			{object}	GetAllSessionsResponse
// @Failure		500			{object}	Error
// @Router			/sessions [get]
func (h *SessionHandler) GetAllSessionsWithFilter(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	categoryIDint, _ := strconv.Atoi(c.DefaultQuery("categoryid", "0"))
	mentorIDint, _ := strconv.Atoi(c.DefaultQuery("mentorid", "0"))
	rating := c.DefaultQuery("rating", "highest")
	schedule := c.DefaultQuery("schedule", "newest")

	sessions, total, err := h.service.GetAllSessionsWithFilter(uint(categoryIDint), uint(mentorIDint), page, pageSize, rating, schedule)

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	var response dto.GetAllSessionsResponse
	var sessionsResponse []dto.Session
	var sessionImageURL *string
	var mentorImageURL *string

	for _, session := range *sessions {
		if session.ImageURL != "" {
			sessionImageURL = &session.ImageURL
		} else {
			sessionImageURL = nil
		}
		if session.User.ImageURL != nil {
			mentorImageURL = session.User.ImageURL
		} else {
			mentorImageURL = nil
		}

		var average_rating float32
		if len(session.SessionReviews) == 0 {
			average_rating = 0
		} else {
			for _, review := range session.SessionReviews {
				average_rating += float32(review.Rating)
			}
			average_rating /= float32(len(session.SessionReviews))
		}

		sessionsResponse = append(sessionsResponse, dto.Session{
			ID:               session.ID,
			MentorDetails:    dto.MentorDetails{Id: session.User.ID, Fullname: session.User.Fullname, ImageURL: *mentorImageURL},
			Category:         session.Category.Category_name,
			Title:            session.Title,
			ShortDescription: session.ShortDescription,
			Detail:           session.Detail,
			Schedule:         session.Schedule.Format(time.RFC3339),
			Duration:         session.EstimateDuration,
			ImageURL:         *sessionImageURL,
			Link:             session.Link,
			Day:              session.Schedule.Weekday().String(),
			Time:             session.Schedule.Format("03:04 PM"),
			AverageRating:    average_rating,
		})
	}

	if rating != "" {
		sort.Slice(sessionsResponse, func(i, j int) bool {
			if rating == "highest" {
				return sessionsResponse[i].AverageRating > sessionsResponse[j].AverageRating
			} else {
				return sessionsResponse[i].AverageRating < sessionsResponse[j].AverageRating
			}
		})
	}

	response = dto.GetAllSessionsResponse{
		Sessions: sessionsResponse,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary		Get session detail by id
// @Description	Get session detail by session_id
// @Schemes
// @Tags		Session
// @Produce	json
// @Param		session_id	path		int	true	"Session ID"
// @Success	200			{object}	Session
// @Failure	400			{object}	Error
// @Failure	500			{object}	Error
// @Router		/sessions/{session_id} [get]
func (h *SessionHandler) GetSessionDetailById(c *gin.Context) {
	sessionID, _ := strconv.Atoi(c.Param("session_id"))
	if sessionID <= 0 {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "session_id is invalid"})
		return
	}

	uintSessionID := uint(sessionID)

	session, err := h.service.GetDetailSessionByID(uintSessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	var sessionImage string
	var mentorImageURL string

	if session.ImageURL != "" {
		sessionImage = session.ImageURL
	} else {
		sessionImage = ""
	}

	if session.User.ImageURL != nil {
		mentorImageURL = *session.User.ImageURL
	} else {
		mentorImageURL = ""
	}

	var average_rating float32
	if len(session.SessionReviews) == 0 {
		average_rating = 0
	} else {
		for _, review := range session.SessionReviews {
			average_rating += float32(review.Rating)
		}
		average_rating /= float32(len(session.SessionReviews))
	}

	response := dto.Session{
		ID:               session.ID,
		MentorDetails:    dto.MentorDetails{Id: session.User.ID, Fullname: session.User.Fullname, ImageURL: mentorImageURL},
		Category:         session.Category.Category_name,
		Title:            session.Title,
		ShortDescription: session.ShortDescription,
		Detail:           session.Detail,
		Schedule:         session.Schedule.Format(time.RFC3339),
		Duration:         session.EstimateDuration,
		ImageURL:         sessionImage,
		Link:             session.Link,
		Day:              session.Schedule.Weekday().String(),
		Time:             session.Schedule.Format("03:04 PM"),
		AverageRating:    average_rating,
	}

	c.JSON(http.StatusOK, response)
}
