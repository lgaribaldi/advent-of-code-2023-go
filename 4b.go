package main

import (
	"log"
	"strings"
)

func day4b(c chan string) {
	var total int
	// 206 is number of cards
	copies := make([]int, 206)
	lineNumber := -1
	for line := range c {
		if len(line) == 0 {
			continue
		}
		lineNumber++
		_, values := splitString(line, ": ")
		cardLeft, cardRight := splitString(values, " | ")
		winNumbers := nonEmptyValues(strings.Split(cardLeft, " "))
		cardNumbers := nonEmptyValues(strings.Split(cardRight, " "))
		matches := intersection(winNumbers, cardNumbers)
		if len(matches) > 0 {
			for i := 1; i <= len(matches) && lineNumber+i+1 <= len(copies); i++ {
				copies[lineNumber+i] += (copies[lineNumber] + 1)
			}
		}
	}
	for row, v := range copies {
		total += (v + 1)
		log.Printf("card: %d copies: %d", row+1, v+1)
	}

	log.Printf("4B Total: %d", total)
}
