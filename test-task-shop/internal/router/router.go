package router

import (
	"TestTaskShop/internal/handler"
	"net/http"
)

func NewRouter(sellerHandler *handler.SellerHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// routes for sellers
	mux.HandleFunc("/sellers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			sellerHandler.CreateSeller(w, r)
		case http.MethodGet:
			sellerHandler.GetSellerByID(w, r)
		case http.MethodPut:
			sellerHandler.UpdateSeller(w, r)
		case http.MethodDelete:
			sellerHandler.DeleteSeller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
