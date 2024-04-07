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

	strInput, _ := r.ReadString('\n')

	if strings.TrimSpace(strInput) == "" {
		fmt.Fprint(w, 0)
	} else {
		strSlice := strings.Split(strings.TrimSpace(strInput), " ")

		fmt.Fprint(w, len(strSlice))
	}

	w.Flush()
}
