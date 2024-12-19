package repository

import (
	"fmt"

	categoryModel "github.com/philipnathan/pijar-backend/internal/category/model"
	dto "github.com/philipnathan/pijar-backend/internal/search/dto"
	sessionModel "github.com/philipnathan/pijar-backend/internal/session/model"
	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
	"gorm.io/gorm"
)

type SearchRepositoryInterface interface {
    SearchSessions(keyword string) ([]dto.SessionDetail, error)
    SearchMentors(keyword string) ([]dto.MentorDetail, error)
    SearchCategories(keyword string) ([]dto.CategoryDetail, error)
}

type SearchRepository struct {
    db *gorm.DB
}

func NewSearchRepository(db *gorm.DB) SearchRepositoryInterface {
    return &SearchRepository{
        db: db,
    }
}

func (r *SearchRepository) SearchSessions(keyword string) ([]dto.SessionDetail, error) {
    var sessions []sessionModel.MentorSession
    err := r.db.Where("title LIKE ? OR short_description LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&sessions).Error
    fmt.Println(sessions)
    fmt.Println(err)
    if err != nil {
        return nil, err
    }

    var sessionDetails []dto.SessionDetail
    for _, session := range sessions {
        sessionDetails = append(sessionDetails, dto.SessionDetail{
            Title:            session.Title,
            ShortDescription: session.ShortDescription,
            Schedule:         session.Schedule.Format("2006-01-02 15:04:05"),
            ImageURL:         session.ImageURL,
        })
    }
    return sessionDetails, nil
}

func (r *SearchRepository) SearchMentors(keyword string) ([]dto.MentorDetail, error) {
    var mentors []userModel.User
    err := r.db.Where("fullname LIKE ? OR email LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&mentors).Error
    if err != nil {
        return nil, err
    }

    var mentorDetails []dto.MentorDetail
    
    for _, mentor := range mentors {
        
        mentorDetail := dto.MentorDetail{
            Fullname: mentor.Fullname,
            Email:    mentor.Email,
        }
        if mentor.ImageURL != nil {
            mentorDetail.ImageURL = *mentor.ImageURL
        }
        mentorDetails = append(mentorDetails, mentorDetail)
    }
    fmt.Println("test")
    return mentorDetails, nil
}

func (r *SearchRepository) SearchCategories(keyword string) ([]dto.CategoryDetail, error) {
    var categories []categoryModel.Category
    err := r.db.Where("category_name LIKE ? ", "%"+keyword+"%").Find(&categories).Error
    if err != nil {
        return nil, err
    }

    var categoryDetails []dto.CategoryDetail
    for _, category := range categories {
        categoryDetails = append(categoryDetails, dto.CategoryDetail{
            Title: category.Category_name,
        })
    }
    return categoryDetails, nil
}