package handler

import (
	"TestTaskShop/internal/model"
	"TestTaskShop/internal/service"
	"TestTaskShop/pkg/authenticator"
	"TestTaskShop/pkg/validator"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type SellerHandler struct {
	Service service.SellerService
	auth    *authenticator.Authenticator
}

func NewSellerHandler(s service.SellerService, a *authenticator.Authenticator) *SellerHandler {
	return &SellerHandler{Service: s, auth: a}
}

func (h *SellerHandler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	if !h.auth.BasicAuth(w, r) {
		return
	}

	var seller model.Seller
	json.NewDecoder(r.Body).Decode(&seller)
	if err := validator.GetValidator().Validate(seller); err != nil {
		http.Error(w, fmt.Sprintf("invalid input values: %v", err), http.StatusBadRequest)
		return
	}

	id, err := h.Service.CreateSeller(seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *SellerHandler) GetSellerByID(w http.ResponseWriter, r *http.Request) {
	if !h.auth.BasicAuth(w, r) {
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	seller, err := h.Service.GetSellerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(seller)
}

func (h *SellerHandler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	if !h.auth.BasicAuth(w, r) {
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var seller model.Seller
	json.NewDecoder(r.Body).Decode(&seller)
	if err := validator.GetValidator().Validate(seller); err != nil {
		http.Error(w, fmt.Sprintf("invalid input values: %v", err), http.StatusBadRequest)
		return
	}
	seller.ID = id

	err := h.Service.UpdateSeller(seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SellerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	if !h.auth.BasicAuth(w, r) {
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := h.Service.DeleteSeller(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
