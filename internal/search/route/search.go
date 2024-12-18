package route

import (
	"github.com/gin-gonic/gin"
	handler "github.com/philipnathan/pijar-backend/internal/search/handler"
	repository "github.com/philipnathan/pijar-backend/internal/search/repository"
	service "github.com/philipnathan/pijar-backend/internal/search/service"
	"gorm.io/gorm"
)

func SearchRoute(r *gin.Engine, db *gorm.DB) {
    repo := repository.NewSearchRepository(db)
    srv := service.NewSearchService(repo)
    hnd := handler.NewSearchHandler(srv)

    r.GET("/api/v1/search", hnd.Search)
}