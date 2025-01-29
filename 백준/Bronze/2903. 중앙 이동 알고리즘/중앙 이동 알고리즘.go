package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	/*
		점의 개수
			초기 정사각형: 가로 2개, 세로 2개
				2*2 = 4개
			1번 변형: 가로 3개, 세로 3개
				3*3 = 9개

			- 변형 될 때마다 점의 개수는 2의 N제곱 + 1 로 늘어남
			- 총 점의 개수는 가로 축에있는 점의 개수 * 세로 축에있는 점의 개수 이므로, 은 결국 2의 N제곱 +1 을 제곱한 값
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n float64
	fmt.Fscanf(r, "%f", &n)

	d := int(math.Pow((math.Pow(2, n) + 1), 2))

	fmt.Fprintln(w, d)
	w.Flush()
}
