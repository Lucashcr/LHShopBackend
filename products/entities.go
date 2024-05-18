package products

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

func ParseProductFromRequest(r *http.Request, p *product) error {
	var err error

	if r.Method != http.MethodPost {
		id := r.PathValue("id")
		if id == "" {
			return errors.New("product ID is required")
		}

		p.ID, err = uuid.Parse(id)
		if err != nil {
			return errors.New("product ID is not a valid UUID")
		}
	}

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

func (p *product) validate() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}

	if p.Description == "" {
		return errors.New("product description is required")
	}

	if p.Price == 0 {
		return errors.New("product price is required")
	}

	return nil
}

type productsResponse struct {
	Result       []product `json:"result"`
	TotalItems   int       `json:"totalItems"`
	TotalPages   int       `json:"totalPages"`
	ItemsPerPage int       `json:"itemsPerPage"`
	NextPage     string    `json:"nextPage,omitempty"`
	PrevPage     string    `json:"prevPage,omitempty"`
}
