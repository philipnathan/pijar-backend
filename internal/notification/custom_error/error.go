package notification

import "errors"

type Error struct {
	Error string `json:"error"`
}

var ErrNotificationHasBeenRead = errors.New("notification has been read")
