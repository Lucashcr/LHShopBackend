package products

import "errors"

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
