package main

import (
	"log"
	"strconv"
	"strings"
)

func day3(c chan string) {
	var total int64

	blueprint := make([][]string, 0, 142)
	for line := range c {
		if len(line) == 0 {
			continue
		}
		blueprint = append(blueprint, strings.Split(line, ""))
	}

	var isPart = false
	var currentNumber = ""
	for line := 0; line < len(blueprint); line++ {
		currentNumber = ""
		isPart = false
		for col := 0; col < len(blueprint[line]); col++ {
			if isNumeric([]byte(blueprint[line][col])[0]) {
				currentNumber += blueprint[line][col]
				if !isPart && ((!isNumeric([]byte(blueprint[line+1][col])[0]) && blueprint[line+1][col] != ".") ||
					(!isNumeric([]byte(blueprint[line+1][col+1])[0]) && blueprint[line+1][col+1] != ".") ||
					(!isNumeric([]byte(blueprint[line+1][col-1])[0]) && blueprint[line+1][col-1] != ".") ||
					(!isNumeric([]byte(blueprint[line][col-1])[0]) && blueprint[line][col-1] != ".") ||
					(!isNumeric([]byte(blueprint[line][col+1])[0]) && blueprint[line][col+1] != ".") ||
					(!isNumeric([]byte(blueprint[line-1][col])[0]) && blueprint[line-1][col] != ".") ||
					(!isNumeric([]byte(blueprint[line-1][col+1])[0]) && blueprint[line-1][col+1] != ".") ||
					(!isNumeric([]byte(blueprint[line-1][col-1])[0]) && blueprint[line-1][col-1] != ".")) {
					isPart = true
				}
			} else {
				if isPart {
					number, err := strconv.ParseInt(string(currentNumber), 10, 64)
					if err != nil {
						log.Fatal(err)
					}
					total += number
				}
				currentNumber = ""
				isPart = false
			}
		}
	}
	log.Printf("3A Total: %d", total)
}
