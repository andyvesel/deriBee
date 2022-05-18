package grid

import "fmt"

func SetupOrders() {
	order := 10
	multiplier := 2
	orders := []int{}
	dep := 200
	lev := dep / 2
	for i := order; i <= lev; i += order {
		order = order + multiplier
		orders = append(orders, order)
	}
	result := 0
	for _, order := range orders {
		result += order
		fmt.Printf("%d\n", result)
	}

	fmt.Println(orders)
	fmt.Println(result - order)
}
