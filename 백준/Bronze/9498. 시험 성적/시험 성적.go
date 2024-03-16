package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	if 90 <= a && a <= 100 {
		fmt.Println("A")
	} else if 80 <= a && a <= 89 {
		fmt.Println("B")
	} else if 70 <= a && a <= 79 {
		fmt.Println("C")
	} else if 60 <= a && a <= 69 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
