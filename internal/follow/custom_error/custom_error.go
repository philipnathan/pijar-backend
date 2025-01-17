package follow

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
	ErrInvalidFollowingID = NewCustomError("Invalid following_id")
	ErrFollowingSelf      = NewCustomError("You can't follow yourself")
	ErrNotLearner         = NewCustomError("User is not a learner")
	ErrNotMentor          = NewCustomError("User is not a mentor")
)
