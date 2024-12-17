package session

import "time"

type MentorSessionResponse struct {
	MentorSessionTitle string        `json:"mentor_session_title"`
	ShortDescription   string        `json:"short_description"`
	Schedule           time.Time     `json:"schedule"`
	Registered         bool          `json:"registered"`
	MentorDetails      MentorDetails `json:"mentor_details"`
}

type GetUpcomingMentorSessionResponse struct {
	Sessions []MentorSessionResponse `json:"sessions"`
}

type MentorDetails struct {
	Id       uint   `json:"id" example:"1"`
	Fullname string `json:"fullname" example:"John Doe"`
	ImageURL string `json:"image_url" example:"https://example.com/image.jpg"`
}
