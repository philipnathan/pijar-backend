package session

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/session/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/session/dto"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
	"github.com/philipnathan/pijar-backend/middleware"
)

type SessionHandler struct {
	service service.SessionService
}

func NewSessionHandler(service service.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

// @Summary		Get sessions for a user
// @Description	Get all sessions for a specific user by user ID
// @Tags			Mentor
// @Accept			json
// @Produce		json
// @Param			user_id	path		int	true	"User ID"
// @Success		200		{array}		MentorSessionResponse
// @Failure		400		{object}	Error	"Invalid user ID"
// @Failure		500		{object}	Error	"Internal server error"
// @Router			/sessions/{user_id} [get]
func (h *SessionHandler) GetSessions(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "Invalid user ID"})
		return
	}

	sessions, err := h.service.GetSessions(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	var response []dto.MentorSessionResponse
	var mentorImageURL string
	var registered bool = false

	for _, session := range sessions {
		if session.User.ImageURL != nil {
			mentorImageURL = *session.User.ImageURL
		} else {
			mentorImageURL = ""
		}

		response = append(response, dto.MentorSessionResponse{
			MentorSessionTitle: session.Title,
			ShortDescription:   session.ShortDescription,
			Schedule:           session.Schedule,
			Registered:         registered,
			MentorDetails: dto.MentorDetails{
				Id:       session.User.ID,
				Fullname: session.User.Fullname,
				ImageURL: mentorImageURL,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

// @Summary		Get learner history session
// @Description	Get learner history session
// @Tags			Learner
// @Produce		json
// @Security		Bearer
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

// @Summary		Get upcoming sessions
// @Schemes
// @Description	Get upcoming sessions
// @Tags			Session
// @Produce		json
// @Param			page		query		int	false	"page"
// @Param			pagesize	query		int	false	"pagesize"
// @Param			categoryid	query		int	false	"categoryid"
// @Success		200	{object}		MentorSessionResponse
// @Failure		400	{object}	Error	"Invalid user ID"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/sessions/upcoming [get]
func (h *SessionHandler) GetUpcommingSessionsLandingPage(c *gin.Context) {
	// check if user is authenticated
	var isAuthenticated bool
	authHeader := c.GetHeader("Authorization")
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

	fmt.Println(categoryIDint)

	// if category_id is available (categoryID is the priority because it is an input from guest/user)
	if categoryIDint > 0 {
		var categoryIDsuint []uint
		categoryIDsuint = append(categoryIDsuint, uint(categoryIDint))
		sessions, total, err := h.service.GetUpcommingSessionsByCategory(categoryIDsuint, page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}

		var sessionDetails []dto.SessionDetail
		for _, session := range *sessions {
			sessionDetails = append(sessionDetails, dto.SessionDetail{
				Day:              session.Schedule.Weekday().String(),
				Time:             session.Schedule.Format("03:04 PM"),
				Title:            session.Title,
				ShortDescription: session.ShortDescription,
				Schedule:         session.Schedule.Format(time.RFC3339),
				ImageURL:         session.ImageURL,
				Registered:       true,
				Duration:         session.EstimateDuration,
			})
		}
		c.JSON(http.StatusOK, dto.GetUpcomingSessionResponse{Sessions: sessionDetails, Page: page, PageSize: pageSize, Total: total})
		return
	}

	// if user is not authenticated (without JWT) and category_id is not available
	if !isAuthenticated && categoryIDint == 0 {
		sessions, total, err := h.service.GetUpcomingSessions(page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}

		var sessionDetails []dto.SessionDetail
		for _, session := range sessions {
			sessionDetails = append(sessionDetails, dto.SessionDetail{
				Day:              session.Schedule.Weekday().String(),
				Time:             session.Schedule.Format("03:04 PM"),
				Title:            session.Title,
				ShortDescription: session.ShortDescription,
				Schedule:         session.Schedule.Format(time.RFC3339),
				ImageURL:         session.ImageURL,
				Registered:       true,
				Duration:         session.EstimateDuration,
			})
		}
		c.JSON(http.StatusOK, dto.GetUpcomingSessionResponse{Sessions: sessionDetails, Page: page, PageSize: pageSize, Total: total})
		return
	}

	// if user is authenticated (with JWT) - get user_id from JWT and get mentor by user interests
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

	sessions, total, err := h.service.GetSessionByLearnerInterests(uint(id), page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	var sessionDetails []dto.SessionDetail

	for _, session := range *sessions {
		sessionDetails = append(sessionDetails, dto.SessionDetail{
			Day:              session.Schedule.Weekday().String(),
			Time:             session.Schedule.Format("03:04 PM"),
			Title:            session.Title,
			ShortDescription: session.ShortDescription,
			Schedule:         session.Schedule.Format(time.RFC3339),
			ImageURL:         session.ImageURL,
			Registered:       true,
			Duration:         session.EstimateDuration,
		})
	}

	c.JSON(http.StatusOK, dto.GetUpcomingSessionResponse{Sessions: sessionDetails, Page: page, PageSize: pageSize, Total: total})
}
