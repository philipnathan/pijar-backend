package category

import "gorm.io/gorm"

type SubCategory struct {
	gorm.Model      `json:"-"`
	CategoryID      uint   `gorm:"not null" json:"category_id"`
	SubCategoryName string `gorm:"type:varchar(50);not null" json:"sub_category_name"`
}
