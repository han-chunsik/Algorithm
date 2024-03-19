package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int

	fmt.Scanln(&a)

	var b, c int

	numList := make([]int, a)
	count := 1

	for i := 0; i < a; i++ {
		fmt.Scanln(&b, &c)
		numList[i] = b + c
	}

	for _, item := range numList {
		fmt.Println("Case #" + strconv.Itoa(count) + ": " + strconv.Itoa(item))
		count++
	}
}
