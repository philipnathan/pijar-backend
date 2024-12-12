package notification

import "gorm.io/gorm"

type NotificationType struct {
	gorm.Model `json:"-"`
	Type       string `gorm:"type:varchar(50);not null" json:"type"`
}
