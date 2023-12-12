package main

import (
	"log"
	"math"
	"strings"
)

type Galaxy struct {
	x, y int
}

func day11(c chan string) {

	var total int
	starMap := make([][]byte, 0, 0)
	for line := range c {
		if len(line) == 0 {
			continue
		}
		starLine := []byte(line)
		if !strings.Contains(line, "#") {
			starMap = append(starMap, starLine)
		}
		starMap = append(starMap, starLine)
	}

	log.Println("Old map")
	for x := range starMap {
		log.Println(string(starMap[x]))
	}

	var newCol bool
column:
	for y := 0; ; y++ {
		if y == len(starMap[0]) {
			break
		}
		if newCol {
			newCol = false
			continue
		}

		for x := range starMap {
			if starMap[x][y] == '#' {
				continue column
			}
		}
		// duplicate column
		log.Println("Duplicating col", y)
		for x := range starMap {
			newLine := append(starMap[x], '.')
			copy(newLine[y+1:], newLine[y:])
			newLine[y+1] = '.'
			starMap[x] = newLine
		}
		newCol = true
	}

	log.Println("New map")
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
			distance := int(math.Abs(float64(galaxy.x-other.x)) + math.Abs(float64(galaxy.y-other.y)))
			log.Println("Distance ", i+1, j+i+1+1, distance)
			total += distance
		}
	}

	log.Printf("11A Total: %d", total)
}
