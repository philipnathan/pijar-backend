package user

type UserProfileResponse struct {
    ID       uint   `json:"id"`
    Fullname string `json:"fullname"`
    Email    string `json:"email"`
    ImageURL string `json:"image_url"`
}