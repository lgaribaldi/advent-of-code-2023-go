package main

import (
	"fmt"
	"log"
	"strings"
)

var memo map[string]bool

func day12b(c chan string) {
	var total int
	memo = make(map[string]bool)

	for line := range c {
		if len(line) == 0 {
			continue
		}
		preSprings, annotationsPart := splitString(line, " ")
		preAnnotations := make([]int, 0, 0)
		for _, v := range readLineOfNumbers(annotationsPart, ",") {
			preAnnotations = append(preAnnotations, int(v))
		}
		// Multiplies input by 5
		annotations := make([]int, 0, 0)
		for i := 0; i < 5; i++ {
			annotations = append(annotations, preAnnotations...)
		}
		springs := preSprings + "?" + preSprings + "?" + preSprings + "?" + preSprings + "?" + preSprings + "."

		possibilities := make(map[string]int)
		nextPossibilities := make(map[string]int)
		nextPossibilities[springs] = 1

		for _, v := range annotations {
			possibilities = nextPossibilities
			nextPossibilities = make(map[string]int)
			for currentSprings, ammount := range possibilities {
				// Cuts input into the smalles possible size
				for i := 0; i+v < len(currentSprings); i++ {
					if checkSpringMatch(v, currentSprings[i:i+v+1]) {
						nextPossibilities[currentSprings[i+v+1:]] += ammount
					}
					if currentSprings[i] == '#' {
						break
					}
				}
			}
		}

		// filter out the ones who left damaged springs
		finalPossibilities := 0
		for v, ammount := range nextPossibilities {
			if strings.Index(v, "#") == -1 {
				finalPossibilities += ammount
			}
		}
		total += finalPossibilities
	}
	log.Printf("12B Total: %d", total)

}

func checkSpringMatch(v int, currentSprings string) bool {
	index := fmt.Sprint(v) + currentSprings
	if elem, ok := memo[index]; ok {
		return elem
	}

	if (currentSprings[0] == '#' || currentSprings[0] == '?') &&
		strings.Index(currentSprings[:int(v)], ".") == -1 &&
		(currentSprings[int(v)] != '#') {
		memo[index] = true
		return true
	}
	memo[index] = false
	return false
}
