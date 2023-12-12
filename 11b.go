package main

import (
	"log"
	"strings"
)

func day11b(c chan string) {

	const EXPANSION_VALUE = 1000000
	var total int
	starMap := make([][]byte, 0, 0)
	expandingCols := make([]int, 0, 0)
	expandingRows := make([]int, 0, 0)
	currentRow := 0
	for line := range c {
		if len(line) == 0 {
			continue
		}
		starLine := []byte(line)
		if !strings.Contains(line, "#") {
			expandingRows = append(expandingRows, currentRow)
		}
		currentRow++
		starMap = append(starMap, starLine)
	}

column:
	for y := range starMap[0] {
		for x := range starMap {
			if starMap[x][y] == '#' {
				continue column
			}
		}
		log.Println("Expanded col", y)
		expandingCols = append(expandingCols, y)
	}

	var galaxies []Galaxy
	for x := range starMap {
		log.Println(string(starMap[x]))
		for y := range starMap[x] {
			if starMap[x][y] == '#' {
				galaxies = append(galaxies, Galaxy{x, y})
			}
		}
	}

	for i, galaxy := range galaxies {
		for j, other := range galaxies[i+1:] {
			distance := 0
			var smallerY, biggerY int
			if galaxy.y >= other.y {
				smallerY = other.y
				biggerY = galaxy.y
			} else {
				smallerY = galaxy.y
				biggerY = other.y
			}
			distance += biggerY - smallerY
			for _, v := range expandingCols {
				if v >= smallerY && v <= biggerY {
					distance += EXPANSION_VALUE - 1
				}
			}

			distance += other.x - galaxy.x
			for _, v := range expandingRows {
				if v >= galaxy.x && v <= other.x {
					distance += EXPANSION_VALUE - 1
				}
			}

			log.Println("Distance ", i+1, j+i+1+1, distance)
			total += distance
		}
	}

	log.Printf("11B Total: %d", total)
}
