package customerror

import (

    "github.com/gin-gonic/gin"
)

func CustomErrorHandler(c *gin.Context) {
    c.Next()

    if len(c.Errors) > 0 {
        c.JSON(-1, gin.H{"errors": c.Errors})
        return
    }
}

func RegisterCustomErrorHandler(router *gin.Engine) {
    router.Use(CustomErrorHandler)
}