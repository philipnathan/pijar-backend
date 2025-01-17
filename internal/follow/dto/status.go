package follow

type IsFollowResponse struct {
	Message     string `json:"message"`
	IsFollowing bool   `json:"is_following"`
}
