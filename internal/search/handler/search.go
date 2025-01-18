package search

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/search/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/search/dto"
	service "github.com/philipnathan/pijar-backend/internal/search/service"
)

type SearchHandler struct {
	service service.SearchServiceInterface
}

func NewSearchHandler(service service.SearchServiceInterface) *SearchHandler {
	return &SearchHandler{service: service}
}

// @Summary		Search for sessions, mentors, and categories
// @Description	Search for sessions, mentors, and categories by keyword
// @Tags			Search
// @Produce		json
// @Param			keyword		query		string	true	"Search Keyword min 3 characters long"
// @Param			page		query		int		false	"Page number for sessions"
// @Param			pagesize	query		int		false	"Page size for sessions"
// @Success		200			{object}	SearchResponse
// @Failure		400			{object}	Error	"Bad Request"
// @Failure		500			{object}	Error	"Internal server error"
// @Router			/search [get]
func (h *SearchHandler) Search(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, custom_error.Error{Message: "Keyword is required"})
		return
	}

	if len(keyword) < 3 {
		c.JSON(http.StatusBadRequest, custom_error.Error{Message: "Keyword must be at least 3 characters long"})
		return
	}

	sessionPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sessionPageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))

	sessions, mentors, categories, total, err := h.service.Search(&keyword, &sessionPage, &sessionPageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Message: err.Error()})
		return
	}

	var sessionDetails []dto.Session
	var sessionImage, mentorImage string

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

		var average_rating float32 = 0

		if len(session.SessionReviews) > 0 {
			for _, rating := range session.SessionReviews {
				average_rating += float32(rating.Rating)
			}

			average_rating /= float32(len(session.SessionReviews))
		}

		sessionDetails = append(sessionDetails, dto.Session{
			ID: session.ID,
			MentorDetails: dto.MentorDetails{
				Id:       session.User.ID,
				Fullname: session.User.Fullname,
				ImageURL: mentorImage,
			},
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
		})
	}

	var mentorDetails []dto.MentorDetailLandingPage
	for _, mentor := range *mentors {
		if mentor.ImageURL == nil {
			mentorImage = ""
		} else {
			mentorImage = *mentor.ImageURL
		}

		mentorDetails = append(mentorDetails, dto.MentorDetailLandingPage{
			Id:         mentor.ID,
			Fullname:   mentor.Fullname,
			ImageURL:   mentorImage,
			Occupation: mentor.MentorExperiences[0].Occupation,
		})
	}

	var categoryDetails []dto.Category
	for _, category := range *categories {
		categoryDetails = append(categoryDetails, dto.Category{
			ID:           category.ID,
			CategoryName: category.Category_name,
			ImageURL:     category.Image_url,
		})
	}

	var allResponse dto.SearchResponse
	allResponse.Sessions = dto.GetAllSessionsResponse{
		Sessions: sessionDetails,
		Total:    total,
		Page:     sessionPage,
		PageSize: sessionPageSize,
	}
	allResponse.Mentors = mentorDetails
	allResponse.Categories = categoryDetails

	c.JSON(http.StatusOK, allResponse)
}
