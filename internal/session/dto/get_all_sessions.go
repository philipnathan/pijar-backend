package session

import (
	"time"
)

type GetAllSessionsResponse struct {
	Sessions []Session `json:"sessions"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
}

type Session struct {
	MentorSessionTitle string        `json:"mentor_session_title"`
	ShortDescription   string        `json:"short_description"`
	ImageURL           string        `json:"image_url"`
	Schedule           time.Time     `json:"schedule"`
	MentorDetails      MentorDetails `json:"mentor_details"`
	AverageRating      float32       `json:"average_rating"`
}
