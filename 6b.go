package main

import (
	"log"
	"strconv"
	"strings"
)

func day6b(c chan string) {
	var time, distance, result int64
	for line := range c {
		if len(line) == 0 {
			continue
		}

		tag, valueString := splitString(line, ":")
		values := strings.ReplaceAll(valueString, " ", "")
		value, err := strconv.ParseInt(values, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if tag == "Time" {
			time = value
		} else {
			distance = value
		}
	}

	for y := 1; int64(y) < time; y++ {
		if int64(y)*(time-int64(y)) > distance {
			result += 1
		}
	}

	log.Printf("6B Result: %d", result)
}
