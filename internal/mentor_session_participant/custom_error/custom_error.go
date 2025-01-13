package mentor_session_participant

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
	ErrInvalidRating          = NewCustomError("invalid rating. Must be between 1 and 5")
	ErrCommentTooLong         = NewCustomError("comment is too long. Must be less than 200 characters")
	ErrUserNotFound           = NewCustomError("user not found")
	ErrSessionNotFound        = NewCustomError("session not found")
	ErrSessionAlreadyFinished = NewCustomError("session already finished")
	ErrUserAlreadyRegistered  = NewCustomError("user already registered")
)
