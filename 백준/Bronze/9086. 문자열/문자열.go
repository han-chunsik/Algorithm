package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var s string
	resultList := []string{}

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		start := string(s[0])
		end := string(s[len(s)-1])
		resultList = append(resultList, start+end)
	}

	for _, result := range resultList {
		fmt.Println(result)
	}

}
