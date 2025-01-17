package follow

import (
	model "github.com/philipnathan/pijar-backend/internal/follow/model"
	"gorm.io/gorm"
)

type FollowRepositoryInterface interface {
	Follow(followerID, followingID *uint) error
	Unfollow(followerID, followingID *uint) error
	IsFollowing(followerID, followingID *uint) (bool, error)
}

type FollowRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) FollowRepositoryInterface {
	return &FollowRepository{
		db: db,
	}
}

func (f *FollowRepository) Follow(followerID, followingID *uint) error {
	var follow model.Follow
	follow.FollowerID = *followerID
	follow.FollowingID = *followingID
	if err := f.db.Create(&follow).Error; err != nil {
		return err
	}
	return nil
}

func (f *FollowRepository) Unfollow(followerID, followingID *uint) error {
	if err := f.db.Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(&model.Follow{}).Error; err != nil {
		return err
	}
	return nil
}

func (f *FollowRepository) IsFollowing(followerID, followingID *uint) (bool, error) {
	var follow model.Follow
	if err := f.db.Where("follower_id = ? AND following_id = ?", followerID, followingID).First(&follow).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
