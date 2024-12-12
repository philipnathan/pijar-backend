package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/category/model"
	repo "github.com/philipnathan/pijar-backend/internal/category/repository"
	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	categories := []model.Category{
		{Category_name: "Development", Image_url: "development.png"},
		{Category_name: "Design", Image_url: "design.png"},
		{Category_name: "Marketing", Image_url: "marketing.png"},
		{Category_name: "Business", Image_url: "business.png"},
		{Category_name: "Photography", Image_url: "photography.png"},
		{Category_name: "Music", Image_url: "music.png"},
		{Category_name: "Finance", Image_url: "finance.png"},
	}

	var count int64
	if err := db.Model(&model.Category{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, category := range categories {
		if err := repo.NewCategoryRepository(db).SaveCategory(&category); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
