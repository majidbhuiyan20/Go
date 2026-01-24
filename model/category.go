package model

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	CreatedAt int64  `json:"created_at"`
}
