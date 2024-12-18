package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/search/custom_error"
	service "github.com/philipnathan/pijar-backend/internal/search/service"
)

type SearchHandler struct {
	service service.SearchServiceInterface
}

func NewSearchHandler(service service.SearchServiceInterface) *SearchHandler {
	return &SearchHandler{service: service}
}

// @Summary Search for sessions, mentors, and categories
// @Description Search for sessions, mentors, and categories by keyword
// @Tags Search
// @Produce json
// @Param keyword query string true "Search Keyword"
// @Success 200 {object} SearchResponse
// @Failure 400 {object} Error "Bad Request"
// @Failure 500 {object} Error "Internal server error"
// @Router /search [get]
func (h *SearchHandler) Search(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, custom_error.Error{Message: "Keyword is required"})
		return
	}

	results, err := h.service.Search(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
