package main

import "fmt"

func main() {
	var a, b int

	for i := 0; i < 1; {
		fmt.Scanln(&a, &b)
		if a == 0 && b == 0 {
			i += 1
		} else {
			fmt.Println(a + b)
		}
	}
}
