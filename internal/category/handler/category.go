package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/philipnathan/pijar-backend/internal/category/service"
)

type CategoryHandler struct {
	service service.CategoryServiceInterface
}

func NewCategoryHandler(service service.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) GetAllCategoriesHandler(c *gin.Context) {
	categories, err := h.service.GetAllCategoriesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}
