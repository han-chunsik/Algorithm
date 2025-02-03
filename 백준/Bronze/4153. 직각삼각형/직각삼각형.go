package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	/*
		주어진 각 변 a, b, c로 직삼각형인지 확인하는 방법
		- a제곱 + b제곱 = c제곱(가장 긴 수) 일 경우 직삼각형
		입력 마지막이 0 0 0 일 경우 입력 종료
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var a, b, c int

	for {
		fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)
		if a+b+c == 0 {
			break
		}

		powA := math.Pow(float64(a), 2)
		powB := math.Pow(float64(b), 2)
		powC := math.Pow(float64(c), 2)

		result := "wrong"
		if a > b && a > c {
			if powA-powB-powC == 0 {
				result = "right"
			}
		}

		if b > a && b > c {
			if powB-powA-powC == 0 {
				result = "right"
			}
		}

		if c > a && c > b {
			if powC-powA-powB == 0 {
				result = "right"
			}
		}

		fmt.Fprintln(w, result)
	}

	w.Flush()
}
