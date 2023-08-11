package order

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Items = []Item

type Order struct {
	number int
	items  Items
}

func NewOrder(number int) (*Order, error) {
	if number <= 0 {
		return nil, errors.New("Invalid order number!")
	}
	return &Order{number, []Item{}}, nil
}

func (order *Order) Add(value *Item) *Order {
	order.items = append(order.items, *value)
	return order
}

func (order *Order) GetItems() Items { return order.items }

func (order *Order) GetTotalPrice() float64 {
	total := 0.0
	for _, item := range order.items {
		total += item.GetTotalPrice()
	}
	return total
}
func (order *Order) ToString() string {
	var result strings.Builder
	result.WriteString("Order: ")
	result.WriteString(strconv.Itoa(order.number))
	result.WriteString(" -> \n{\n")
	for _, item := range order.items {
		result.WriteString(item.ToString())
	}
	result.WriteString("} => Total Price: ")
	result.WriteString(fmt.Sprintf("%.2f", order.GetTotalPrice()))
	return result.String()
}
