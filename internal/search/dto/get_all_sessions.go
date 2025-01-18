package search

type GetAllSessionsResponse struct {
	Sessions []Session `json:"sessions"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
}

type Session struct {
	ID               uint          `json:"session_id"`
	MentorDetails    MentorDetails `json:"mentor_details"`
	Category         string        `json:"category"`
	Title            string        `json:"title"`
	ShortDescription string        `json:"short_description"`
	Detail           string        `json:"detail"`
	Schedule         string        `json:"schedule"`
	Duration         int           `json:"duration"`
	ImageURL         string        `json:"image_url"`
	Link             string        `json:"link"`

	Day           string  `json:"day"`
	Time          string  `json:"time"`
	AverageRating float32 `json:"average_rating"`
}

type MentorDetails struct {
	Id       uint   `json:"id" example:"1"`
	Fullname string `json:"fullname" example:"John Doe"`
	ImageURL string `json:"image_url" example:"https://example.com/image.jpg"`
}
