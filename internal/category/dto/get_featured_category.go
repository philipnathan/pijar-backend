package dto

type FeaturedCategoryResponseDto struct {
    CategoryName string `json:"category_name" example:"Coding Basics"`
    ImageURL     string `json:"image_url" example:"https://example.com/image.png"`
}