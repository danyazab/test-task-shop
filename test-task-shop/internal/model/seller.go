package model

type Seller struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required,max=255"`
	Phone string `json:"phone" validate:"required,max=50"`
}
