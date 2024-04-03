package products

import (
	"log"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /products", CreateProductHandler)
	mux.HandleFunc("GET /products", ListProductsHandler)
	mux.HandleFunc("GET /products/{id}", DeatilProductHandler)
	mux.HandleFunc("PUT /products/{id}", UpdateProductHandler)
	mux.HandleFunc("DELETE /products/{id}", DeleteProductHandler)

	log.Println("[INFO]: Products handlers registered!")
}
