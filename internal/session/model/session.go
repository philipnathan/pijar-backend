package session

import (
    "time"
    "github.com/philipnathan/pijar-backend/internal/user/model"
)

type MentorSession struct {
    ID               uint      `gorm:"primaryKey"`
    UserID           uint      `gorm:"not null;index" json:"user_id"`
    User             user.User `gorm:"foreignKey:UserID;references:ID"`
    CategoryID       uint      `gorm:"not null" json:"category_id"`
    Title            string    `gorm:"size:144;not null" json:"title"`
    ShortDescription string    `gorm:"size:256;not null" json:"short_description"`
    Detail           string    `gorm:"type:text;not null" json:"detail"`
    Schedule         time.Time `gorm:"not null" json:"schedule"`
    EstimateDuration int       `gorm:"not null" json:"estimate_duration"`
    ImageURL         string    `gorm:"type:text;not null" json:"image_url"`
    Link             string    `gorm:"type:text;not null" json:"link"`
    CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MentorSessionParticipant struct {
    ID              uint          `gorm:"primaryKey"`
    MentorSessionID uint          `gorm:"not null;index" json:"mentor_session_id"`
    MentorSession   MentorSession `gorm:"foreignKey:MentorSessionID;references:ID"`
    UserID          uint          `gorm:"not null;index" json:"user_id"`
    User            user.User     `gorm:"foreignKey:UserID;references:ID"`
    Status          string        `gorm:"type:mentor_session_participants_status;default:'registered';not null" json:"status"`
    Rating          float32       `json:"rating"`
    RegisteredAt    time.Time     `gorm:"autoCreateTime" json:"registered_at"`
    UpdatedAt       time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}