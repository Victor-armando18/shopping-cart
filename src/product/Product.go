package product

import (
	"errors"
	"regexp"
)

type Product struct {
	name  string
	price float64
}

func NewProduct(name string, price float64) (*Product, error) {
	if len(name) == 0 {
		return nil, errors.New("Enter the product name!")
	}
	return &Product{name, price}, nil
}

func isNoValidName(value string) bool {
	compileNameRegexp := regexp.MustCompile(`^[a-zA-Zãâàáêèéìíîôóòùûú\s]+[0-9]*$`)
	isValid := compileNameRegexp.MatchString(value)
	return !isValid
}

func (product *Product) GetName() string   { return product.name }
func (product *Product) GetPrice() float64 { return product.price }
