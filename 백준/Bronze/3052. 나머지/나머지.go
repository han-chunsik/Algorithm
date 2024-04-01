package main

import "fmt"

func main() {
	var a int
	nums := make([]int, 0, 10)

	resultList := []int{}

	for i := 0; i < 10; i++ {
		fmt.Scanln(&a)
		nums = append(nums, a%42)
	}

	for _, num := range nums {
		found := false
		for _, j := range resultList {
			if j == num {
				found = true
			}
		}
		if !found {
			resultList = append(resultList, num)
		}
	}
	fmt.Println(len(resultList))

}
