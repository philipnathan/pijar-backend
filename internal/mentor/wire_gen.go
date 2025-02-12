// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package mentor

import (
	"github.com/google/wire"
	"github.com/philipnathan/pijar-backend/internal/learner/repository"
	"github.com/philipnathan/pijar-backend/internal/mentor/handler"
	mentor2 "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	mentor3 "github.com/philipnathan/pijar-backend/internal/mentor/service"
	"gorm.io/gorm"
)

// Injectors from mentor_bio_wire.go:

func InitializedMentorBio(db *gorm.DB) (mentor.MentorBioHandlerInterface, error) {
	mentorBioRepositoryInterface := mentor2.NewMentorBioRepository(db)
	mentorBioServiceInterface := mentor3.NewMentorBioService(mentorBioRepositoryInterface)
	mentorBioHandlerInterface := mentor.NewMentorBioHandler(mentorBioServiceInterface)
	return mentorBioHandlerInterface, nil
}

// Injectors from mentor_wire.go:

func InitializedMentor(db *gorm.DB) (mentor.MentorHandlerInterface, error) {
	mentorRepositoryInterface := mentor2.NewMentorRepository(db)
	learnerRepositoryInterface := learner.NewLearnerRepository(db)
	mentorServiceInterface := mentor3.NewMentorService(mentorRepositoryInterface, learnerRepositoryInterface)
	mentorHandlerInterface := mentor.NewMentorHandler(mentorServiceInterface)
	return mentorHandlerInterface, nil
}

// mentor_bio_wire.go:

var MentorBioProviderSet = wire.NewSet(mentor2.NewMentorBioRepository, mentor3.NewMentorBioService, mentor.NewMentorBioHandler)

// mentor_wire.go:

var MentorProviderSet = wire.NewSet(mentor2.NewMentorRepository, mentor3.NewMentorService, mentor.NewMentorHandler)
