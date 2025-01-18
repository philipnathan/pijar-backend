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
		case custom_error.ErrAlreadyLearner:
			c.JSON(http.StatusConflict, custom_error.ErrAlreadyLearner)
			return
		case custom_error.ErrAlreadyMentor:
			c.JSON(http.StatusConflict, custom_error.ErrAlreadyMentor)
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.NewCustomError(err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"access_token": access_token, "refresh_token": refresh_token})
}

func (h *GoogleAuthHandler) GoogleLoginCallback(c *gin.Context) {
	entity := "login-" + c.Param("entity")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", entity))

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		switch err {
		case custom_error.ErrUserNotFound:
			c.JSON(http.StatusNotFound, custom_error.ErrUserNotFound)
			return
		case custom_error.ErrNotUsingGoogle:
			c.JSON(http.StatusBadRequest, custom_error.ErrNotUsingGoogle)
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.NewCustomError(err.Error()))
			return
		}
	}

	access_token, refresh_token, err := h.service.GoogleLogin(&user.Email, &entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": access_token, "refresh_token": refresh_token})
}

func (h *GoogleAuthHandler) GoogleRegister(c *gin.Context) {
	entity := c.Param("entity")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", entity))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (h *GoogleAuthHandler) GoogleLogin(c *gin.Context) {
	entity := "login-" + c.Param("entity")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", entity))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}
