package main

import "fmt"

func main() {
	var n int

	scanList := make([]int, 0, 28)

	for j := 0; j < 28; j++ {
		fmt.Scanln(&n)
		scanList = append(scanList, n)
	}
	resultList := make([]int, 0, 2)

	for i := 1; i <= 30; i++ {
		found := false
		for _, m := range scanList {
			if i == m {
				found = true
			}
		}
		if !found {
			resultList = append(resultList, i)
		}
	}

	max := resultList[0]

	min := resultList[1]

	for _, b := range resultList {
		if max < b {
			max = b
		}
		if min > b {
			min = b
		}
	}

	fmt.Println(min)
	fmt.Println(max)
}
