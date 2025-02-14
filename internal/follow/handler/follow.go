package follow

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/follow/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/follow/dto"
	service "github.com/philipnathan/pijar-backend/internal/follow/service"
)

type FollowHandlerInterface interface {
	FollowUnfollowHandler(c *gin.Context)
	IsFollowingHandler(c *gin.Context)
}

type FollowHandler struct {
	service service.FollowServiceInterface
}

func NewFollowHandler(service service.FollowServiceInterface) FollowHandlerInterface {
	return &FollowHandler{
		service: service,
	}
}

// @Summary		Follow or Unfollow Mentor
// @Description	Follow or Unfollow Mentor
// @Scheme
// @Tags		Follow
// @Accept		json
// @Produce	json
// @Param		mentorid	path	int	true	"Mentor ID"
// @Success	200	{object}	FollowUnffolowResponse
// @Failure	400	{object}	CustomError
// @Failure	500	{object}	CustomError
// @Router		/mentors/{mentorid}/follow [post]
func (h *FollowHandler) FollowUnfollowHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	following, err := strconv.Atoi(c.Param("mentor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("Invalid mentor ID"))
		return
	}

	if following <= 0 {
		c.JSON(http.StatusBadRequest, custom_error.ErrInvalidFollowingID)
		return
	}

	if int(following) == int(id) {
		c.JSON(http.StatusBadRequest, custom_error.ErrFollowingSelf)
		return
	}

	followerID := uint(id)
	followingID := uint(following)
	ctx := c.Request.Context()

	err = h.service.FollowUnfollow(ctx, &followerID, &followingID)
	if err != nil {
		switch err {
		case custom_error.ErrNotLearner:
			c.JSON(http.StatusBadRequest, custom_error.ErrNotLearner)
			return
		case custom_error.ErrNotMentor:
			c.JSON(http.StatusBadRequest, custom_error.ErrNotMentor)
			return
		default:
			c.JSON(http.StatusInternalServerError, custom_error.NewCustomError(err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, dto.FollowUnffolowResponse{Messsage: "Follow/Unfollow successful"})
}

// @Summary		Check status of following
// @Description	Check status of following
// @Scheme
// @Tags		Follow
// @Accept		json
// @Produce	json
// @Param		mentorid	path	int	true	"Mentor ID"
// @Success	200	{object}	IsFollowResponse
// @Failure	400	{object}	CustomError
// @Failure	500	{object}	CustomError
// @Router		/mentors/{mentorid}/status [get]
func (h *FollowHandler) IsFollowingHandler(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	following, err := strconv.Atoi(c.Param("mentor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, custom_error.NewCustomError("Invalid mentor_id"))
		return
	}

	if following <= 0 {
		c.JSON(http.StatusBadRequest, custom_error.ErrInvalidFollowingID)
		return
	}

	if int(following) == int(id) {
		c.JSON(http.StatusBadRequest, custom_error.ErrFollowingSelf)
		return
	}

	followerID := uint(id)
	followingID := uint(following)

	status, err := h.service.IsFollowing(&followerID, &followingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, custom_error.NewCustomError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.IsFollowResponse{Message: "Successfully checked", IsFollowing: status})
}
