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

	croatia := []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}

	s, _ := r.ReadString('\n')
	s = strings.TrimSpace(s)

	replaced := s

	for _, i := range croatia {
		replaced = strings.Replace(replaced, i, "@", -1)
	}

	fmt.Fprintln(w, len(replaced))
	w.Flush()
}
