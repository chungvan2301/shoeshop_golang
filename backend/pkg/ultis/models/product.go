package models

import "mime/multipart"

type ProductResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"imageURL"`
	Gender   string  `json:"gender"`
}

type ProductInput struct {
	Name          string                `form:"name"`
	Brand         string                `form:"brand"`
	Gender        string                `form:"gender"`
	Category      string                `form:"category"`
	Price         float64               `form:"price"`
	IsInInventory bool                  `form:"is_in_inventory"`
	ItemsLeft     int                   `form:"items_left"`
	ImageURL      string                `form:"imageURL"`
	Slug          string                `form:"slug"`
	ImageFile     *multipart.FileHeader `form:"image"`
}

type ProductCart struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"imageURL"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
