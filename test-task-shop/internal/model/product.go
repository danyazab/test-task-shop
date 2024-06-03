package model

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	SellerID    int    `json:"seller_id"`
	IsDeleted   bool   `json:"is_deleted"`
}
