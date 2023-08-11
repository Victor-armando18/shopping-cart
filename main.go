package main

import (
	"fmt"

	"github.com/Victor-armando18/shopping-cart/src/order"
	"github.com/Victor-armando18/shopping-cart/src/product"
)

func main() {
	newOrder, _ := order.NewOrder(1)
	firstItem, _ := product.NewProduct("Recarga DSTV Mensal", 7600)
	secondItem, _ := product.NewProduct("Net Casa UNITEL 10MB/S - 30 Dias", 15000)
	newOrder.Add(order.NewItem(firstItem, 2, 0))
	newOrder.Add(order.NewItem(secondItem, 2, 0))
	fmt.Print(newOrder.ToString())
}
