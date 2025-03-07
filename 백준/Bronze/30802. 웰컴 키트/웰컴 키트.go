package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
	* 티셔츠 종류: S, M, L, XL, XXL, XXXL
	* 조건:
	*	1. 티셔츠는 T장 묶음으로만 주문 가능
	*	2. 펜은 한 종류로 P자루 묶음으로 구매하거나, 한 자루씩 주문 가능
	*	3. N명의 참가자 중 S, M, L, XL, XXL, XXXL 사이즈의 티셔츠를 신청한 사람은 각각 S, M, L, XL, XXL, XXXL명
	*	4. 티셔츠는 남아도 되지만, 모자르면 안됨
	*	5. 펜은 정확한 숫자로 준비되야 함
	* 티셔츠를 T장씩 최소 몇 묶음 구매해야하는지, 펜을 P자루씩 최대 몇 묶음 주문, 이 때 펜을 한 자루씩 몇 개 주문하는지
	* 입력:
	* 	첫째 줄: 참가자 수 N
	*	둘째 줄: 사이즈별 신청자의 수가 S, M, L, XL, XXL, XXXL 순서로 공백으로 구분
	*	셋째 줄: 정수 티셔츠와 펜의 묶음 수를 의미하는 정수 T와 P가 공백으로 구분
	* 출력:
	*	첫째 줄: T장 씩 최소 묶음
	* 	둘째 줄: P자루씩 최대 묶음, 펜 한자루 씩 몇자루 공백으로 구분
	 */

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, t, p int
	sizes := make([]int, 6)

	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fscanf(r, "%d %d %d %d %d %d\n", &sizes[0], &sizes[1], &sizes[2], &sizes[3], &sizes[4], &sizes[5])
	fmt.Fscanf(r, "%d %d\n", &t, &p)

	tBundleCount := 0
	for _, size := range sizes {
		tBundleCount += size / t
		if size%t > 0 {
			tBundleCount++
		}
	}

	pBundleCount := n / p
	pCount := n % p

	fmt.Fprintln(w, tBundleCount)
	fmt.Fprintln(w, pBundleCount, pCount)

	w.Flush()
}
