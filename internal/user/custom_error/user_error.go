package user

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
	ErrUserNotFound                      = NewCustomError("user not found")
	ErrUserExist                         = NewCustomError("user already exist")
	ErrPhoneNumberExist                  = NewCustomError("phone number already exist")
	ErrEmailExist                        = NewCustomError("email already exist")
	ErrLogin                             = NewCustomError("invalid email or password")
	ErrToken                             = NewCustomError("invalid token")
	ErrWrongPassword                     = NewCustomError("wrong password")
	ErrSamePassword                      = NewCustomError("new password cannot be the same as old password")
	ErrWrongPasswordAndLearnerRegistered = NewCustomError("user has been registered as learner, please use that password to register")
	ErrAlreadyMentor                     = NewCustomError("user is already a mentor")
	ErrAlreadyLearner                    = NewCustomError("user is already a learner")
	ErrWrongPasswordAndMentorRegistered  = NewCustomError("user has been registered as mentor, please use that password to register")
	ErrNotUsingGoogle                    = NewCustomError("user is not using google account")
)

type Error struct {
	Error string `json:"error" example:"user not found"`
}
