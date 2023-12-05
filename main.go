package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categoryController"
	"go-web-native/controllers/homeController"
	"go-web-native/controllers/productController"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	// 1.homepage
	http.HandleFunc("/", homeController.Welcome)

	// 2.categories
	http.HandleFunc("/categories", categoryController.Index)
	http.HandleFunc("/categories/add", categoryController.Add)
	http.HandleFunc("/categories/edit", categoryController.Edit)
	http.HandleFunc("/categories/delete", categoryController.Delete)

	// Product
	http.HandleFunc("/products", productController.Index)
	http.HandleFunc("/products/add", productController.Add)
	http.HandleFunc("/products/detail", productController.Detail)
	http.HandleFunc("/products/edit", productController.Edit)
	http.HandleFunc("/products/delete", productController.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8090", nil)
}
