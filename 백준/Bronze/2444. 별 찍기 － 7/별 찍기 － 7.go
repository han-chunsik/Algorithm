package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		blank := strings.Repeat(" ", n-i)
		repeated := strings.Repeat("*", 2*i-1)
		fmt.Println(blank + repeated)
	}
	for j := n - 1; j > 0; j-- {
		blank := strings.Repeat(" ", n-j)
		repeated := strings.Repeat("*", 2*j-1)
		fmt.Println(blank + repeated)
	}
}
