package product

import (
	"errors"
	"regexp"
)

type Product struct {
	name  string
	price float64
}

func New(name string, price float64) (*Product, error) {
	if isNoValidName(name) {
		return nil, errors.New("Enter the product name!")
	}
	return &Product{name, price}, nil
}

func isNoValidName(value string) bool {
	compileNameRegexp := regexp.MustCompile(`^[a-zA-Zãâàáêèéìíîôóòùûú\s]+$`)
	isValid := compileNameRegexp.MatchString(value)
	return !isValid
}

func (product *Product) GetName() string   { return product.name }
func (product *Product) GetPrice() float64 { return product.price }
