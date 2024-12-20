package session

type GetUpcomingSessionResponse struct {
	Sessions []SessionDetail `json:"sessions"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Total    int             `json:"total"`
}

type SessionDetail struct {
	Day              string `json:"day"`
	Time             string `json:"time"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Schedule         string `json:"schedule"`
	ImageURL         string `json:"image_url"`
	Registered       bool   `json:"registered"`
	Duration         int    `json:"duration"`
}
