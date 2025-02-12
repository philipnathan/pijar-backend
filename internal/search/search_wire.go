//go:build wireinject
// +build wireinject

package search

import (
	"github.com/google/wire"
	handler "github.com/philipnathan/pijar-backend/internal/search/handler"
	repo "github.com/philipnathan/pijar-backend/internal/search/repository"
	service "github.com/philipnathan/pijar-backend/internal/search/service"
	"gorm.io/gorm"
)

var SearchProviderSet = wire.NewSet(
	repo.NewSearchRepository,
	service.NewSearchService,
	handler.NewSearchHandler,
)

func InitializedSearch(db *gorm.DB) (handler.SearchHandlerInterface, error) {
	wire.Build(SearchProviderSet)
	return nil, nil
}
