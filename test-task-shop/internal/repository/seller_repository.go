package repository

import (
	"TestTaskShop/internal/database"
	"TestTaskShop/internal/model"
)

type SellerRepository interface {
	Create(seller model.Seller) (int, error)
	GetByID(id int) (*model.Seller, error)
	Update(seller model.Seller) error
	Delete(id int) error
}

type sellerRepository struct {
	db database.Database
}

func NewSellerRepository(db database.Database) SellerRepository {
	return &sellerRepository{db: db}
}

func (r *sellerRepository) Create(seller model.Seller) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO sellers (name, phone) VALUES ($1, $2) RETURNING id",
		seller.Name, seller.Phone).Scan(&id)
	return id, err
}

func (r *sellerRepository) GetByID(id int) (*model.Seller, error) {
	var seller model.Seller
	err := r.db.QueryRow(
		"SELECT id, name, phone FROM sellers WHERE id = $1",
		id).Scan(&seller.ID, &seller.Name, &seller.Phone)
	return &seller, err
}

func (r *sellerRepository) Update(seller model.Seller) error {
	_, err := r.db.Exec(
		"UPDATE sellers SET name = $1, phone = $2 WHERE id = $3",
		seller.Name, seller.Phone, seller.ID)
	return err
}

func (r *sellerRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM sellers WHERE id = $1", id)
	return err
}
