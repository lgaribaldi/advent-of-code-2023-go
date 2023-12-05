package main

import (
	"log"
	"math"
	"strings"
)

func day4(c chan string) {
	var total int64

	for line := range c {
		if len(line) == 0 {
			continue
		}
		_, values := splitString(line, ": ")
		cardLeft, cardRight := splitString(values, " | ")
		winNumbers := nonEmptyValues(strings.Split(cardLeft, " "))
		cardNumbers := nonEmptyValues(strings.Split(cardRight, " "))
		matches := intersection(winNumbers, cardNumbers)
		var cardScore float64
		if len(matches) > 0 {
			cardScore = math.Pow(2, float64(len(matches)-1))

			log.Printf(line+" Score: %f", cardScore)
			total += int64(cardScore)
		}
	}

	log.Printf("4A Total: %d", total)
}
