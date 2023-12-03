package main

import (
	"log"
	"strconv"
	"strings"
)

func day3b(c chan string) {
	var total int64

	blueprint := make([][]string, 0, 142)
	for line := range c {
		if len(line) == 0 {
			continue
		}
		blueprint = append(blueprint, strings.Split(line, ""))
	}

	for line := 0; line < len(blueprint); line++ {
		for col := 0; col < len(blueprint[line]); col++ {
			if blueprint[line][col] == "*" {
				adjacentNumbers := []int64{
					readCompleteNumber(&blueprint, line-1, col-1),
					readCompleteNumber(&blueprint, line-1, col),
					readCompleteNumber(&blueprint, line-1, col+1),
					readCompleteNumber(&blueprint, line, col+1),
					readCompleteNumber(&blueprint, line, col-1),
					readCompleteNumber(&blueprint, line+1, col-1),
					readCompleteNumber(&blueprint, line+1, col),
					readCompleteNumber(&blueprint, line+1, col+1),
				}
				adjacentNumbers = nonZeroValues(unique(adjacentNumbers))
				if len(adjacentNumbers) == 2 {
					log.Printf("Gear at: l%d c%d", line, col)
					log.Printf("Adjacent numbers: A%d B%d", adjacentNumbers[0], adjacentNumbers[1])
					total = total + (adjacentNumbers[0] * adjacentNumbers[1])
				}
			}
		}
	}
	log.Printf("3B Total: %d", total)
}

func readCompleteNumber(blueprint *[][]string, line int, col int) int64 {
	var currentNumber string
	if !isNumeric([]byte((*blueprint)[line][col])[0]) {
		return 0
	}
	startingPostion := col
	for ; isNumeric([]byte((*blueprint)[line][startingPostion])[0]); startingPostion-- {
	}
	startingPostion++
	for i := startingPostion; isNumeric([]byte((*blueprint)[line][i])[0]); i++ {
		currentNumber += (*blueprint)[line][i]
	}
	number, err := strconv.ParseInt(string(currentNumber), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func unique[T comparable](s []T) []T {
	var unique []T
	m := map[T]bool{}

	for _, v := range s {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

func nonZeroValues(s []int64) []int64 {
	var nonZero []int64
	for _, v := range s {
		if v != 0 {
			nonZero = append(nonZero, v)
		}
	}
	return nonZero
}
