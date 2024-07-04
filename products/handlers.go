package products

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Lucashcr/LHShopBackend/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	var newProduct product
	err = ParseProductFromRequest(r, &newProduct)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error parsing product from request: %s", r.URL, http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Error parsing product from request: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	err = insertProductIntoDatabase(&newProduct)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error inserting product into database: %s", r.URL, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Error inserting product into database: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	json, err := json.Marshal(newProduct)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error marshalling JSON: %s", r.URL, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Error marshalling JSON: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	log.Printf("[INFO] %s (%d): Product created!", r.URL, http.StatusCreated)
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	pageInt, err := utils.VerifyPageQuery(r, w)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error verifying page query: Invalid page number", r.URL, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error verifying page query: Invalid page number"))
		return
	}

	itemsPerPageInt, err := utils.VerifyItemsPerPageQuery(r, w)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error verifying itemsPerPage query: Invalid itemsPerPage number", r.URL, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error verifying itemsPerPage query: Invalid itemsPerPage number"))
		return
	}

	productsList, productsCount, err := fetchProductList(pageInt, itemsPerPageInt)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error fetching products list: ", r.URL, http.StatusInternalServerError)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error verifying itemsPerPage query: Invalid itemsPerPage number"))
		return
	}

	if pageInt < 1 || pageInt > productsCount/itemsPerPageInt+1 {
		log.Printf("[ERROR] %s (%d): Error verifying page query: Invalid page number", r.URL, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error verifying page query: Invalid page number"))
		return
	}

	prevPage, nextPage := utils.WritePrevAndNextPages(pageInt, productsCount, itemsPerPageInt)

	totalPages := utils.CalculateTotalPages(productsCount, itemsPerPageInt)

	productsResponse := productsResponse{
		Result:       productsList,
		TotalItems:   productsCount,
		TotalPages:   totalPages,
		ItemsPerPage: itemsPerPageInt,
		PrevPage:     prevPage,
		NextPage:     nextPage,
	}

	jsonProducts, err := json.Marshal(productsResponse)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error marshalling JSON: %s", r.URL, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling JSON"))
		return
	}

	log.Printf("[INFO] %s (%d): Products listed!", r.URL, http.StatusOK)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonProducts)
}

func DeatilProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	product, err := fetchProductByID(id)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error fetching product by ID: %s", r.URL, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error fetching product by ID: Product not found"))
		return
	}

	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error marshalling JSON: %s", r.URL, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Error marshalling JSON: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonProduct)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var updatedProduct product

	err = ParseProductFromRequest(r, &updatedProduct)

	if err != nil {
		log.Printf("[ERROR] %s (%d): Error parsing product from request: %s", r.URL, http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Error parsing product from request: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	rowsAffected, err := updateProductByID(&updatedProduct)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error updating product: %s", r.URL, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Error updating product: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	if rowsAffected == 0 {
		log.Printf("[ERROR] %s (%d): Error updating product: Product not found", r.URL, http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error updating product: Product not found"))
		return
	}

	json, err := json.Marshal(updatedProduct)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error marshalling JSON: %s", r.URL, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("[ERROR]: Error marshalling JSON: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	log.Printf("[INFO] %s (%d): Product updated!", r.URL, http.StatusOK)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	rowsAffected, err := deleteProductByID(id)
	if err != nil {
		log.Printf("[ERROR] %s (%d): Error deleting product: %s", r.URL, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("[ERROR]: Error deleting product: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	if rowsAffected == 0 {
		log.Printf("[ERROR] %s (%d): Error deleting product: Product not found", r.URL, http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error deleting product: Product not found"))
		return
	}

	log.Printf("[INFO] %s (%d): Product deleted!", r.URL, http.StatusNoContent)
	w.WriteHeader(http.StatusNoContent)
}
