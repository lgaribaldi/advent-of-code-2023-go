package main

import (
	"log"
	"strconv"
	"strings"
)

func day6(c chan string) {
	var total int
	var time, distance, result []int
	for line := range c {
		if len(line) == 0 {
			continue
		}

		tag, valueString := splitString(line, ":")
		values := strings.Split(valueString, " ")
		var numericValues []int
		for _, v := range values {
			if v == "" {
				continue
			}
			value, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			numericValues = append(numericValues, int(value))
		}
		if tag == "Time" {
			time = numericValues
		} else {
			distance = numericValues
		}
	}

	for i := range time {
		result = append(result, 0)
		for y := 1; y < time[i]; y++ {
			if y*(time[i]-y) > distance[i] {
				result[i] += 1
			}
		}
	}

	total = 1
	for _, v := range result {
		total *= v
	}

	log.Printf("6A Total: %d", total)
}
