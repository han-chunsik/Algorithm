package main

import "fmt"

func main() {
	var a int

	fmt.Scanln(&a)

	var b, c int
	numbers := make([]int, a)

	for i := 0; i < a; i++ {
		fmt.Scanln(&b, &c)
		numbers[i] = b + c
	}

	for _, number := range numbers {
		fmt.Println(number)
	}

}
