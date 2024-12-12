package mentor

import "errors"

type Error struct {
	Error string `json:"error"`
}

var (
	ErrMentorBioNotFound = errors.New("mentor bio not found")
	ErrMentorNotFound    = errors.New("mentor not found")
)
