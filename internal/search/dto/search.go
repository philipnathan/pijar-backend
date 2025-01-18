package search

type SearchResponse struct {
	Sessions   GetAllSessionsResponse    `json:"sessions"`
	Mentors    []MentorDetailLandingPage `json:"mentors"`
	Categories []Category                `json:"categories"`
}

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	ImageURL     string `json:"image_url"`
}
