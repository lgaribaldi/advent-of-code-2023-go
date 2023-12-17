package main

import (
	"log"
)

func day16b(c chan string) {

	grid := make([][]byte, 0, 0)
	total := 0
	for line := range c {
		if len(line) == 0 {
			continue
		}
		gridLine := []byte(line)
		grid = append(grid, gridLine)
	}

	for i := range grid {
		if numE := checkEdge(i, 0, 'E', &grid); numE > total {
			total = numE
		}
		if numW := checkEdge(i, len(grid[0])-1, 'W', &grid); numW > total {
			total = numW
		}
	}
	for i := range grid[0] {
		if numS := checkEdge(0, i, 'S', &grid); numS > total {
			total = numS
		}
		if numN := checkEdge(len(grid)-1, i, 'N', &grid); numN > total {
			total = numN
		}
	}
	log.Printf("16B Most energized: %d", total)
}

func checkEdge(startX, startY int, direction rune, grid *[][]byte) int {
	eGrid := make([][]int, len(*grid))
	for i := range eGrid {
		eGrid[i] = make([]int, len((*grid)[0]))
	}
	energized := 0
	followBeamPath(startX, startY, direction, grid, &eGrid)
	for i := range eGrid {
		for j := range eGrid[i] {
			if eGrid[i][j] > 0 {
				energized++
			}
		}
	}
	println(startX, startY, direction, energized)
	return energized
}
