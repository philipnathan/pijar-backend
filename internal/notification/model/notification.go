package notification

import (
	model "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model         `json:"-"`
	UserID             uint   `gorm:"not null" json:"user_id"`
	NotificationTypeID uint   `gorm:"not null" json:"notification_type_id"`
	Message            string `gorm:"type:text;not null" json:"message"`
	IsRead             bool   `gorm:"type:bool;default:false" json:"is_read"`

	NotificationType NotificationType `gorm:"foreignKey:NotificationTypeID;references:ID" json:"notification_type"`
	User             model.User       `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
