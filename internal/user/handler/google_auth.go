package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/user/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/user/dto"
	service "github.com/philipnathan/pijar-backend/internal/user/service"
	"github.com/philipnathan/pijar-backend/utils"
	"golang.org/x/oauth2"
)

type GoogleAuthHandler struct {
	service     service.GoogleAuthServiceInterface
	oauthConfig *oauth2.Config
}

func NewGoogleAuthHandler(service service.GoogleAuthServiceInterface, oauthConfig *oauth2.Config) *GoogleAuthHandler {
	return &GoogleAuthHandler{
		service:     service,
		oauthConfig: oauthConfig,
	}
}

type googleInfo struct {
	Family_name string
	Name        string
	Picture     string
	Email       string
	Given_name  string
	ID          string
	Verified    bool
}

//	@Summary		Register using Google
//	@Description	Register using Google. Need authorization code from google
//	@Scheme
//	@Tags		OAuth
//	@Param		entity	path		string	true	"learner/mentor"
//	@Param		code	query		string	true	"authorization code from Google"
//	@Success	200		{object}	RegisterUserResponseDto
//	@Failure	400		{object}	CustomError
//	@Failure	500		{object}	CustomError
//	@Router		/auth/google/{entity}/register [get]
func (h *GoogleAuthHandler) GoogleRegisterCallback(c *gin.Context) {
	entity := c.Param("entity")
	code := c.DefaultQuery("code", "")

	if code == "" {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("code not found"))
	}

	email, name, err := h.authenticateAndGetUserDetails(c, code)
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

//	@Summary		Login using Google
//	@Description	Login using Google. Need authorization code from google
//	@Scheme
//	@Tags		OAuth
//	@Param		entity	path		string	true	"learner/mentor"
//	@Param		code	query		string	true	"authorization code from Google"
//	@Success	200		{object}	LoginUserResponseDto
//	@Failure	400		{object}	CustomError
//	@Failure	500		{object}	CustomError
//	@Router		/auth/google/{entity}/login [get]
func (h *GoogleAuthHandler) GoogleLoginCallback(c *gin.Context) {
	entity := c.Param("entity")
	code := c.DefaultQuery("code", "")

	if code == "" {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("code not found"))
	}

	email, _, err := h.authenticateAndGetUserDetails(c, code)
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

func (h *GoogleAuthHandler) googleExchange(c *gin.Context, code string) (*oauth2.Token, error) {
	fmt.Println(code)

	token, err := h.oauthConfig.Exchange(c.Request.Context(), code)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return token, nil
}

func (h *GoogleAuthHandler) getUserInfo(token *oauth2.Token) (*string, *string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var userInfo googleInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, nil, err
	}

	return &userInfo.Email, &userInfo.Name, nil
}

func (h *GoogleAuthHandler) authenticateAndGetUserDetails(c *gin.Context, code string) (*string, *string, error) {
	token, err := h.googleExchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("Failed to exchange token"))
		return nil, nil, err
	}

	email, name, err := h.getUserInfo(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("Failed to get user info"))
		return nil, nil, err
	}

	return email, name, nil
}
