package order

import (
	"fmt"

	"github.com/Victor-armando18/shopping-cart/src/product"
)

type Product = *product.Product

type Item struct {
	product  Product
	price    float64
	amount   int
	discount float64
}

func NewItem(product Product, amount int, discount float64) *Item {
	return &Item{
		product:  product,
		price:    product.GetPrice(),
		amount:   amount,
		discount: discount,
	}
}

func (item *Item) GetName() string       { return item.product.GetName() }
func (item *Item) GetUnitPrice() float64 { return item.product.GetPrice() }
func (item *Item) GetAmount() int        { return item.amount }
func (item *Item) GetDiscount() float64  { return item.discount }
func (item *Item) GetTotalPrice() float64 {
	return (item.product.GetPrice() * float64(item.amount)) - item.discount
}
func (item *Item) ToString() string {
	return fmt.Sprintf("Item: %s -> { price: %.2f, amount: %d, discount: %.2f}\n",
		item.GetName(), item.GetUnitPrice(), item.GetAmount(), item.GetDiscount())
}
