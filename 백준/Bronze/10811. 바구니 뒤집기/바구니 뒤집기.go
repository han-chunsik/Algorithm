package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, m int

	fmt.Scanln(&n, &m)

	numList := make([]string, 0, n)
	resultList := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		numList = append(numList, strconv.Itoa(i))
		resultList = append(resultList, strconv.Itoa(i))
	}

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Scanln(&a, &b)
		for j, k := a-1, b-1; j <= b-1; j, k = j+1, k-1 {
			resultList[j] = numList[k]
		}
		copy(numList, resultList)
	}

	output := strings.Join(resultList, " ")
	fmt.Println(output)
}
