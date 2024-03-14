package main

import (
	"fmt"
	"strconv"
)

// https://www.educative.io/answers/how-to-reverse-a-string-in-golang
func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	numbersStr := strconv.Itoa(b)
	reverseStr := reverseString(numbersStr)
	for _, numStr := range reverseStr {
		fmt.Println(a * int(numStr-'0'))
	}

	fmt.Println(a * b)

}
