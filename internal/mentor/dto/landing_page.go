package mentor

import (
	userModel "github.com/philipnathan/pijar-backend/internal/user/model"
)

type MentorDetails struct {
	Id         uint   `json:"id" example:"1"`
	Fullname   string `json:"fullname" example:"John Doe"`
	ImageURL   string `json:"image_url" example:"https://example.com/image.jpg"`
	Occupation string `json:"occupation" example:"Software Engineer"`
}

type MentorLandingPageResponseDto struct {
	Message     string `json:"message" example:"success"`
	Data        *[]MentorDetails
	CountData   int
	CurrentPage int
	DataPerPage int
}

func NewMentorLandingPageResponseDto(data *[]userModel.User, total, page, pageSize int) *MentorLandingPageResponseDto {
	var mentorDetails []MentorDetails

	var mentorImageURL string

	for _, mentor := range *data {
		if mentor.ImageURL != nil {
			mentorImageURL = *mentor.ImageURL
		} else {
			mentorImageURL = ""
		}

		mentorDetails = append(mentorDetails, MentorDetails{
			Id:         mentor.ID,
			Fullname:   mentor.Fullname,
			ImageURL:   mentorImageURL,
			Occupation: mentor.MentorExperiences[0].Occupation,
		})
	}

	return &MentorLandingPageResponseDto{
		Message:     "success",
		Data:        &mentorDetails,
		CountData:   total,
		CurrentPage: page,
		DataPerPage: pageSize,
	}
}
