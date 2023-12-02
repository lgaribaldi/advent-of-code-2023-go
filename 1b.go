package main

import (
	"log"
	"strconv"
	"strings"
)

func day1b(c chan string) {
	total := 0
	for line := range c {
		if len(line) == 0 {
			continue
		}
		line = replaceNumbers(line)

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

	log.Printf("1B Total: %d", total)

}

func replaceNumbers(line string) string {
	newLine := line
	newLine = strings.Replace(newLine, "one", "o1e", -1)
	newLine = strings.Replace(newLine, "two", "t2o", -1)
	newLine = strings.Replace(newLine, "three", "t3e", -1)
	newLine = strings.Replace(newLine, "four", "f4r", -1)
	newLine = strings.Replace(newLine, "five", "f5e", -1)
	newLine = strings.Replace(newLine, "six", "s6x", -1)
	newLine = strings.Replace(newLine, "seven", "s7n", -1)
	newLine = strings.Replace(newLine, "eight", "e8t", -1)
	newLine = strings.Replace(newLine, "nine", "n9e", -1)

	return newLine
}
