package utils

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, access_token, refresh_token string) {
	PATH := os.Getenv("PATH")
	DOMAIN := os.Getenv("DOMAIN")
	SECURE, err := strconv.ParseBool(os.Getenv("SECURE"))
	if err != nil {
		SECURE = false
	}

	c.SetCookie("access_token", access_token, 3600, PATH, DOMAIN, SECURE, true)
	c.SetCookie("refresh_token", refresh_token, 3600, PATH, DOMAIN, SECURE, true)
}

func DeleteCookie(c *gin.Context) {
	PATH := os.Getenv("PATH")
	DOMAIN := os.Getenv("DOMAIN")
	SECURE, err := strconv.ParseBool(os.Getenv("SECURE"))
	if err != nil {
		SECURE = false
	}

	c.SetCookie("access_token", "", -1, PATH, DOMAIN, SECURE, true)
	c.SetCookie("refresh_token", "", -1, PATH, DOMAIN, SECURE, true)
}
