package user

import (
	"database/sql/driver"
	"errors"
	"time"

	learner "github.com/philipnathan/pijar-backend/internal/learner/model"
	"gorm.io/gorm"
)

type CustomTime struct {
	time.Time
}

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(50);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:text;not null" json:"password"`
	Fullname string `gorm:"type:varchar(100);not null" json:"fullname"`
	BirthDate CustomTime `gorm:"type:DATE;not null" json:"birthdate"`
	PhoneNumber string `gorm:"type:varchar(13);not null;unique" json:"phonenumber"`
	IsMentor *bool `gorm:"type:bool;not null;default:false"`
	ImageURL *string `gorm:"type:text" json:"imageurl"`
	LearnerBio *learner.LearnerBio `gorm:"foreignKey:UserID;references:ID" json:"learner_bio"`
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	str := string(b)
	str = str[1:len(str)-1]

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