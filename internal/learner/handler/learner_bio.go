package learner

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/learner/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/learner/dto"
	service "github.com/philipnathan/pijar-backend/internal/learner/service"
)

type LearnerBioHandlerInterface interface {
	CreateLearnerBio(c *gin.Context)
	GetLearnerBio(c *gin.Context)
	UpdateLearnerBio(c *gin.Context)
}

type LearnerBioHandler struct {
	service service.LearnerBioServiceInterface
}

func NewLearnerBioHandler(service service.LearnerBioServiceInterface) LearnerBioHandlerInterface {
	return &LearnerBioHandler{
		service: service,
	}
}

// @Summary	Create learner bio
// @Schemes
// @Description	Create learner bio
// @Tags			Learner Bio
// @Accept			json
// @Produce		json
// @Param			bio	body		CreateLearnerBioDto	true	"bio"
// @Success		201	{object}	CreateLearnerBioResponseDto
// @Failure		400	{object}	Error	"Invalid request body"
// @Failure		401	{object}	Error	"Unauthorized"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/learners/biographies [post]
func (h *LearnerBioHandler) CreateLearnerBio(c *gin.Context) {
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

	var input dto.CreateLearnerBioDto
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
		return
	}

	err := h.service.CreateLearnerBio(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateLearnerBioResponseDto{Message: "bio added successfully"})
}

func (h *LearnerBioHandler) GetLearnerBio(c *gin.Context) {
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

	bio, err := h.service.GetLearnerBio(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.GetLearnerBioResponseDto{Message: "bio fetched successfully", Bio: bio})
}

// @Summary	Update learner bio
// @Schemes
// @Description	Update learner bio
// @Tags			Learner Bio
// @Accept			json
// @Produce		json
// @Param			bio	body		UpdateLearnerBioDto	true	"bio"
// @Success		200	{object}	UpdateLearnerBioResponseDto
// @Failure		400	{object}	Error	"Invalid request body"
// @Failure		401	{object}	Error	"Unauthorized"
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/learners/biographies [put]
func (h *LearnerBioHandler) UpdateLearnerBio(c *gin.Context) {
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

	var input dto.UpdateLearnerBioDto
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: err.Error()})
		return
	}

	err := h.service.UpdateLearnerBio(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.UpdateLearnerBioResponseDto{Message: "bio updated successfully"})

}
