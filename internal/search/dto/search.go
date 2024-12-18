package search

type SearchResponse struct {
	Sessions []SessionDetail  `json:"sessions"`
	Mentors  []MentorDetail   `json:"mentors"`
	Topics   []CategoryDetail `json:"topics"`
}

type SessionDetail struct {
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Schedule         string `json:"schedule"`
	ImageURL         string `json:"image_url"`
}

type MentorDetail struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	ImageURL string `json:"image_url"`
}

type CategoryDetail struct {
	Title string `json:"category_name"`
}