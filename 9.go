package main

import (
	"log"
	"strconv"
	"strings"
)

func day9(c chan string) {
	var total int

	sequence := make([][]int, 0, 0)
	// read file
	for line := range c {
		if len(line) == 0 {
			continue
		}
		numbersStr := strings.Split(line, " ")
		numbers := make([]int, 0, 0)
		for _, v := range numbersStr {
			number, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, int(number))
		}
		sequence = append(sequence, numbers)
	}

	// generate subSequences for each line
	for j := 0; j < len(sequence); j++ {
		subSequence := make([][]int, 0, 0)
		subLevel := 0
		subSequence = append(subSequence, sequence[j])
		for {
			nextLevel := make([]int, 0, 0)
			prevValue := subSequence[subLevel][0]
			for i := 1; i < len(subSequence[subLevel]); i++ {
				nextLevel = append(nextLevel, subSequence[subLevel][i]-prevValue)
				prevValue = subSequence[subLevel][i]
			}
			subSequence = append(subSequence, nextLevel)
			subLevel++
			shouldBreak := true
			for _, v := range nextLevel {
				if v != 0 {
					shouldBreak = false
					break
				}
			}
			if shouldBreak {
				break
			}
		}

		// calculate next value
		nextValue := 0
		for level, ss := range subSequence {
			log.Printf("SubSequence level %d", level)
			for i, v := range ss {
				if i+1 == len(ss) {
					nextValue += v
				}
				log.Printf("%d", v)
			}
		}
		log.Printf("Next value for sequence %d", nextValue)
		total += nextValue

	}

	log.Printf("9A Total: %d", total)

}
