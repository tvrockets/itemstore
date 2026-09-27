package model

type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description, omitempty"`
	Price       int    `json:"price"`
}
