package notification

type Notif struct {
	ID      uint   `json:"id" example:"1"`
	Type    string `json:"type" example:"schedule_change"`
	Message string `json:"message" example:"schedule for python basic has been changed"`
	IsRead  bool   `json:"is_read" example:"false"`
}

type GetAllNotificationsResponseDto struct {
	Message       string  `json:"message" example:"No notifications found"`
	Notifications []Notif `json:"notification"`
}
