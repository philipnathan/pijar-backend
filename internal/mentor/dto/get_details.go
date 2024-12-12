package mentor

import (
	model "github.com/philipnathan/pijar-backend/internal/mentor/model"
)

type MentorExperiences struct {
	Ocupation   *string           `json:"occupation"`
	CompanyName *string           `json:"company_name"`
	StartDate   *model.CustomTime `json:"start_date"`
	EndDate     *model.CustomTime `json:"end_date"`
}

type MentorExpertises struct {
	Expertise *string `json:"expertise"`
	Category  *string `json:"category"`
}

type GetMentorDetailsDto struct {
	UserID    uint    `json:"user_id" example:"1"`
	Fullname  string  `json:"fullname" example:"John Doe"`
	ImageURL  *string `json:"image_url" example:"https://example.com/image.jpg"`
	MentorBio string  `json:"mentor_bio"`

	MentorExperiences []*MentorExperiences `json:"mentor_experiences"`
	MentorExpertises  []*MentorExpertises  `json:"mentor_expertise"`
}
