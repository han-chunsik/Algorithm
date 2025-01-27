package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
	10진수 값과 N을 입력받는다.
	10진수를 N으로 나누어 몫과 나머지를 구한다.
	나머지를 기록한다.
	몫이 0이 될 때까지 2단계와 3단계를 반복한다.
	기록된 나머지를 거꾸로 조합하여 N진수 결과를 얻는다.
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	nums := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var n int
	var b int

	fmt.Fscanf(r, "%d %d", &n, &b)

	result := ""

	for n > 0 {
		rem := n % b
		result = string(nums[rem]) + result
		n /= b
	}

	fmt.Fprintln(w, result)

	w.Flush()
}
