package custom_error

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExist = errors.New("user already exist")
	ErrPhoneNumberExist = errors.New("phone number already exist")
	ErrEmailExist = errors.New("email already exist")
	ErrLogin = errors.New("invalid email or password")
	ErrToken = errors.New("invalid token")
)