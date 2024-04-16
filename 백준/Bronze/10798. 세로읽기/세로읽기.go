package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	ws := make([]string, 0)
	for i := 0; i < 5; i++ {
		strInput, _ := r.ReadString('\n')
		w := strings.TrimSpace(strInput)
		ws = append(ws, w)
	}

	result := ""

	for i := 0; i < 15; i++ {
		for j := 0; j < 5; j++ {
			if len(ws[j]) > i {
				result += string(ws[j][i])
			}
		}
	}

	fmt.Fprint(w, result)
	w.Flush()
}
