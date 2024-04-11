package main

import (
	"log"
	"net/http"

	"github.com/Lucashcr/LHShopBackend/dbconn"
	"github.com/Lucashcr/LHShopBackend/middlewares"
	"github.com/Lucashcr/LHShopBackend/products"
)

func main() {
	db := dbconn.GetDB()
	defer db.Close()

	mux := http.NewServeMux()

	products.RegisterHandlers(mux)

	log.Println("[INFO]: Server running on port 8000...")
	err := http.ListenAndServe(":8000", middlewares.CorsMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
