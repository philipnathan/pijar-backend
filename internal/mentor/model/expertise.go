package mentor

import (
	category "github.com/philipnathan/pijar-backend/internal/category/model"
	"gorm.io/gorm"
)

type MentorExpertises struct {
	gorm.Model  `json:"-"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Expertise   string `gorm:"type:varchar(50);not null" json:"expertise"`
	Category_id uint   `gorm:"not null" json:"category_id"`

	Category category.Category `gorm:"foreignKey:Category_id;references:ID" json:"category"`
}
