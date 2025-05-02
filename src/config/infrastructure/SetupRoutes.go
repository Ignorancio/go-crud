package infrastructure

import (
	"backend/src/product/infrastructure"
	"net/http"
)

func SetupRoutes(productController *infrastructure.ProductController) {
	// Product API endpoints
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productController.GetAllProducts(w, r)
		case "POST":
			productController.CreateProduct(w, r)
		case "PUT":
			productController.UpdateProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productController.GetProductById(w, r)
		case "DELETE":
			productController.DeleteProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
