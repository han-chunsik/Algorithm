package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		정문에서 가장 짧은 거리에 있는 방
		w개의 방이 있는 h층 건물
		앨리베이터는 가장 왼쪽에 있으며, 정문도 거기에 가깝게 있음, 정문과 앨리베이터의 거리는 고려하지 않음
	*/

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanf(r, "%d\n", &t)

	resultList := make([]int, t)

	var H, W, N int
	var f, room int
	for i := 0; i < t; i++ {
		fmt.Fscanf(r, "%d %d %d\n", &H, &W, &N)

		f = N % H
		room = N/H + 1
		if f == 0 {
			room -= 1
			f += H
		}

		resultList[i] = f*100 + room
	}

	for _, result := range resultList {
		fmt.Fprintln(w, result)
	}

	w.Flush()
}
