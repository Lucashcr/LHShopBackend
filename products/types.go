package products

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

func ParseProductFromRequest(r *http.Request, p *Product) error {
	var err error

	err = json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		return err
	}

	err = p.validate()
	if err != nil {
		return err
	}

	return nil
}

func (p *Product) validate() error {
	if p.Name == "" {
		return errors.New("Product name is required")
	}

	if p.Description == "" {
		return errors.New("Product description is required")
	}

	if p.Price == 0 {
		return errors.New("Product price is required")
	}

	return nil
}

type ProductsResponse struct {
	Result       []Product `json:"result"`
	TotalItems   int       `json:"totalItems"`
	TotalPages   int       `json:"totalPages"`
	ItemsPerPage int       `json:"itemsPerPage"`
	NextPage     string    `json:"nextPage,omitempty"`
	PrevPage     string    `json:"prevPage,omitempty"`
}
