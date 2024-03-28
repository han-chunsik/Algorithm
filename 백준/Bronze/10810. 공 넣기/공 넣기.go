package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, m int

	fmt.Scanln(&n, &m)

	b := make([]string, n)

	for i := range b {
		b[i] = "0"
	}

	var i, j, k int

	for x := 0; x < m; x++ {
		fmt.Scanln(&i, &j, &k)

		for z := i; z < j+1; z++ {
			b[z-1] = strconv.Itoa(k)
		}
	}

	output := strings.Join(b, " ")
	fmt.Println(output)
}
