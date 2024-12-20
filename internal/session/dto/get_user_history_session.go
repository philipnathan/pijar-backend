package session

import "time"

type GetUserHistorySessionResponseDto struct {
	Histories []History `json:"histories"`
}

type History struct {
	MentorSessionTitle string        `json:"mentor_session_title"`
	ShortDescription   string        `json:"short_description"`
	Schedule           time.Time     `json:"schedule"`
	Status             string        `json:"status"`
	MentorDetails      MentorDetails `json:"mentor_details"`
}
