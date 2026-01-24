package model

type Category struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Image    string    `json:"image"`
	Products []Product `json:"product"`
}
