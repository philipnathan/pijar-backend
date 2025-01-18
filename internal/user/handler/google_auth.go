package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
)

type GoogleAuthHandler struct {
	service service.GoogleAuthServiceInterface
}

func NewGoogleAuthHandler(service service.GoogleAuthServiceInterface) *GoogleAuthHandler {
	return &GoogleAuthHandler{
		service: service,
	}
}

func (h *GoogleAuthHandler) GoogleAuthCallback(c *gin.Context) {
	entity := c.Param("entity")

	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", entity))

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("from handler: ", user)

	access_token, refresh_token, err := h.service.GoogleRegister(c, &user.Email, &user.Name, entity)

	if err != nil {
		switch err {
		case custom_error.ErrUserAlreadyLearner:
			c.JSON(http.StatusConflict, custom_error.ErrUserAlreadyLearner)
			return
		case custom_error.ErrUserAlreadyMentor:
			c.JSON(http.StatusConflict, custom_error.ErrUserAlreadyMentor)
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.NewCustomError(err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"access_token": access_token, "refresh_token": refresh_token})
}

func (h *GoogleAuthHandler) GoogleAuth(c *gin.Context) {
	entity := c.Param("entity")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", entity))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}
