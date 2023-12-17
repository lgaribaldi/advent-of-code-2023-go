package main

import (
	"log"
)

func day16b(c chan string) {
	ch := make(chan int)

	grid := make([][]byte, 0, 0)
	total := 0
	jobs := 0
	for line := range c {
		if len(line) == 0 {
			continue
		}
		gridLine := []byte(line)
		grid = append(grid, gridLine)
	}
	for i := range grid {
		jobs += 2
		go checkEdge(i, 0, 'E', &grid, &ch)
		go checkEdge(i, len(grid[0])-1, 'W', &grid, &ch)
	}
	for i := range grid[0] {
		jobs += 2
		go checkEdge(0, i, 'S', &grid, &ch)
		go checkEdge(len(grid)-1, i, 'N', &grid, &ch)
	}
	for i := 0; i < jobs; i++ {
		if num := <-ch; num > total {
			total = num
		}
	}
	log.Printf("16B Most energized: %d", total)
}

func checkEdge(startX, startY int, direction rune, grid *[][]byte, ch *chan int) {
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
	*ch <- energized
}
