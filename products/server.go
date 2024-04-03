package products

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created"))
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid page number"))
		return
	}

	itemsPerPage := r.URL.Query().Get("itemsPerPage")
	if itemsPerPage == "" {
		itemsPerPage = "30"
	}

	itemsPerPageInt, err := strconv.Atoi(itemsPerPage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid itemsPerPage number"))
		return
	}

	products, productsCount := GetProductList(pageInt, itemsPerPageInt)

	if pageInt < 1 || pageInt > productsCount/itemsPerPageInt+1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid page number"))
		return
	}

	var prevPage string
	if pageInt > 1 {
		prevPage = fmt.Sprintf("/products?page=%d", pageInt-1)
	}

	var nextPage string
	if productsCount > pageInt*itemsPerPageInt {
		nextPage = fmt.Sprintf("/products?page=%d", pageInt+1)
	}

	if itemsPerPageInt != 30 {
		prevPage = fmt.Sprintf("%s&itemsPerPage=%d", prevPage, itemsPerPageInt)
		nextPage = fmt.Sprintf("%s&itemsPerPage=%d", nextPage, itemsPerPageInt)
	}

	totalPages := productsCount / itemsPerPageInt
	if productsCount%itemsPerPageInt != 0 {
		totalPages++
	}

	productsResponse := ProductsResponse{
		Result:       products,
		TotalItems:   productsCount,
		TotalPages:   totalPages,
		ItemsPerPage: itemsPerPageInt,
		PrevPage:     prevPage,
		NextPage:     nextPage,
	}

	jsonProducts, err := json.Marshal(productsResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonProducts)
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

	log.Println("[INFO]: Products handlers registered!")
}
