package main

import (
	"log"
)

func day16(c chan string) {

	grid := make([][]byte, 0, 0)
	eGrid := make([][]int, 0, 0)
	total := 0
	for line := range c {
		if len(line) == 0 {
			continue
		}
		gridLine := []byte(line)
		grid = append(grid, gridLine)
		eGrid = append(eGrid, make([]int, len(gridLine)))
	}
	followBeamPath(0, 0, 'E', &grid, &eGrid)
	for i := range eGrid {
		println()
		for j := range eGrid[i] {
			print(eGrid[i][j])

			if eGrid[i][j] > 0 {
				total++
			}
		}
	}
	println()
	log.Printf("16A Energized: %d", total)
}

func followBeamPath(startX, startY int, direction rune, grid *[][]byte, eGrid *[][]int) {
	nextTileX := startX
	nextTileY := startY
	if (*eGrid)[nextTileX][nextTileY] > 1 {
		return
	}
	for {
		// out of bounds or passed here many times, must be a loop
		if nextTileX < 0 || nextTileX >= len(*grid) || nextTileY < 0 || nextTileY >= len((*grid)[0]) || (*eGrid)[nextTileX][nextTileY] > 8 {
			break
		}
		(*eGrid)[nextTileX][nextTileY] += 1
		switch direction {
		case 'E':
			switch (*grid)[nextTileX][nextTileY] {
			case '\\':
				nextTileX += 1
				direction = 'S'
			case '/':
				nextTileX -= 1
				direction = 'N'
			case '|':
				nextTileX -= 1
				direction = 'N'
				followBeamPath(nextTileX+1, nextTileY, 'S', grid, eGrid)
			default:
				nextTileY += 1
			}
		case 'N':
			switch (*grid)[nextTileX][nextTileY] {
			case '\\':
				nextTileY -= 1
				direction = 'W'
			case '/':
				nextTileY += 1
				direction = 'E'
			case '-':
				nextTileY -= 1
				direction = 'W'
				followBeamPath(nextTileX, nextTileY+1, 'E', grid, eGrid)
			default:
				nextTileX -= 1
			}
		case 'W':
			switch (*grid)[nextTileX][nextTileY] {
			case '\\':
				nextTileX -= 1
				direction = 'N'
			case '/':
				nextTileX += 1
				direction = 'S'
			case '|':
				nextTileX -= 1
				direction = 'N'
				followBeamPath(nextTileX+1, nextTileY, 'S', grid, eGrid)
			default:
				nextTileY -= 1
			}
		case 'S':
			switch (*grid)[nextTileX][nextTileY] {
			case '\\':
				nextTileY += 1
				direction = 'E'
			case '/':
				nextTileY -= 1
				direction = 'W'
			case '-':
				nextTileY -= 1
				direction = 'W'
				followBeamPath(nextTileX, nextTileY+1, 'E', grid, eGrid)
			default:
				nextTileX += 1
			}
		}
	}
}
