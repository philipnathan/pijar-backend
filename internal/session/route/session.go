package session

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	handler "github.com/philipnathan/pijar-backend/internal/session/handler"
	repository "github.com/philipnathan/pijar-backend/internal/session/repository"
	service "github.com/philipnathan/pijar-backend/internal/session/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
)

func SessionRoute(router *gin.RouterGroup, db *gorm.DB) {
	apiV1 := "/api/v1/sessions"
	
	sessionRepository := repository.NewSessionRepository(db)
	sessionService := service.NewSessionService(sessionRepository)
	sessionHandler := handler.NewSessionHandler(sessionService)

	router.GET(apiV1+"/upcoming", middleware.AuthMiddleware(), sessionHandler.GetUpcomingSessions) 

}