package internal

import (
	"fmt"
	"log"
	"net/http"

	"TestTaskShop/internal/configs"
	"TestTaskShop/internal/database"
	migrate "TestTaskShop/internal/database/migration"
	"TestTaskShop/internal/handler"
	"TestTaskShop/internal/repository"
	"TestTaskShop/internal/router"
	"TestTaskShop/internal/service"
	"TestTaskShop/pkg/authenticator"
)

func RunApi() error {
	cfg, err := configs.Load()
	if err != nil {
		return fmt.Errorf("can not load config: %e", err)
	}

	db, err := database.NewPostgreSQLDB(cfg)
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}
	defer db.Close() // close the connection to the database when the work is completed

	if err = migrate.RunMigrations(db); err != nil {
		return fmt.Errorf("could not migrate database: %w", err)
	}

	sellerRepo := repository.NewSellerRepository(db)
	sellerService := service.NewSellerService(sellerRepo)

	auth := authenticator.NewAuthenticator(
		cfg.Admin1,
		cfg.Admin2,
	)
	sellerHandler := handler.NewSellerHandler(sellerService, auth)

	r := router.NewRouter(sellerHandler)
	log.Println(fmt.Sprintf("Server is running on port %d", cfg.Api.Port))

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Api.Port), r)
}
