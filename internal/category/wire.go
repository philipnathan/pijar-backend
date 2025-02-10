// go:build wireinject
//go:build wireinject
// +build wireinject

package category

import (
	"github.com/google/wire"
	handler "github.com/philipnathan/pijar-backend/internal/category/handler"
	repo "github.com/philipnathan/pijar-backend/internal/category/repository"
	service "github.com/philipnathan/pijar-backend/internal/category/service"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	repo.NewCategoryRepository,
	service.NewCategoryService,
	handler.NewCategoryHandler,
)

func InitializedCategory(db *gorm.DB) (*handler.CategoryHandler, error) {
	wire.Build(ProviderSet)

	return nil, nil
}
