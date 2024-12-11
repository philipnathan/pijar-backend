package mentor

import "gorm.io/gorm"

type MentorBiographies struct {
	gorm.Model `json:"-"`
	UserID     uint   `gorm:"not null" json:"user_id"`
	Bio        string `gorm:"type:text;not null" json:"bio"`
}
