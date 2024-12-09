package learner

import (
	category "github.com/philipnathan/pijar-backend/internal/category/model"
	"gorm.io/gorm"
)

type LearnerBio struct {
	gorm.Model  `json:"-"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Bio         string `gorm:"type:text;not null" json:"bio"`
	Occupation  string `gorm:"type:varchar(50);not null" json:"occupation"`
	Institution string `gorm:"type:varchar(50);not null" json:"institution"`
}

type LearnerInterest struct {
	gorm.Model `json:"-"`
	UserID     uint `gorm:"not null" json:"user_id"`
	CategoryID uint `gorm:"type:varchar(50);not null" json:"category_id"`

	Category category.Category `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
}
