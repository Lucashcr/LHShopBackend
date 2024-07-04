package products

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

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
