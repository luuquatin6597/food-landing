package models

// Food đại diện cho thông tin món ăn
type Food struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	Price       string `json:"price"`
	ImageURL    string `json:"image_url"`
	Color       string `json:"color"`
	Region      string `json:"region"`
}
