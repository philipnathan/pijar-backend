package search

type MentorDetailLandingPage struct {
	Id         uint   `json:"id" example:"1"`
	Fullname   string `json:"fullname" example:"John Doe"`
	ImageURL   string `json:"image_url" example:"https://example.com/image.jpg"`
	Occupation string `json:"occupation" example:"Software Engineer"`
}

type MentorLandingPageResponseDto struct {
	Message     string `json:"message" example:"success"`
	Data        *[]MentorDetails
	CountData   int
	CurrentPage int
	DataPerPage int
}
