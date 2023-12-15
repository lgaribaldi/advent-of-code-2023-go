package main

import (
	"log"
	"strconv"
	"strings"
)

type lens struct {
	label  string
	length int
}

func day15b(c chan string) {
	var total int
	boxes := make([][]lens, 256, 256)

	for line := range c {
		if len(line) == 0 {
			continue
		}
		commands := strings.Split(line, ",")
		for _, command := range commands {
			value := 0
			var operation rune
			var length int
			var label string
			for i, char := range command {
				if char == '-' || char == '=' {
					operation = char
					label = command[:i]
					if i+1 < len(command) {
						if num, err := strconv.ParseInt(command[i+1:], 10, 32); err == nil {
							length = int(num)
						}
					}
					break
				}
				value += int(char)
				value *= 17
				value = value % 256
			}

			if operation == '-' {
				boxes[value] = removeLens(boxes[value], label)
			}
			if operation == '=' {
				index := findLensIndex(boxes[value], label)
				if index > -1 {
					boxes[value][index] = lens{label, length}
				} else {
					boxes[value] = append(boxes[value], lens{label, length})
				}
			}
		}
	}
	for i, box := range boxes {
		println("Box:", i)
		for j := range box {
			println(box[j].label, box[j].length)
			total += (i + 1) * (j + 1) * box[j].length
		}
	}

	log.Printf("15B Total: %d", total)
}

func removeLens(list []lens, label string) []lens {
	result := make([]lens, 0)
	for _, lens := range list {
		if lens.label != label {
			result = append(result, lens)
		}
	}
	return result
}

func findLensIndex(list []lens, label string) int {
	idx := -1
	for i, lens := range list {
		if lens.label == label {
			idx = i
			break
		}
	}
	return idx
}
