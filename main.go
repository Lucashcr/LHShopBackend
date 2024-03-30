package main

import (
	"fmt"
	"net/http"

	"github.com/Lucashcr/LHShopBackend/products"
)

func main() {
	mux := http.NewServeMux()

	products.RegisterHandlers(mux)

	fmt.Println("Server running on port 8000...")
	http.ListenAndServe(":8000", mux)
}
