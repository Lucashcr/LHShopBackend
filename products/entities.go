package products

import (
	"github.com/google/uuid"
)

type product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

type productsResponse struct {
	Result       []product `json:"result"`
	TotalItems   int       `json:"totalItems"`
	TotalPages   int       `json:"totalPages"`
	ItemsPerPage int       `json:"itemsPerPage"`
	NextPage     string    `json:"nextPage,omitempty"`
	PrevPage     string    `json:"prevPage,omitempty"`
}
