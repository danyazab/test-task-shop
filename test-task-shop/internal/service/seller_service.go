package service

import (
	"TestTaskShop/internal/model"
	"TestTaskShop/internal/repository"
)

type SellerService interface {
	CreateSeller(seller model.Seller) (int, error)
	GetSellerByID(id int) (*model.Seller, error)
	UpdateSeller(seller model.Seller) error
	DeleteSeller(id int) error
}

type sellerService struct {
	repo repository.SellerRepository
}

func NewSellerService(repo repository.SellerRepository) SellerService {
	return &sellerService{repo: repo}
}

func (s *sellerService) CreateSeller(seller model.Seller) (int, error) {
	return s.repo.Create(seller)
}

func (s *sellerService) GetSellerByID(id int) (*model.Seller, error) {
	return s.repo.GetByID(id)
}

func (s *sellerService) UpdateSeller(seller model.Seller) error {
	return s.repo.Update(seller)
}

func (s *sellerService) DeleteSeller(id int) error {
	return s.repo.Delete(id)
}
