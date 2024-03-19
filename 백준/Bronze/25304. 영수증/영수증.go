package main

import (
	"fmt"
)

func main() {
	var total, kind int
	var price, amount int

	fmt.Scanln(&total)
	fmt.Scanln(&kind)

	ProdPriceList := make([]int, kind)

	for i := 0; i < kind; i++ {
		fmt.Scanln(&price, &amount)
		ProdPriceList[i] = price * amount
	}

	sum := 0

	for _, ProdPrice := range ProdPriceList {
		sum += ProdPrice
	}

	if sum == total {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}

}
