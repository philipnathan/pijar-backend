package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/user/dto"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
	"github.com/philipnathan/pijar-backend/utils"
)

type googleBodyResponse struct {
	FamilyName    string `json:"family_name"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	GivenName     string `json:"given_name"`
	ID            string `json:"id"`
	VerifiedEmail bool   `json:"verified_email"`
}

type GoogleAuthHandler struct {
	service service.GoogleAuthServiceInterface
}

func NewGoogleAuthHandler(service service.GoogleAuthServiceInterface) *GoogleAuthHandler {
	return &GoogleAuthHandler{
		service: service,
	}
}

// @Summary		Register using Google
// @Description	Register using Google. Need authorization code from google
// @Scheme
// @Tags		OAuth
// @Param		entity			query		string	true	"learner/mentor"
// @Param		access_token	query		string	true	"acess_token from Google"
// @Success	200				{object}	RegisterUserResponseDto
// @Failure	400				{object}	CustomError
// @Failure	500				{object}	CustomError
// @Router		/auth/google/register [get]
func (h *GoogleAuthHandler) GoogleRegisterCallback(c *gin.Context) {
	entity := c.DefaultQuery("entity", "")
	access_token := c.DefaultQuery("access_token", "")

	if access_token == "" {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("access_token not found"))
		return
	}

	if entity == "" {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("entity not found"))
		return
	}

	email, name, err := h.getUserInfo(&access_token)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError(err.Error()))
		return
	}

	access_token, refresh_token, err := h.service.GoogleRegister(c, email, name, entity)

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

	utils.SetCookie(c, access_token, refresh_token)

	c.JSON(http.StatusOK, dto.RegisterUserResponseDto{
		Message: "user registered successfully",
	})
}

// @Summary		Login using Google
// @Description	Login using Google. Need authorization code from google
// @Scheme
// @Tags		OAuth
// @Param		entity			query		string	true	"learner/mentor"
// @Param		access_token	query		string	true	"authorization code from Google"
// @Success	200				{object}	LoginUserResponseDto
// @Failure	400				{object}	CustomError
// @Failure	500				{object}	CustomError
// @Router		/auth/google/login [get]
func (h *GoogleAuthHandler) GoogleLoginCallback(c *gin.Context) {
	entity := c.DefaultQuery("entity", "")
	access_token := c.DefaultQuery("access_token", "")

	if access_token == "" {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("access_token not found"))
		return
	}

	if entity == "" {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("entity not found"))
		return
	}

	email, _, err := h.getUserInfo(&access_token)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError(err.Error()))
		return
	}

	access_token, refresh_token, err := h.service.GoogleLogin(email, &entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SetCookie(c, access_token, refresh_token)

	c.JSON(http.StatusOK, dto.LoginUserResponseDto{
		Message: "user logged in successfully",
	})
}

func (h *GoogleAuthHandler) getUserInfo(access_token *string) (*string, *string, error) {
	url := "https://www.googleapis.com/oauth2/v2/userinfo"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *access_token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var userInfo googleBodyResponse
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, nil, err
	}

	if userInfo.Email == "" || userInfo.Name == "" {
		return nil, nil, custom_error.ErrUserNotFound
	}

	return &userInfo.Email, &userInfo.Name, nil

}
