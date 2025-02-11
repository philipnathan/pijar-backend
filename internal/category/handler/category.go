package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/category/custom_error"
	service "github.com/philipnathan/pijar-backend/internal/category/service"
)

type CategoryHandlerInterface interface {
	GetAllCategoriesHandler(c *gin.Context)
	GetFeaturedCategoriesHandler(c *gin.Context)
}

type CategoryHandler struct {
	service service.CategoryServiceInterface
}

func NewCategoryHandler(service service.CategoryServiceInterface) CategoryHandlerInterface {
	return &CategoryHandler{
		service: service,
	}
}

// @Summary	Get all categories
// @Schemes
// @Description	Get all categories
// @Tags			Category
// @Produce		json
// @Success		200	{array}		Category
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/categories [get]
func (h *CategoryHandler) GetAllCategoriesHandler(c *gin.Context) {
	categories, err := h.service.GetAllCategoriesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// @Summary	Get featured categories
// @Schemes
// @Description	Get featured categories
// @Tags			Category
// @Produce		json
// @Success		200	{array}		dto.FeaturedCategoryResponseDto
// @Failure		500	{object}	Error	"Internal server error"
// @Router			/categories/featured [get]
func (h *CategoryHandler) GetFeaturedCategoriesHandler(c *gin.Context) {
	categories, err := h.service.GetFeaturedCategoriesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}
