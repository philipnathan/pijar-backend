package mentor

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/mentor/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/mentor/dto"
	service "github.com/philipnathan/pijar-backend/internal/mentor/service"
	middleware "github.com/philipnathan/pijar-backend/middleware"
)

type MentorHandler struct {
	service service.MentorServiceInterface
}

func NewMentorHandler(service service.MentorServiceInterface) *MentorHandler {
	return &MentorHandler{
		service: service,
	}
}

// @Summary	Get mentor details
// @Schemes
// @Description	Get mentor details
// @Tags			Mentor
// @Produce		json
// @Param			mentor_id	path		uint	true	"mentor_id"
// @Success		200			{object}	GetMentorDetailsDto
// @Failure		400			{object}	Error	"Invalid mentor_id"
// @Failure		404			{object}	Error	"Mentor not found"
// @Failure		500			{object}	Error	"Internal server error"
// @Router			/mentors/{mentor_id} [get]
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
		occupation := experience.Occupation
		companyName := experience.CompanyName
		startDate := experience.StartDate.FormatToString()

		var endDate string
		if !experience.EndDate.IsZero() {
			endDate = experience.EndDate.FormatToString()
		} else {
			endDate = ""
		}

		MentorExperiences = append(MentorExperiences, &dto.MentorExperiences{
			Ocupation:   occupation,
			CompanyName: companyName,
			StartDate:   startDate,
			EndDate:     endDate,
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
		Occupation:        MentorExperiences[0].Ocupation,
		MentorExperiences: MentorExperiences,
		MentorExpertises:  MentorExpertises,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary	Get mentor landing page
// @Schemes
// @Description	Get mentor landing page
// @Tags			Mentor
// @Produce		json
// @Param			page		query		int	false	"page"
// @Param			pagesize	query		int	false	"pagesize"
// @Param			categoryid	query		int	false	"categoryid"
// @Success		200			{object}	MentorLandingPageResponseDto
// @Failure		500			{object}	Error	"Internal server error"
// @Router			/mentors/landingpage [get]
func (h *MentorHandler) UserGetMentorLandingPage(c *gin.Context) {
	// check if user is authenticated
	var isAuthenticated bool
	authHeader, err := c.Cookie("access_token")
	if err != nil {
		isAuthenticated = false
	}
	if authHeader != "" {
		middleware.AuthMiddleware()(c)

		if !c.IsAborted() {
			isAuthenticated = true
		}
	}

	// get page and page_size from query
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	categoryIDint, _ := strconv.Atoi(c.DefaultQuery("categoryid", "0"))
	var categoryIDuint uint

	// if category_id is not available or negative (invalid)
	if categoryIDint < 0 {
		c.JSON(http.StatusBadRequest, custom_error.Error{Error: "category_id is invalid"})
		return
	}

	// if category_id is available (categoryID is the priority because it is an input fro guest/user)
	if categoryIDint > 0 {
		categoryIDuint = uint(categoryIDint)
		mentors, total, err := h.service.GetMentorLandingPageByCategory(categoryIDuint, page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.NewMentorLandingPageResponseDto(&mentors, total, page, pageSize))
		return
	}

	// if user is not authenticated (without JWT) and category_id is not available
	if !isAuthenticated && categoryIDint == 0 {
		mentors, total, err := h.service.GetAllMentors(page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.NewMentorLandingPageResponseDto(mentors, total, page, pageSize))
		return
	}

	// if user is authenticated (with JWT) - get user_id from JWT and get mentor by user interests
	UserID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	userIDuint, ok := UserID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, custom_error.Error{Error: "Unauthorized"})
		return
	}

	mentors, total, err := h.service.GetMentorLandingPageByUserInterests(uint(userIDuint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.NewMentorLandingPageResponseDto(&mentors, total, page, pageSize))
}
