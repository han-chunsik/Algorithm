package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
				"OOXXOXXOOO"의 점수는 1+2+0+0+1+0+0+1+2+3 = 10점
		        OX퀴즈의 결과가 주어졌을 때, 점수를 구하는 프로그램
				5
				OOXXOXXOOO
				OOXXOOXXOO
				OXOXOXOXOXOXOX
				OOOOOOOOOO
				OOOOXOOOOXOOOOX
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanf(r, "%d\n", &n)
	scoreList := make([]int, n)

	var tc string
	var t string
	for i := 0; i < n; i++ {
		cnt := 0
		fmt.Fscanf(r, "%s\n", &tc)
		for _, digit := range tc {
			t = string(digit)
			if t == "O" {
				cnt += 1
			} else {
				cnt = 0
			}
			scoreList[i] += cnt
		}
	}

	for _, score := range scoreList {
		fmt.Fprintln(w, score)
	}
	w.Flush()
}
