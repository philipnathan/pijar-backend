package mentor

import (
	"database/sql/driver"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CustomTime struct {
	time.Time
}

type MentorExperiences struct {
	gorm.Model  `json:"-"`
	UserID      uint       `gorm:"not null" json:"user_id"`
	Occupation  string     `gorm:"type:varchar(50);not null" json:"occupation"`
	CompanyName string     `gorm:"type:varchar(100);not null" json:"company_name"`
	StartDate   CustomTime `gorm:"type:DATE;not null" json:"start_date"`
	EndDate     CustomTime `gorm:"type:DATE" json:"end_date"`
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

func (t *CustomTime) FormatToString() string {
	return t.Time.Format("2006-01-02")
}
