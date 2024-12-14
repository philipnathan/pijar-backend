package session

import (
    "time"
    "gorm.io/gorm"
)

type MentorSession struct {
    ID               uint      `gorm:"primaryKey"`
    UserID           uint      `gorm:"not null" json:"user_id"`
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
    ID              uint      `gorm:"primaryKey"`
    UserID          uint      `gorm:"not null" json:"user_id"`
    MentorSessionID uint      `gorm:"not null" json:"mentor_session_id"`
    Status          string    `gorm:"type:enum('registered', 'confirmed', 'cancelled_by_mentor', 'cancelled_by_learner', 'complete');default:'registered';not null" json:"status"`
    Rating          float64   `gorm:"default:0;not null" json:"rating"`
    RegisteredAt    time.Time `gorm:"autoCreateTime" json:"registered_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}