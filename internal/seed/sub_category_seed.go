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
		{CategoryID: 1, SubCategoryName: "Java"},
		{CategoryID: 1, SubCategoryName: "Python"},
		{CategoryID: 1, SubCategoryName: "C++"},
		{CategoryID: 1, SubCategoryName: "C#"},
		{CategoryID: 2, SubCategoryName: "Photoshop"},
		{CategoryID: 2, SubCategoryName: "Illustrator"},
		{CategoryID: 2, SubCategoryName: "Figma"},
		{CategoryID: 2, SubCategoryName: "Sketch"},
		{CategoryID: 3, SubCategoryName: "Facebook"},
		{CategoryID: 3, SubCategoryName: "Instagram"},
		{CategoryID: 3, SubCategoryName: "Twitter"},
		{CategoryID: 3, SubCategoryName: "Shopee"},
		{CategoryID: 3, SubCategoryName: "Tokopedia"},
		{CategoryID: 4, SubCategoryName: "Finance"},
		{CategoryID: 4, SubCategoryName: "Marketing"},
		{CategoryID: 4, SubCategoryName: "Business"},
		{CategoryID: 5, SubCategoryName: "Photography"},
		{CategoryID: 5, SubCategoryName: "Videography"},
		{CategoryID: 6, SubCategoryName: "Music"},
		{CategoryID: 6, SubCategoryName: "Podcast"},
		{CategoryID: 7, SubCategoryName: "Finance"},
		{CategoryID: 7, SubCategoryName: "Marketing"},
		{CategoryID: 7, SubCategoryName: "Business"},
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
