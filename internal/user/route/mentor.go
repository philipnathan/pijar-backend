package user

import (
	"github.com/gin-gonic/gin"
	initMentor "github.com/philipnathan/pijar-backend/internal/user"
	"gorm.io/gorm"
)

func MentorUserRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/users"

	handler, err := initMentor.InitializedMentor(db)
	if err != nil {
		panic(err)
	}

	nonProtectedRoutes := r.Group(apiV1)
	{
		nonProtectedRoutes.POST("/registermentor", handler.RegisterMentor)
	}
}
