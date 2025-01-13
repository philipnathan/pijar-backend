package mentor_session_participant

import "time"

type EnrollmentResponse struct {
	Message     string              `json:"message"`
	Enrollments []EnrollmentDetails `json:"enrollments"`
	Total       int                 `json:"total"`
	Page        int                 `json:"page"`
	PageSize    int                 `json:"page_size"`
}

type EnrollmentDetails struct {
	MentorSessionParticipantID uint           `json:"mentor_session_participant_id"`
	SessionDetails             SessionDetails `json:"session_details"`
	Status                     string         `json:"status"`
}

type SessionDetails struct {
	MentorSessionID    uint      `json:"mentor_session_id"`
	MentorSessionTitle string    `json:"mentor_session_title"`
	ShortDescription   string    `json:"short_description"`
	ImageURL           string    `json:"image_url"`
	Schedule           time.Time `json:"schedule"`
}
