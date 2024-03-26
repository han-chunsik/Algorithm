package main

import "fmt"

func main() {
	var a int
	numList := make([]int, 0, 9)

	for i := 0; i < 9; i++ {
		fmt.Scanln(&a)
		numList = append(numList, a)
	}

	max := 0
	index := 0

	for i, num := range numList {
		if max < num {
			max = num
			index = i + 1
		}
	}
	fmt.Println(max)
	fmt.Println(index)
}
