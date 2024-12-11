package mentor

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/mentor/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/mentor/dto"
	service "github.com/philipnathan/pijar-backend/internal/mentor/service"
)

type MentorBioHandler struct {
	service service.MentorBioServiceInterface
}

func NewMentorBioHandler(service service.MentorBioServiceInterface) *MentorBioHandler {
	return &MentorBioHandler{
		service: service,
	}
}

func (h *MentorBioHandler) MentorGetBio(c *gin.Context) {
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	idFloat, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	id := uint(idFloat)
	bio, err := h.service.GetMentorBio(&id)
	if err != nil {
		if err == custom_error.ErrMentorBioNotFound {
			c.JSON(http.StatusOK, dto.GetMentorBioResponseDto{Message: "bio not found", Bio: ""})
			return
		}
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.GetMentorBioResponseDto{Message: "bio fetched successfully", Bio: bio.Bio})
}

func (h *MentorBioHandler) UserGetBio(c *gin.Context) {
	MentorIDStr := c.Param("mentor_id")
	if MentorIDStr == "" {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "mentor_id is required"})
		return
	}

	MentorID, err := strconv.ParseUint(MentorIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "mentor_id is invalid"})
		return
	}

	uintMentorID := uint(MentorID)

	mentorBio, err := h.service.GetMentorBio(&uintMentorID)
	if err != nil {
		if err == custom_error.ErrMentorBioNotFound {
			c.JSON(http.StatusOK, dto.GetMentorBioResponseDto{Message: "bio not found", Bio: ""})
			return
		}
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.GetMentorBioResponseDto{Message: "bio fetched successfully", Bio: mentorBio.Bio})
}
