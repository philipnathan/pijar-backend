package dto

type MentorSessionResponse struct {
    Day                string `json:"day"`
    Time               string `json:"time"`
    MentorSessionTitle string `json:"mentor_session_title"`
    ShortDescription   string `json:"short_description"`
    Schedule           string `json:"schedule"`
    Registered         bool   `json:"registered"`
}

type GetUpcomingMentorSessionResponse struct {
    Sessions []MentorSessionResponse `json:"sessions"`
}