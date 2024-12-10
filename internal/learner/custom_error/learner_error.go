package learner

import (
	"errors"
)

type Error struct {
	Error string `json:"error" example:"interest not found"`
}

var (
	ErrLearnerBioNotFound     = errors.New("learner bio not found. Please add bio first")
	ErrLearnerBioAlreadyExist = errors.New("bio already exist. Please update bio")
)
