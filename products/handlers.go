package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lucashcr/LHShopBackend/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	var newProduct Product
	err = ParseProductFromRequest(r, &newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Error parsing product from request: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	err = InsertProductIntoDatabase(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Error inserting product into database: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	json, err := json.Marshal(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Error marshalling JSON: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	pageInt, err := utils.VerifyPageQuery(r, w)
	if err != nil {
		return
	}

	itemsPerPageInt, err := utils.VerifyItemsPerPageQuery(r, w)
	if err != nil {
		return
	}

	productsList, productsCount := FetchProductList(pageInt, itemsPerPageInt)

	if pageInt < 1 || pageInt > productsCount/itemsPerPageInt+1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid page number"))
		return
	}

	prevPage, nextPage := utils.WritePrevAndNextPages(pageInt, productsCount, itemsPerPageInt)

	totalPages := utils.CalculateTotalPages(productsCount, itemsPerPageInt)

	productsResponse := ProductsResponse{
		Result:       productsList,
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

func DeatilProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	product, err := FetchProductByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
		return
	}

	jsonProduct, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonProduct)
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
