package mentor

import "gorm.io/gorm"

type MentorExpertises struct {
	gorm.Model  `json:"-"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Expertise   string `gorm:"type:varchar(50);not null" json:"expertise"`
	Category_id uint   `gorm:"not null" json:"category_id"`
}
