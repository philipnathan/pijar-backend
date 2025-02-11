package user

import (
	"github.com/gin-gonic/gin"
	initGoogle "github.com/philipnathan/pijar-backend/internal/user"
	"gorm.io/gorm"
)

func GoogleAuthRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/"

	handler, err := initGoogle.InitializedGoogleAuth(db)
	if err != nil {
		panic(err)
	}

	nonProtected := r.Group(apiV1)
	{
		nonProtected.GET("/auth/google/register", handler.GoogleRegisterCallback)
		nonProtected.GET("/auth/google/login", handler.GoogleLoginCallback)
	}
}
