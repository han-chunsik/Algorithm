package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		A × B × C를 계산한 결과에 0부터 9까지 각각의 숫자가 몇 번씩 쓰였는지를 구하는 프로그램
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var a, b, c int
	fmt.Fscanf(r, "%d\n", &a)
	fmt.Fscanf(r, "%d\n", &b)
	fmt.Fscanf(r, "%d\n", &c)

	resultList := [10]int{}
	strNums := fmt.Sprintf("%d", a*b*c)

	for i := 0; i < 10; i++ {

		for _, digit := range strNums {
			strNum := string(digit)
			num, _ := strconv.Atoi(strNum)

			if num == i {
				resultList[i] += 1
			}
		}
	}

	for _, result := range resultList {
		fmt.Fprintln(w, result)
	}
	w.Flush()
}
