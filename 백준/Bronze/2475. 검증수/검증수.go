package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	/*
		5개의 숫자를 각각 제곱한 수의 합을 10으로 나눈 나머지
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var a, b, c, d, e float64

	fmt.Fscanf(r, "%f %f %f %f %f", &a, &b, &c, &d, &e)

	sumNum := (math.Pow(a, 2) + math.Pow(b, 2) + math.Pow(c, 2) + math.Pow(d, 2) + math.Pow(e, 2))

	result := int(sumNum) % 10

	fmt.Fprintln(w, result)
	w.Flush()
}
