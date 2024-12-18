package mentor

type MentorExperiences struct {
	Ocupation   string `json:"occupation"`
	CompanyName string `json:"company_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date" example:"2022-01-01" omitempty:"true"`
}

type MentorExpertises struct {
	Expertise *string `json:"expertise"`
	Category  *string `json:"category"`
}

type GetMentorDetailsDto struct {
	UserID     uint    `json:"user_id" example:"1"`
	Fullname   string  `json:"fullname" example:"John Doe"`
	ImageURL   *string `json:"image_url" example:"https://example.com/image.jpg"`
	MentorBio  string  `json:"mentor_bio"`
	Occupation string  `json:"occupation" example:"Software Engineer"`

	MentorExperiences []*MentorExperiences `json:"mentor_experiences"`
	MentorExpertises  []*MentorExpertises  `json:"mentor_expertise"`
}
