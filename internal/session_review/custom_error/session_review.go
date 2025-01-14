package session_review

type CustomError struct {
	Message string `json:"error"`
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(message string) *CustomError {
	return &CustomError{Message: message}
}

var (
	ErrUserNotFound        = NewCustomError("user not found")
	ErrSessionNotFound     = NewCustomError("session not found")
	ErrInvalidRating       = NewCustomError("invalid rating. Must be between 1 and 5")
	ErrReviewTooLong       = NewCustomError("review is too long. Must be less than 250 characters")
	ErrUserAlreadyReviewed = NewCustomError("user already reviewed")
	ErrLearnerNotEnrolled  = NewCustomError("learner not enrolled")
)
