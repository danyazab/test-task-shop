package model

type items struct {
	ID       int `json:"id"`
	Product  int `json:"name"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

type Order struct {
	ID         int `json:"id"`
	CustomerID int `json:"customer_id"`
	SellerID   int `json:"seller_id"`
	items      []items
}
