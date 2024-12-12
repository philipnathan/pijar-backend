package mentor

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/mentor/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/mentor/dto"
	service "github.com/philipnathan/pijar-backend/internal/mentor/service"
)

type MentorHandler struct {
	service service.MentorServiceInterface
}

func NewMentorHandler(service service.MentorServiceInterface) *MentorHandler {
	return &MentorHandler{
		service: service,
	}
}

func (h *MentorHandler) UserGetMentorDetails(c *gin.Context) {
	MentorIDStr := c.Param("mentor_id")
	if MentorIDStr == "" {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "mentor_id is required"})
		return
	}

	MentorID, err := strconv.ParseUint(MentorIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "mentor_id is invalid"})
		return
	}

	uintMentorID := uint(MentorID)
	mentorDetails, err := h.service.GetMentorDetails(uintMentorID)
	if err != nil {
		if err == custom_error.ErrMentorNotFound {
			c.JSON(http.StatusNotFound, custom_error.Error{Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		}
	}

	var MentorExperiences []*dto.MentorExperiences
	for _, experience := range mentorDetails.MentorExperiences {
		occupation := experience.Ocupation
		companyName := experience.CompanyName
		startDate := experience.StartDate
		endDate := experience.EndDate

		MentorExperiences = append(MentorExperiences, &dto.MentorExperiences{
			Ocupation:   &occupation,
			CompanyName: &companyName,
			StartDate:   &startDate,
			EndDate:     &endDate,
		})
	}

	var MentorExpertises []*dto.MentorExpertises
	for _, expertise := range mentorDetails.MentorExpertises {
		category := expertise.Category.Category_name
		expertise := expertise.Expertise

		MentorExpertises = append(MentorExpertises, &dto.MentorExpertises{
			Expertise: &expertise,
			Category:  &category,
		})
	}

	response := &dto.GetMentorDetailsDto{
		UserID:            mentorDetails.ID,
		Fullname:          mentorDetails.Fullname,
		ImageURL:          mentorDetails.ImageURL,
		MentorBio:         mentorDetails.MentorBio.Bio,
		MentorExperiences: MentorExperiences,
		MentorExpertises:  MentorExpertises,
	}

	c.JSON(http.StatusOK, response)
}
