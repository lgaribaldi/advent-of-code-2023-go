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

		annotations := make([]int, 0, 0)
		for i := 0; i < 3; i++ {
			annotations = append(annotations, preAnnotations...)
		}
		springs := preSprings + "?" + preSprings + "?" + preSprings + "." // + "?" + preSprings + "?" + preSprings

		possibilities := make([]string, 0, 0)
		nextPossibilities := make([]string, 0, 0)
		nextPossibilities = append(possibilities, springs)

		//TODO store a second map for bigger chunks

		for _, v := range annotations {
			fmt.Println(len(nextPossibilities))
			possibilities = nextPossibilities
			nextPossibilities = make([]string, 0, 0)
			for _, currentSprings := range possibilities {
				for i := 0; i+v < len(currentSprings); i++ {
					// nextPossibilities = append(nextPossibilities, springs[i:i+v+1])

					//for j := range next {
					if doTheThing(v, currentSprings[i:i+v+1]) {
						nextPossibilities = append(nextPossibilities, currentSprings[i+v+1:])
						//next[j] += currentSprings[i+annotations[0]+1:]
						//println("Next", next[j])
					}
					if currentSprings[i] == '#' {
						break
					}

				}
				//nextPossibilities = append(nextPossibilities, doTheThing(v, currentSprings)...)
			}
		}

		// filter out the ones who left damaged springs
		finalPossibilities := 0
		for _, v := range nextPossibilities {
			if strings.Index(v, "#") == -1 {
				finalPossibilities++
			}
		}
		log.Printf("line: %s matches: %d", line, finalPossibilities)
		total += finalPossibilities
	}

	log.Println(memo)
	log.Printf("12A Total: %d", total)

}

func doTheThing(v int, currentSprings string) bool {
	index := fmt.Sprint(v) + currentSprings
	if elem, ok := memo[index]; ok {
		return elem
	}

	//answer := make([]string, 0, 0)
	// var start int
	/*for i := range currentSprings {
		if currentSprings[i] != '.' {
			start = i
			break
		}
	}*/
	//currentSprings = currentSprings[start:]
	//for i := range currentSprings {
	// Assigns annotation to current failed spring
	/*if i > 0 && currentSprings[i-1] == '#' {
		break
	}*/

	if (currentSprings[0] == '#' || currentSprings[0] == '?') &&
		strings.Index(currentSprings[:int(v)], ".") == -1 &&
		(currentSprings[int(v)] != '#') {
		// possibility for next annotation
		//if len(currentSprings) < int(v)+2 {
		//answer = append(answer, "")
		memo[index] = true
		return true
		/*} else {
			println("Acho que nao cai mais aqui agora")
			// answer = append(answer, string(currentSprings[int(v)+1:]))
		}*/
	}
	//}
	memo[index] = false
	return false

	//return answer
}
