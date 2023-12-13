package main

import (
	"log"
	"strings"
)

func day12(c chan string) {
	var total int

	for line := range c {
		if len(line) == 0 {
			continue
		}
		springs, annotationsPart := splitString(line, " ")
		annotations := readLineOfNumbers(annotationsPart, ",")

		possibilities := make([]string, 0, 0)
		nextPossibilities := make([]string, 0, 0)
		nextPossibilities = append(possibilities, springs)

		for _, v := range annotations {
			possibilities = nextPossibilities
			nextPossibilities = make([]string, 0, 0)
			for _, currentSprings := range possibilities {

				for i := range currentSprings {
					// Assigns annotation to current failed spring
					if i > 0 && currentSprings[i-1] == '#' {
						break
					}

					if (currentSprings[i] == '#' || currentSprings[i] == '?') &&
						len(currentSprings[i:]) >= int(v) && strings.Index(currentSprings[i:i+int(v)], ".") == -1 &&
						(len(currentSprings) == i+int(v) || currentSprings[i+int(v)] != '#') {
						// possibility for next annotation
						if len(currentSprings) < i+int(v)+2 {
							nextPossibilities = append(nextPossibilities, "")
						} else {
							nextPossibilities = append(nextPossibilities, string(currentSprings[i+int(v)+1:]))
						}
					}
				}

			}

		}

		// filter out the ones who left damaged springs
		finalPossibilities := make([]string, 0, 0)
		for _, v := range nextPossibilities {
			if strings.Index(v, "#") == -1 {
				finalPossibilities = append(finalPossibilities, v)
			}
		}
		log.Printf("line: %s matches: %d", line, len(finalPossibilities))
		total += len(finalPossibilities)
	}

	log.Printf("12A Total: %d", total)

}
