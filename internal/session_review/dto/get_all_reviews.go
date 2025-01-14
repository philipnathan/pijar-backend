package session_review

type GetAllReviewsResponse struct {
	Message   string         `json:"message"`
	SessionID uint           `json:"session_id"`
	Page      int            `json:"page"`
	PageSize  int            `json:"page_size"`
	Total     int            `json:"total"`
	Reviews   []ReviewDetail `json:"reviews"`
}

type UserDetails struct {
	Fullname string `json:"fullname"`
	ImageURL string `json:"image_url"`
}

type ReviewDetail struct {
	ReviewID    uint        `json:"review_id"`
	UserDetails UserDetails `json:"user_details"`
	Rating      uint        `json:"rating"`
	Review      string      `json:"review"`
}
