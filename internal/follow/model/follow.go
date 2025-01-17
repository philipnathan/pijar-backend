package follow

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model  `json:"-"`
	FollowerID  uint `gorm:"not null" json:"follower_id"`
	FollowingID uint `gorm:"not null" json:"following_id"`
}
