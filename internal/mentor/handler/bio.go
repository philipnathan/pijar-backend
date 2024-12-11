package mentor

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/mentor/custom_error"
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
			c.JSON(http.StatusNotFound, custom_error.Error{Error: "Mentor bio not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bio)
}
