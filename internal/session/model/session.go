package mentor

import (
	session "github.com/philipnathan/pijar-backend/internal/session/model"
	"gorm.io/gorm"
	"time"
)

type MentorSession struct {
    ID               uint      `gorm:"primaryKey"`
    UserID           uint      `gorm:"not null" json:"user_id"`
    CategoryID       uint      `gorm:"not null" json:"category_id"`
    Title            string(144)    `gorm:"size:144;not null" json:"title"`
    ShortDescription string(256)    `gorm:"size:256;not null" json:"short_description"`
    Detail           text    `gorm:"not null" json:"detail"`
    Schedule         time.Time `gorm:"not null" json:"schedule"`
    EstimateDuration int       `gorm:"not null" json:"estimate_duration"`
    ImageURL         text    `gorm:"not null" json:"image_url"`
    Link             text    `gorm:"not null" json:"link"`
    CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type MentorSessionParticipant struct {
    ID             uint   `gorm:"primaryKey"`
    UserID         uint   `gorm:"not null" json:"user_id"`
    MentorSessionID uint  `gorm:"not null" json:"mentor_session_id"`
    Status         enums `gorm:"default:registered;not null" json:"status"`
    Rating         float64 `gorm:"default:0;not null" json:"rating"`
    RegisteredAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"registered_at"`
    UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
