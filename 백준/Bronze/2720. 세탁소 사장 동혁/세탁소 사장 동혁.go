package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
	목표: 거슬러 주어야 하는 금액은 항상 $5.00 이하이고, 최소 동전 개수로 거슬러 주기
	관련 알고리즘: 그리디 알고리즘

	사용할 동전:
		쿼터($0.25)
		다임($0.10)
		니켈($0.05)
		페니($0.01)
	
	알고리즘 설계
		1. 테스트 케이스 개수 입력받기
		2. 테스트 케이스 만큼 거스름 돈 입력받기
		3. 큰 단위부터 순서대로 나누어 필요한 동전 개수 얻기
		4. 0이 될때까지 반복하며 배열에 저장
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var t int

	fmt.Fscanln(r, &t)

	coins := []int{25, 10, 5, 1}
	coinCount := make([][]int, t)

	for i := 0; i < t; i++ {
		var amount int
		fmt.Fscanln(r, &amount)
		coinCount[i] = make([]int, len(coins))

		for j, coin := range coins {
			coinCount[i][j] = amount / coin
			amount %= coin
		}
	}

	for _, row := range coinCount {
		for _, value := range row {
			fmt.Fprint(w, value, " ")
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}
