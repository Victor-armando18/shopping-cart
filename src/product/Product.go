package product

import "errors"

type Product struct {
	name string
}

func New(name string) (*Product, error) {
	if len(name) == 0 {
		return nil, errors.New("Enter the product name!")
	}
	return &Product{name}, nil
}

func (product *Product) GetName() string { return product.name }
