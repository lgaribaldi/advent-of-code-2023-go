package main

import (
	"log"
)

func day13b(c chan string) {
	var total int
	var pattern []string
	lineN := 0

	for line := range c {
		lineN++
		if len(line) == 0 {
			// horizontal
			h := checkHorizontalReflectionWithSmudges(pattern) * 100
			// vertical
			invertedPattern := make([]string, 0, 0)
			for i := 0; i < len(pattern[0]); i++ {
				row := make([]byte, 0, 0)
				for j := 0; j < len(pattern); j++ {
					row = append(row, pattern[j][i])
				}
				invertedPattern = append(invertedPattern, string(row))
			}
			v := checkHorizontalReflectionWithSmudges(invertedPattern)

			total += h + v
			pattern = nil
			continue
		}
		pattern = append(pattern, line)
	}

	log.Printf("13B Total: %d", total)

}

func checkHorizontalReflectionWithSmudges(pattern []string) int {
	for i := 1; i < len(pattern); i++ {
		smudges := 0
		ok, sm := smudgiMatch(pattern[i], pattern[i-1])
		if ok {
			smudges = sm
			couldThisBeTheOne := true
			for j := i + 1; j < len(pattern); j++ {
				if i-(j-i+1) >= 0 {
					ok, sm := smudgiMatch(pattern[j], pattern[i-(j-i+1)])
					smudges += sm
					if !ok || smudges > 1 {
						couldThisBeTheOne = false
						break
					}
				}
			}
			if couldThisBeTheOne && smudges == 1 {
				return i
			}
		}
	}
	return 0
}

func smudgiMatch(v1, v2 string) (bool, int) {
	smudges := 0
	for i := range v1 {
		if v1[i] != v2[i] {
			smudges++
			if smudges > 1 {
				return false, 0
			}
		}
	}
	return true, smudges
}
