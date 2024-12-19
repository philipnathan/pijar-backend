package seed

import (
	"time"

	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
	userRepo "github.com/philipnathan/pijar-backend/internal/user/repository"
	utils "github.com/philipnathan/pijar-backend/utils"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	user01, _ := utils.HashPassword("hashed_password_andi")
	user02, _ := utils.HashPassword("hashed_password_nurul")
	user03, _ := utils.HashPassword("hashed_password_siti")
	user04, _ := utils.HashPassword("hashed_password_agus")

	users := []userModel.User{
		{
			Email:       "andi.budi@example.com",
			Password:    user01,
			Fullname:    "Andi Budi",
			IsMentor:    &[]bool{true}[0],
			BirthDate:   &userModel.CustomTime{Time: time.Date(1985, 3, 21, 0, 0, 0, 0, time.UTC)},
			ImageURL:    &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ29FsGvdObT4DjbQEj_gA-5NqGs7m7e5raog&s"}[0],
			PhoneNumber: &[]string{"081234567890"}[0],
		},
		{
			Email:       "nurul.aini@example.com",
			Password:    user02,
			Fullname:    "Nurul Aini",
			IsMentor:    &[]bool{false}[0],
			BirthDate:   &userModel.CustomTime{Time: time.Date(1997, 7, 12, 0, 0, 0, 0, time.UTC)},
			ImageURL:    &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTF69UQPqIO-jP-3C6LeR7NxRL5FVVosp56VA&s"}[0],
			PhoneNumber: &[]string{"081987654321"}[0],
		},
		{
			Email:    "siti.maemunah@example.com",
			Password: user03,
			Fullname: "Siti Maemunah",
			IsMentor: &[]bool{true}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1990, 2, 15, 0, 0, 0, 0, time.UTC),
			},
			ImageURL:    &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsPqnrRwfV3Oa-Cuc_f08NCCLvS2O9yzBHNg&s"}[0],
			PhoneNumber: &[]string{"081345678901"}[0],
		},
		{
			Email:    "agus.suryadi@example.com",
			Password: user04,
			Fullname: "Agus Suryadi",
			IsMentor: &[]bool{false}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1988, 11, 5, 0, 0, 0, 0, time.UTC),
			},
			ImageURL:    &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQKlbHGf_ERPpChbcnvpjn6aC2TMycroaR8IQ&s"}[0],
			PhoneNumber: &[]string{"081678912345"}[0],
		},
	}

	var count int64
	if err := db.Model(&userModel.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for _, user := range users {
		if err := userRepo.NewUserRepository(db).SaveUser(&user); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
