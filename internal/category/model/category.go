package category

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model    `json:"-"`
	ID            uint   `gorm:"primaryKey" json:"id"`
	Category_name string `gorm:"type:varchar(50);not null" json:"category_name"`
	Image_url     string `gorm:"type:text" json:"image_url"`

	SubCategories []SubCategory `gorm:"foreignKey:CategoryID;references:ID" json:"sub_categories"`
}
