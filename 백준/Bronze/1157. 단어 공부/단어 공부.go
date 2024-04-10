package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

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
		fmt.Println("?")
	} else {
		fmt.Printf("%c\n", maxIndex)
	}
}
