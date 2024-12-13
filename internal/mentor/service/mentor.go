package mentor

import (
	learnerRepo "github.com/philipnathan/pijar-backend/internal/learner/repository"
	custom_error "github.com/philipnathan/pijar-backend/internal/mentor/custom_error"
	repo "github.com/philipnathan/pijar-backend/internal/mentor/repository"
	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
)

type MentorServiceInterface interface {
	GetMentorDetails(MentorID uint) (*userModel.User, error)
	GetMentorLandingPageByUserInterests(UserID uint, page, pageSize int) ([]userModel.User, int, error)
	GetMentorLandingPageByCategory(category_id uint, page, pageSize int) ([]userModel.User, int, error)
}

type MentorService struct {
	repo                 repo.MentorRepositoryInterface
	learnerInterestsRepo learnerRepo.LearnerRepositoryInterface
}

func NewMentorService(repo repo.MentorRepositoryInterface, learnerInterestsRepo learnerRepo.LearnerRepositoryInterface) *MentorService {
	return &MentorService{
		repo:                 repo,
		learnerInterestsRepo: learnerInterestsRepo,
	}
}

func (s *MentorService) GetMentorDetails(MentorID uint) (*userModel.User, error) {
	mentor, err := s.repo.GetMentorDetails(&MentorID)
	if err != nil {
		return nil, err
	}

	return mentor, nil
}

func (s *MentorService) GetMentorLandingPageByUserInterests(UserID uint, page, pageSize int) ([]userModel.User, int, error) {
	userInterests, err := s.learnerInterestsRepo.GetLearnerInterest(UserID)
	if err != nil {
		return nil, 0, custom_error.ErrUserNotFound
	}

	// convert user interest to slice of uint
	var interests []uint
	for _, interest := range userInterests {
		interests = append(interests, interest.CategoryID)
	}

	// if user doesn't have any interest, return all mentors
	if interests == nil {
		category_id := []uint{1}
		mentor, total, err := s.repo.GetMentorByMentorExpertisesCategory(category_id, page, pageSize)
		if err != nil {
			return nil, 0, err
		}

		return mentor, total, nil
	}

	// if user has interest, return mentors with that interest
	mentor, total, err := s.repo.GetMentorByMentorExpertisesCategory(interests, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return mentor, total, nil
}

func (s *MentorService) GetMentorLandingPageByCategory(category_id uint, page, pageSize int) ([]userModel.User, int, error) {
	// convert category_id to slice of uint
	interests := []uint{category_id}
	mentor, total, err := s.repo.GetMentorByMentorExpertisesCategory(interests, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return mentor, total, nil
}
