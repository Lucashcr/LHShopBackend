package products

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

func CreateProduct(name string, description string, price float64) (product *Product, err error) {
	if len(name) == 0 {
		log.Println("[ERROR]: Product name cannot be empty")
		return nil, errors.New("Product name cannot be empty")
	}

	if len(name) > 255 {
		log.Println("[ERROR]: Product name cannot be longer than 255 characters")
		return nil, errors.New("Product name cannot be longer than 255 characters")
	}

	if len(description) > 65535 {
		log.Println("[ERROR]: Product description cannot be longer than 65535 characters")
		return nil, errors.New("Product description cannot be longer than 65535 characters")
	}

	if price <= 0 {
		log.Println("[ERROR]: Product price must be greater than 0")
		return nil, errors.New("Product price must be greater than 0")
	}

	id := uuid.New()
	return &Product{id, name, description, price}, nil
}

type ProductsResponse struct {
	StatusCode   int       `json:"statusCode"`
	Result       []Product `json:"result"`
	TotalItems   int       `json:"totalItems"`
	TotalPages   int       `json:"totalPages"`
	ItemsPerPage int       `json:"itemsPerPage"`
	NextPage     string    `json:"nextPage,omitempty"`
	PrevPage     string    `json:"prevPage,omitempty"`
}
