package main

import (
	"log"
)

func day13(c chan string) {
	var total int
	var pattern []string

	for line := range c {
		if len(line) == 0 {
			// horizontal
			h := checkHorizontalReflection(pattern) * 100
			// vertical
			invertedPattern := make([]string, 0, 0)
			for i := 0; i < len(pattern[0]); i++ {
				row := make([]byte, 0, 0)
				for j := 0; j < len(pattern); j++ {
					row = append(row, pattern[j][i])
				}
				invertedPattern = append(invertedPattern, string(row))
			}
			v := checkHorizontalReflection(invertedPattern)

			total += h + v
			pattern = nil
			continue
		}
		pattern = append(pattern, line)
	}

	log.Printf("13A Total: %d", total)

}

func checkHorizontalReflection(pattern []string) int {
	for i := 1; i < len(pattern); i++ {
		if pattern[i] == pattern[i-1] {
			couldThisBeTheOne := true
			for j := i + 1; j < len(pattern); j++ {
				if i-(j-i+1) >= 0 && pattern[j] != pattern[i-(j-i+1)] {
					couldThisBeTheOne = false
					break
				}
			}
			if couldThisBeTheOne {
				return i
			}
		}
	}
	return 0
}
