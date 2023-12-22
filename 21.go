package main

import (
	"image"
	"log"
	"strings"
)

func day21(c chan string) {

	grid := make(map[image.Point]int)
	lineNumber := 0
	var startingPosition image.Point
	for line := range c {
		if len(line) == 0 {
			continue
		}
		lineNumber++
		positions := strings.Split(line, "")
		for i := range positions {
			position := image.Point{lineNumber, i}
			switch positions[i] {
			case "S":
				grid[position] = 9999
				startingPosition = image.Point{lineNumber, i}
			case ".":
				grid[position] = 9999
			case "#":
				grid[position] = -1
			}
		}
	}

	floodFillWalk(startingPosition, 0, &grid)

	steps := 0
	for i := 0; i < 131; i++ {
		println()
		for j := 0; j < 131; j++ {
			if v, ok := grid[image.Point{i, j}]; ok {
				if v < 0 {
					print("  # ")
				} else if v > 9 {
					print(" ", v, " ")
				} else {
					print("  ", v, " ")
				}
				if v < 65 && v%2 == 0 {
					steps++
				}

			}
		}
		println()
	}

	log.Printf("21A : %d", steps)
}

func floodFillWalk(pos image.Point, dist int, grid *map[image.Point]int) {
	curr, ok := (*grid)[pos]
	if !ok || curr == -1 || dist > 64 {
		return
	}
	if dist < curr {
		(*grid)[pos] = dist
		floodFillWalk(pos.Add(image.Point{0, 1}), dist+1, grid)
		floodFillWalk(pos.Add(image.Point{0, -1}), dist+1, grid)
		floodFillWalk(pos.Add(image.Point{1, 0}), dist+1, grid)
		floodFillWalk(pos.Add(image.Point{-1, 0}), dist+1, grid)
	}

}
