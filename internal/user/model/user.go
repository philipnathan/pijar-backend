package user

import (
	"database/sql/driver"
	"errors"
	"time"

	learner "github.com/philipnathan/pijar-backend/internal/learner/model"
	mentor "github.com/philipnathan/pijar-backend/internal/mentor/model"
	"gorm.io/gorm"
)

type CustomTime struct {
	time.Time
}

type User struct {
	gorm.Model   `json:"-"`
	Email        string      `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Password     string      `gorm:"type:text" json:"password"`
	Fullname     string      `gorm:"type:varchar(100);not null" json:"fullname"`
	BirthDate    *CustomTime `gorm:"type:DATE" json:"birth_date"`
	PhoneNumber  *string     `gorm:"type:varchar(13);unique" json:"phonenumber"`
	IsLearner    bool        `gorm:"type:bool;default:false" json:"is_leaner"`
	IsMentor     *bool       `gorm:"type:bool;default:false" json:"is_mentor"`
	ImageURL     *string     `gorm:"type:text" json:"image_url"`
	AuthProvider string      `gorm:"type:enum('email','google');not null;default:email" json:"auth_provider"`

	LearnerBio *learner.LearnerBio `gorm:"foreignKey:UserID;references:ID" json:"learner_bio"`

	MentorBio         *mentor.MentorBiographies   `gorm:"foreignKey:UserID;references:ID" json:"mentor_bio"`
	MentorExperiences []*mentor.MentorExperiences `gorm:"foreignKey:UserID;references:ID" json:"mentor_experience"`
	MentorExpertises  []*mentor.MentorExpertises  `gorm:"foreignKey:UserID;references:ID" json:"mentor_expertise"`
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	str := string(b)

	date, err := time.Parse("2006-01-02", str)

	if err != nil {
		return err
	}

	t.Time = date
	return
}

// Value mengubah CustomTime menjadi nilai yang bisa disimpan ke database
func (t CustomTime) Value() (driver.Value, error) {
	return t.Time.Format("2006-01-02"), nil
}

// Scan mengubah nilai dari database menjadi CustomTime
func (t *CustomTime) Scan(value interface{}) error {
	if value == nil {
		return errors.New("null value for date")
	}

	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case string:
		parsedTime, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		t.Time = parsedTime
		return nil
	default:
		return errors.New("invalid type for CustomTime")
	}
}
