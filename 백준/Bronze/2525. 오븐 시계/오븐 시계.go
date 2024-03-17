package main

import (
	"fmt"
)

func main() {
	var h, m, t int

	fmt.Scanln(&h, &m)
	fmt.Scanln(&t)

	fmt.Print((h+(m+t)/60)%24, (m+t)%60)

}
