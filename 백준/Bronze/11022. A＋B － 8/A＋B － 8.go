package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int

	fmt.Scanln(&a)

	rtList := make([]string, a)

	var b, c int
	for i := 0; i < a; i++ {
		fmt.Scanln(&b, &c)
		result := "Case #" + strconv.Itoa(i+1) + ": " + strconv.Itoa(b) + " + " + strconv.Itoa(c) + " = " + strconv.Itoa(b+c)
		rtList[i] = result
	}

	for _, item := range rtList {
		fmt.Println(item)
	}
}
