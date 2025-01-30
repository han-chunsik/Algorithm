package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		A, B, C 가 주어질 때,
		첫째 줄에는 정수로 계산한 A + B - C
		둘째 줄에는 문자열로 계산한 A + B - C

		예시)
		A = 3
		B = 4
		C = 5

		첫째 줄: (3 + 4) - 5 = 2
		둘째 줄: "3" + "4" - "5" = "34" - "5" = 29
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var a, b, c int
	fmt.Fscanf(r, "%d\n", &a)
	fmt.Fscanf(r, "%d\n", &b)
	fmt.Fscanf(r, "%d\n", &c)

	resultInt := (a + b) - c

	sumStr := strconv.Itoa(a) + strconv.Itoa(b)
	sumInt, _ := strconv.Atoi(sumStr)
	resultStr := sumInt - c

	fmt.Fprintln(w, resultInt)
	fmt.Fprintln(w, resultStr)
	w.Flush()
}
