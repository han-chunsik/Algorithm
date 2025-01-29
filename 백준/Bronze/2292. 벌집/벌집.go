package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		중앙 1부터 N 사이의 방의 개수(시작, 끝 포함)
		- 각 숫자가 몇번째 줄에 위치하는지 구해야함.
			1번 줄: 1
			2번 줄: 2~7(6개)
			3번 줄: 8~19(12개)
			4번 줄: 20~37(18개)
			5번 줄: 38~61(24개)
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanf(r, "%d", &n)

	m := 1
	i := 1
	for m < n {
		m += i * 6
		i++
	}

	fmt.Fprintln(w, i)
	w.Flush()
}
