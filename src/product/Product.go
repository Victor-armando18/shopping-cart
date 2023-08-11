package product

import (
	"errors"
	"regexp"
)

type Product struct {
	name string
}

func New(name string) (*Product, error) {
	if isNoValidName(name) {
		return nil, errors.New("Enter the product name!")
	}
	return &Product{name}, nil
}

func isNoValidName(value string) bool {
	compileNameRegexp := regexp.MustCompile(`^[a-zA-Zãâàáêèéìíîôóòùûú\s]+$`)
	isValid := compileNameRegexp.MatchString(value)
	return !isValid
}

func (product *Product) GetName() string { return product.name }
