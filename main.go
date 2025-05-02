package main

import (
	config "backend/src/config/infrastructure"
	app "backend/src/product/application"
	"backend/src/product/domain"
	infra "backend/src/product/infrastructure"
	"log"
	"net/http"
)

func main() {
	// Set up the product repository, service, and controller
	productRepo := domain.NewMemoryRepository()
	productService := app.NewProductService(productRepo)
	productController := infra.NewProductController(productService)

	config.SetupRoutes(productController)

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
