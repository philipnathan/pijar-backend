package seed

import (
	model "github.com/philipnathan/pijar-backend/internal/category/model"
	"gorm.io/gorm"
)

func SeedSubCategory(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	subCategories := []model.SubCategory{
		{
			CategoryID:      1,
			SubCategoryName: "Pertanian Organik",
		},
		{
			CategoryID:      2,
			SubCategoryName: "Kewirausahaan Mikro",
		},
		{
			CategoryID:      3,
			SubCategoryName: "Daur Ulang Plastik",
		},
		{
			CategoryID:      4,
			SubCategoryName: "Olahan Pisang",
		},
	}

	var count int64
	if err := db.Model(&model.SubCategory{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, subCategory := range subCategories {
		if err := db.Create(&subCategory).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
