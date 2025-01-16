package session

type GetAllSessionsResponse struct {
	Sessions []Session `json:"sessions"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
}

type Session struct {
	ID               uint          `json:"session_id"`
	Day              string        `json:"day"`
	Time             string        `json:"time"`
	Title            string        `json:"title"`
	ShortDescription string        `json:"short_description"`
	ImageURL         string        `json:"image_url"`
	Schedule         string        `json:"schedule"`
	MentorDetails    MentorDetails `json:"mentor_details"`
	AverageRating    float32       `json:"average_rating"`
	Duration         int           `json:"duration"`
}

type MentorDetails struct {
	Id       uint   `json:"id" example:"1"`
	Fullname string `json:"fullname" example:"John Doe"`
	ImageURL string `json:"image_url" example:"https://example.com/image.jpg"`
}
