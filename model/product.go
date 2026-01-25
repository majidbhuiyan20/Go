package model

type Product struct {
	ID          string  `json:"id"`
	CategoryID  string  `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
	CreatedAt   int64   `json:"created_at"`
}
