package main

import (
	"fmt"
	"log"
	"strings"
)

var memo map[string][]int

func day12b(c chan string) {
	var total int
	memo = make(map[string][]int)

	for line := range c {
		if len(line) == 0 {
			continue
		}
		preSprings, annotationsPart := splitString(line, " ")
		preAnnotations := readLineOfNumbers(annotationsPart, ",")

		annotations := make([]int64, 0, 0)
		for i := 0; i < 5; i++ {
			annotations = append(annotations, preAnnotations...)
		}
		springs := preSprings + "?" + preSprings + "?" + preSprings + "?" + preSprings + "?" + preSprings

		possibilities := make([]int, 0, 0)
		nextPossibilities := make([]int, 0, 0)
		nextPossibilities = append(possibilities, len(springs))

		for _, v := range annotations {
			possibilities = nextPossibilities
			nextPossibilities = make([]int, 0, 0)
			for _, currentSprings := range possibilities {
				nextPossibilities = append(nextPossibilities, doTheThing(v, springs[len(springs)-currentSprings:])...)
			}
		}

		// filter out the ones who left damaged springs
		finalPossibilities := 0
		for _, v := range nextPossibilities {
			if strings.Index(springs[len(springs)-v:], "#") == -1 {
				finalPossibilities++
			}
		}
		log.Printf("line: %s matches: %d", line, finalPossibilities)
		total += finalPossibilities
	}

	log.Println(memo)
	log.Printf("12A Total: %d", total)

}

func doTheThing(v int64, currentSprings string) []int {
	index := fmt.Sprint(v) + currentSprings
	if elem, ok := memo[index]; ok {
		return elem
	}

	answer := make([]int, 0, 0)
	var start int
	for i := range currentSprings {
		if currentSprings[i] != '.' {
			start = i
			break
		}
	}
	currentSprings = currentSprings[start:]
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
				answer = append(answer, 0)
			} else {
				answer = append(answer, len(string(currentSprings[i+int(v)+1:])))
			}
		}
	}
	memo[index] = answer
	return answer
}
