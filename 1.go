package main

import (
	"log"
	"strconv"
)

func day1(c chan string) {
	total := 0
	for line := range c {
		if len(line) == 0 {
			continue
		}

		var first, last byte
		for i := 0; i < len(line) && (first == 0 || last == 0); i++ {
			if first == 0 && isNumeric(line[i]) {
				first = line[i]
			}
			if last == 0 && isNumeric(line[len(line)-1-i]) {
				last = line[len(line)-1-i]
			}
		}
		lineTotal, err := strconv.ParseInt(string(first)+string(last), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Line: %s", line)
		log.Println(lineTotal)
		total += int(lineTotal)
	}

	log.Printf("1A Total: %d", total)
}
