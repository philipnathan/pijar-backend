package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, access_token, refresh_token string) {
	frontend_url := os.Getenv("FRONTEND_URL")

	c.SetCookie("access_token", access_token, 3600, "/", frontend_url, false, true)
	c.SetCookie("refresh_token", refresh_token, 3600, "/", frontend_url, false, true)
}

func DeleteCookie(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", os.Getenv("FRONTEND_URL"), false, true)
	c.SetCookie("refresh_token", "", -1, "/", os.Getenv("FRONTEND_URL"), false, true)
}
