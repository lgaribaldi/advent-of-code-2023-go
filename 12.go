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
			//log.Println("Annotation: %d possibilities: ", v, nextPossibilities)
			possibilities = nextPossibilities
			nextPossibilities = make([]string, 0, 0)
			for _, currentSprings := range possibilities {
				for i := range currentSprings {
					// TODO refactor these ifs into one
					//can start in current position

					// skips chained fails
					if i > 0 && currentSprings[i-1] == '#' {
						break
					}
					//log.Printf("current string: %s v: %d i: %d", currentSprings, v, i)

					if currentSprings[i] == '#' || currentSprings[i] == '?' {
						// number of adjacent failures and ? fits the annotation
						if len(currentSprings[i:]) >= int(v) && strings.Index(currentSprings[i:i+int(v)], ".") == -1 {
							// next one is not another failures
							if len(currentSprings) == i+int(v) || currentSprings[i+int(v)] != '#' {
								// possibility for next annotation

								if len(currentSprings) < i+int(v)+2 {
									//log.Printf("possibility: %s", currentSprings[i:])
									nextPossibilities = append(nextPossibilities, "")
								} else {
									//log.Println("len(currentSprings)", len(currentSprings), "i+int(v)+2", i+int(v)+2)

									//log.Printf("possibility: %s", currentSprings[i+int(v)+1:])
									nextPossibilities = append(nextPossibilities, string(currentSprings[i+int(v)+1:]))
								}
							} else {
								//log.Println("else 3")
							}
						} else {
							//log.Println("else 2")
						}
					} else {
						//log.Println("else 1")
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
	// answer is 6949
	log.Printf("12A Total: %d", total)

}
