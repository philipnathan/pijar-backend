package session_review

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/session_review/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/session_review/dto"
	service "github.com/philipnathan/pijar-backend/internal/session_review/service"
)

type SessionReviewHandler struct {
	service service.SessionReviewServiceInterface
}

func NewSessionReviewHandler(service service.SessionReviewServiceInterface) *SessionReviewHandler {
	return &SessionReviewHandler{
		service: service,
	}
}

// @Summary		Create session review
// @Description	Create session review
// @Schemes
// @Tags		Session
// @Accept		json
// @Produce	json
// @Security Bearer
// @Param		session_id		path		string					true	"Session ID"
// @Param		session_review	body		SessionReviewRequest	true	"Session Review"
// @Success	200				{object}	SessionReviewResponse
// @Failure	400				{object}	CustomError	"User not found"
// @Failure	500				{object}	CustomError	"Internal server error"
// @Router		/reviews/{session_id} [post]
func (h *SessionReviewHandler) CreateSessionReviewHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.NewCustomError("Unauthorized"))
		return
	}

	userIDFloat, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.NewCustomError("Unauthorized"))
		return
	}

	sessionIDStr := c.Param("session_id")
	if sessionIDStr == "" {
		c.JSON(http.StatusBadRequest, custom_error.ErrSessionNotFound)
		return
	}

	SessionID, err := strconv.ParseUint(sessionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.ErrSessionNotFound)
		return
	}

	var request dto.SessionReviewRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError(err.Error()))
		return
	}

	if request.Rating < 1 || request.Rating > 5 {
		c.JSON(http.StatusBadRequest, custom_error.ErrInvalidRating)
		return
	}

	if request.Review != nil {
		if len(*request.Review) > 250 {
			c.JSON(http.StatusBadRequest, custom_error.ErrReviewTooLong)
			return
		}
	}

	userIDuint := uint(userIDFloat)
	sessionIDuint := uint(SessionID)

	err = h.service.CreateSessionReview(&userIDuint, &sessionIDuint, &request.Rating, request.Review)

	if err != nil {
		switch err {
		case custom_error.ErrUserNotFound:
			c.JSON(http.StatusBadRequest, custom_error.ErrUserNotFound)
			return
		case custom_error.ErrSessionNotFound:
			c.JSON(http.StatusBadRequest, custom_error.ErrSessionNotFound)
			return
		case custom_error.ErrUserAlreadyReviewed:
			c.JSON(http.StatusBadRequest, custom_error.ErrUserAlreadyReviewed)
			return
		case custom_error.ErrLearnerNotEnrolled:
			c.JSON(http.StatusBadRequest, custom_error.ErrLearnerNotEnrolled)
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.NewCustomError(err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, dto.SessionReviewResponse{Message: "Review successfully created"})
}