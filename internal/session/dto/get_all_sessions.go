package session

type GetAllSessionsResponse struct {
	AllSessions []MentorSessionsDetail `json:"sessions"`
	Page        int                    `json:"page"`
	PageSize    int                    `json:"page_size"`
	Total       int                    `json:"total"`
}

type MentorSessionsDetail struct {
	MentorSessions []SessionDetail `json:"mentor_sessions"`
	MentorDetails  MentorDetails   `json:"mentor_details"`
}
