package session

import "errors"

type SessionError struct {
	Err error `json:"error"`
}

var (
	ErrSessionNotFound = errors.New("session not found")
)

