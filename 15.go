package main

import (
	"log"
	"strings"
)

func day15(c chan string) {
	var total int

	for line := range c {
		if len(line) == 0 {
			continue
		}
		commands := strings.Split(line, ",")
		for _, command := range commands {
			value := 0
			for _, char := range command {
				value += int(char)
				value *= 17
				value = value % 256
			}
			total += value
		}
	}

	log.Printf("15A Total: %d", total)
}
