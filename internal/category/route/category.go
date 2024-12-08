package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	handler "github.com/philipnathan/pijar-backend/internal/category/handler"
	repository "github.com/philipnathan/pijar-backend/internal/category/repository"
	service "github.com/philipnathan/pijar-backend/internal/category/service"
)

func CategoryRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/categories"

	repo := repository.NewCategoryRepository(db)
	services := service.NewCategoryService(repo)
	handler := handler.NewCategoryHandler(services)

	categoryRoutes := r.Group(apiV1)
	{
		categoryRoutes.GET("/", handler.GetAllCategoriesHandler)
	}
}