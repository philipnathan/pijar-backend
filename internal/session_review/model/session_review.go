package session_review

import (
	"gorm.io/gorm"
)

type SessionReview struct {
	gorm.Model `json:"-"`
	UserID     uint    `gorm:"not null" json:"user_id"`
	SessionID  uint    `gorm:"not null" json:"session_id"`
	Review     *string `gorm:"type:text" json:"review"`
	Rating     uint    `gorm:"type:integer;not null" json:"rating"`
}
