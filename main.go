package main

import (
	"log"
	"net/http"

	"github.com/Lucashcr/LHShopBackend/dbconn"
	"github.com/Lucashcr/LHShopBackend/middlewares"
	"github.com/Lucashcr/LHShopBackend/products"
)

func main() {
	dbconn.GetDB()
	defer dbconn.CloseDB()

	mux := http.NewServeMux()

	products.RegisterHandlers(mux)

	log.Println("[INFO]: Server running on port 8000...")
	defer log.Println("[INFO]: Server stopped!")
	err := http.ListenAndServe(":8000", middlewares.CorsMiddleware(mux))
	if err != nil {
		log.Fatalf("[ERROR]: Error starting server! %v", err)
	}
}
