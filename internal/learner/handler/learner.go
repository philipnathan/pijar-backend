package learner

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/learner/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/learner/dto"
	service "github.com/philipnathan/pijar-backend/internal/learner/service"
)

type LearnerHandler struct {
	service service.LearnerServiceInterface
}

func NewLearnerHandler(service service.LearnerServiceInterface) *LearnerHandler {
	return &LearnerHandler{
		service: service,
	}
}

func (h *LearnerHandler) GetLearnerInterests(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	learnerInterests, err := h.service.GetLearnerInterests(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	var response []dto.AddLearnerInterestsResponseDto

	for _, interest := range learnerInterests {
		response = append(response, dto.AddLearnerInterestsResponseDto{
			CategoryID:   interest.Category.ID,
			CategoryName: interest.Category.Category_name,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *LearnerHandler) AddLearnerInterests(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	var input dto.AddLearnerInterestsDto
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
		return
	}

	if err := h.service.AddLearnerInterests(uint(id), input.CategoryID); err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
