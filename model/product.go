package model

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"` // image path or URL
	Stock       int      `json:"stock"`
	SKU         string   `json:"sku"`
	Categories  []string `json:"categories"`
	Tags        []string `json:"tags"`
	IsFeatured  bool     `json:"is_featured"`
	IsActive    bool     `json:"is_active"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
}
