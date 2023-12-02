package main

import (
	"log"
	"strconv"
	"strings"
)

func day2(c chan string) {
	var total int64
	for line := range c {
		if len(line) == 0 {
			continue
		}

		var r, g, b int64
		gametag, roundString := splitString(line, ": ")
		gameId, err := strconv.ParseInt(gametag[5:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		rounds := strings.Split(roundString, "; ")
		for _, round := range rounds {
			plays := strings.Split(round, ", ")
			for _, play := range plays {
				numberString, color := splitString(play, " ")
				number, err := strconv.ParseInt(numberString, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				switch color {
				case "red":
					if number > r {
						r = number
					}
				case "green":
					if number > g {
						g = number
					}
				case "blue":
					if number > b {
						b = number
					}
				}
			}
		}
		if r <= 12 && g <= 13 && b <= 14 {
			log.Printf("Adding game: %d", gameId)
			total += gameId
		} else {
			log.Printf("Skipping game: %d", gameId)
		}

	}

	log.Printf("2A Total: %d", total)
}
