package main

import (
	"fmt"
	"strconv"
)

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var n, m string
	fmt.Scan(&n, &m)

	intN, _ := strconv.Atoi(reverseString(n))
	intM, _ := strconv.Atoi(reverseString(m))

	if intN > intM {
		fmt.Print(intN)
	} else {
		fmt.Print(intM)
	}
}
