package main

import "fmt"

func main() {
	var a, b int

	for {
		_, err := fmt.Scanln(&a, &b)
		if err != nil {
			break
		} else {
			fmt.Println(a + b)
		}
	}
}
