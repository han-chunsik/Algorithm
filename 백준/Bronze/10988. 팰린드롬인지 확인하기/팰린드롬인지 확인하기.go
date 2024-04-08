package main

import (
	"fmt"
)

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func main() {
	var s string
	fmt.Scan(&s)

	if s != reverseString(s) {
		fmt.Print(0)
	} else {
		fmt.Print(1)
	}

}
