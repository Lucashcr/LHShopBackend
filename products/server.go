package products

import (
	"fmt"
	"net/http"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created"))
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("List of products"))
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	msg := fmt.Sprintf("Product with ID %s", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	msg := fmt.Sprintf("Product with ID %s updated", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	msg := fmt.Sprintf("Product with ID %s deleted", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /products", CreateProductHandler)
	mux.HandleFunc("GET /products", ListProductsHandler)
	mux.HandleFunc("GET /products/{id}", GetProductHandler)
	mux.HandleFunc("PUT /products/{id}", UpdateProductHandler)
	mux.HandleFunc("DELETE /products/{id}", DeleteProductHandler)
}
