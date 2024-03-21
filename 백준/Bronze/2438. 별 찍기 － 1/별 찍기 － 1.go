package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	fmt.Scan(&a)

	for i := 1; i < a+1; i++ {
		repeated := strings.Repeat("*", i)
		fmt.Println(repeated)
	}
}
