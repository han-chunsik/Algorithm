package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	/*
	* 소수: 1과 자기 자신으로만 나눠지는 1보다 큰 자연수
	* 입력: 수의 개수 N
	* 출력: 주어진 수들 중 소수의 개수
	 */

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(r, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &nums[i])
	}

	cnt := 0

	for _, num := range nums {
		if num > 1 {
			isPrime := true
			for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
				if num%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				cnt++
			}
		}
	}
	fmt.Fprintln(w, cnt)

	w.Flush()
}
