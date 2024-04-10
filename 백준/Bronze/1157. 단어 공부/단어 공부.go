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

	s, _ := r.ReadString('\n')

	s = strings.ToUpper(strings.TrimSpace(s))

	freqMap := make(map[byte]int)
	var maxFreq int

	for i := 0; i < len(s); i++ {
		char := s[i]
		freqMap[char]++
		if freqMap[char] > maxFreq {
			maxFreq = freqMap[char]
		}
	}

	var maxCount, maxIndex int
	for char, freq := range freqMap {
		if freq == maxFreq {
			maxCount++
			maxIndex = int(char)
		}
	}

	if maxCount > 1 {
		fmt.Fprintln(w, "?")
	} else {
		fmt.Fprintf(w, "%c\n", maxIndex)
	}
	w.Flush()
}
