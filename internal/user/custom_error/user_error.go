package user

import "errors"

var (
	ErrUserNotFound                      = errors.New("user not found")
	ErrUserExist                         = errors.New("user already exist")
	ErrPhoneNumberExist                  = errors.New("phone number already exist")
	ErrEmailExist                        = errors.New("email already exist")
	ErrLogin                             = errors.New("invalid email or password")
	ErrToken                             = errors.New("invalid token")
	ErrWrongPassword                     = errors.New("wrong password")
	ErrSamePassword                      = errors.New("new password cannot be the same as old password")
	ErrWrongPasswordAndLearnerRegistered = errors.New("user has been registered as learner, please use that password to register")
	ErrAlreadyMentor                     = errors.New("user is already a mentor")
)

type Error struct {
	Error string `json:"error" example:"user not found"`
}
