package follow

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	custom_error "github.com/philipnathan/pijar-backend/internal/follow/custom_error"
	dto "github.com/philipnathan/pijar-backend/internal/follow/dto"
	service "github.com/philipnathan/pijar-backend/internal/follow/service"
)

type FollowHandler struct {
	service service.FollowServiceInterface
}

func NewFollowHandler(service service.FollowServiceInterface) *FollowHandler {
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
// @Security	Bearer
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

	following, err := strconv.Atoi(c.Param("mentorid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid following_id"})
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

	err = h.service.FollowUnfollow(&followerID, &followingID)
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
