package dto

import "time"

type GetUpcomingSessionResponse struct {
	Sessions []SessionDetail `json:"sessions"`
}

type SessionDetail struct {
	Day                string    `json:"day"`                  // Day of the week
	Time               string    `json:"time"`                 // Time in HH:MM AM/PM format
	MentorSessionTitle string    `json:"mentor_session_title"` // Title of the session
	ShortDescription   string    `json:"short_description"`    // Brief description
	Schedule           string    `json:"schedule"`             // Date of the session (YYYY-MM-DD format)
	Registered         bool      `json:"registered"`           // Registration status
}
