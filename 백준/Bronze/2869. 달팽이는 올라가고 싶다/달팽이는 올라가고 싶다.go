package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b, v int
	fmt.Fscanf(r, "%d %d %d\n", &a, &b, &v)

	// 달팽이가 매일 올라가는 높이 a-b
	// 마지막 날에는 미끄러지지 않기 때문에 달팽이가 올라가야 하는 높이는 v-b
	// => 마지막 날에는 a를 모두 올라가지만, 계산의 편의성을 위해 v-b를 사용하는 것이 적합
	day := int((v - b) / (a - b))

	// 정상에서는 미끄러지지 않으니, 나머지 만큼 마지막으로 올라가야하니까 하루 추가
	if (v-b)%(a-b) != 0 {
		day++
	}

	fmt.Fprintln(w, day)
}
