package session

import "time"

type GetDetailSessionResponse struct {
	SessionID          uint          `json:"session_id"`
	MentorSessionTitle string        `json:"mentor_session_title"`
	ShortDescription   string        `json:"short_description"`
	ImageURL           string        `json:"image_url"`
	Schedule           time.Time     `json:"schedule"`
	MentorDetails      MentorDetails `json:"mentor_details"`
	AverageRating      float32       `json:"average_rating"`
}
