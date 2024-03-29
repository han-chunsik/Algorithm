package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	buckets := make([]string, 0, n)

	fmt.Scanln(&n, &m)

	for x := 1; x < n+1; x++ {
		buckets = append(buckets, strconv.Itoa(x))
	}

	var i, j int

	for z := 0; z < m; z++ {
		fmt.Scanln(&i, &j)
		buckets[i-1], buckets[j-1] = buckets[j-1], buckets[i-1]
	}

	output := strings.Join(buckets, " ")
	fmt.Println(output)
}
