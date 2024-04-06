package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	for i := 97; i < 123; i++ {
		index := strings.Index(s, string(i))
		fmt.Printf("%d ", index)
	}
}