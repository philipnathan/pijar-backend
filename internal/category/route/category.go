package category

import (
	"github.com/gin-gonic/gin"
	categoryInit "github.com/philipnathan/pijar-backend/internal/category"
	"gorm.io/gorm"
)

func CategoryRoute(r *gin.Engine, db *gorm.DB) {
	apiV1 := "/api/v1/categories"

	handler, err := categoryInit.InitializedCategory(db)
	if err != nil {
		panic(err)
	}

	categoryRoutes := r.Group(apiV1)
	{
		categoryRoutes.GET("", handler.GetAllCategoriesHandler)
		categoryRoutes.GET("/featured", handler.GetFeaturedCategoriesHandler)
	}
}
