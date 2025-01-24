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
	user05, _ := utils.HashPassword("hashed_password_wayan")
	user06, _ := utils.HashPassword("hashed_password_komang")
	user07, _ := utils.HashPassword("hashed_password_ari")
	user08, _ := utils.HashPassword("hashed_password_supriyadi")
	user09, _ := utils.HashPassword("Password123")
	user10, _ := utils.HashPassword("Password123")
	user11, _ := utils.HashPassword("Password123")
	user12, _ := utils.HashPassword("Password123")

	users := []userModel.User{
		{
			Email:       "andi.budi@example.com",
			Password:    user01,
			Fullname:    "Andi Budi",
			IsLearner:   false,
			IsMentor:    &[]bool{true}[0],
			BirthDate:   &userModel.CustomTime{Time: time.Date(1985, 3, 21, 0, 0, 0, 0, time.UTC)},
			ImageURL:    &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ29FsGvdObT4DjbQEj_gA-5NqGs7m7e5raog&s"}[0],
			PhoneNumber: &[]string{"081234567890"}[0],
		},
		{
			Email:       "nurul.aini@example.com",
			Password:    user02,
			Fullname:    "Nurul Aini",
			IsLearner:   true,
			IsMentor:    &[]bool{false}[0],
			BirthDate:   &userModel.CustomTime{Time: time.Date(1997, 7, 12, 0, 0, 0, 0, time.UTC)},
			ImageURL:    &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTF69UQPqIO-jP-3C6LeR7NxRL5FVVosp56VA&s"}[0],
			PhoneNumber: &[]string{"081987654321"}[0],
		},
		{
			Email:     "siti.maemunah@example.com",
			Password:  user03,
			Fullname:  "Siti Maemunah",
			IsLearner: false,
			IsMentor:  &[]bool{true}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1990, 2, 15, 0, 0, 0, 0, time.UTC),
			},
			ImageURL:    &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsPqnrRwfV3Oa-Cuc_f08NCCLvS2O9yzBHNg&s"}[0],
			PhoneNumber: &[]string{"081345678901"}[0],
		},
		{
			Email:     "agus.suryadi@example.com",
			Password:  user04,
			Fullname:  "Agus Suryadi",
			IsLearner: true,
			IsMentor:  &[]bool{false}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1988, 11, 5, 0, 0, 0, 0, time.UTC),
			},
			ImageURL:    &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQKlbHGf_ERPpChbcnvpjn6aC2TMycroaR8IQ&s"}[0],
			PhoneNumber: &[]string{"081678912345"}[0],
		},
		{
			Email:     "wayan.budiasa@example.com",
			Password:  user05,
			Fullname:  "Wayan Budiasa",
			IsLearner: false,
			IsMentor:  &[]bool{true}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1960, 3, 21, 0, 0, 0, 0, time.UTC),
			},
			ImageURL: &[]string{"https://st.depositphotos.com/1715570/2652/i/450/depositphotos_26521259-stock-photo-portrait-of-a-handsome-young.jpg"}[0],
		},
		{
			Email:     "komang.ariasa@example.com",
			Password:  user06,
			Fullname:  "Komang Ariasa",
			IsLearner: true,
			IsMentor:  &[]bool{false}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1997, 7, 12, 0, 0, 0, 0, time.UTC),
			},
			ImageURL: &[]string{"https://img.freepik.com/free-photo/artist-white_1368-3546.jpg"}[0],
		},
		{
			Email:     "ari.komang@example.com",
			Password:  user07,
			Fullname:  "Ari Komang",
			IsLearner: false,
			IsMentor:  &[]bool{true}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1960, 3, 21, 0, 0, 0, 0, time.UTC),
			},
			ImageURL: &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTvi7HpQ-_PMSMOFrj1hwjp6LDcI-jm3Ro0Xw&s"}[0],
		},
		{
			Email:     "supriyadi@example.com",
			Password:  user08,
			Fullname:  "Supriyadi",
			IsLearner: true,
			IsMentor:  &[]bool{false}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1997, 7, 12, 0, 0, 0, 0, time.UTC),
			},
			ImageURL: &[]string{"https://static.vecteezy.com/system/resources/thumbnails/009/887/693/small_2x/male-man-african-american-black-diversity-person-afro-hair-ethnic-happy-smile-model-close-up-face-enjoyment-hashion-lifestyle-professional-human-father-boy-business-education-young-adult-teenage-photo.jpg"}[0],
		},
		{
			Email:     "sandhika.galih@example.com",
			Password:  user09,
			Fullname:  "Sandhika Galih",
			IsLearner: false,
			IsMentor:  &[]bool{true}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1982, 3, 21, 0, 0, 0, 0, time.UTC),
			},
			ImageURL: &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSyKzYF_Z55a7kUWIFk5lmK_QJUiMqNANxKUA&s"}[0],
		},
		{
			Email:     "della.wardhini@example.com",
			Password:  user10,
			Fullname:  "Della Wardhini",
			IsLearner: false,
			IsMentor:  &[]bool{true}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1960, 3, 21, 0, 0, 0, 0, time.UTC),
			},
			ImageURL: &[]string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT-q3W9cs5ucieReoKu7IhRRx0dBzwUHs9gcA&s "}[0],
		},
		{
			Email:     "nandia.saya@example.com",
			Password:  user11,
			Fullname:  "Nandia Saya",
			IsLearner: false,
			IsMentor:  &[]bool{true}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1999, 3, 21, 0, 0, 0, 0, time.UTC),
			},
			ImageURL: &[]string{"https://pngimg.com/d/thinking_woman_PNG11633.png"}[0],
		},
		{
			Email:     "budiyanto.jilal@example.com",
			Password:  user12,
			Fullname:  "Budiyanto Jilal",
			IsLearner: false,
			IsMentor:  &[]bool{true}[0],
			BirthDate: &userModel.CustomTime{
				Time: time.Date(1990, 3, 21, 0, 0, 0, 0, time.UTC),
			},
			ImageURL: &[]string{"https://png.pngtree.com/png-vector/20240611/ourmid/pngtree-male-teacher-holding-a-book-wearing-eye-glasses-png-image_12705620.png"}[0],
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
