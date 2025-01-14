package session_review

type SessionReviewRequest struct {
	Review *string `json:"review"`
	Rating uint    `json:"rating" `
}

type SessionReviewResponse struct {
	Message string `json:"message"`
}
