package products

import "github.com/google/uuid"

type product struct {
	id          uuid.UUID
	name        string
	description string
	price       float64
}

func CreateProduct(name string, description string, price float64) product {
	if len(name) == 0 {
		panic("Product name cannot be empty")
	}

	if len(name) > 255 {
		panic("Product name cannot be longer than 100 characters")
	}

	if len(description) > 65535 {
		panic("Product description cannot be longer than 65535 characters")
	}

	if price <= 0 {
		panic("Product price must be greater than 0")
	}

	id := uuid.New()
	return product{id, name, description, price}
}

func (p product) GetName() string {
	return p.name
}

func (p product) GetPrice() float64 {
	return p.price
}

func (p product) GetDescription() string {
	return p.description
}
