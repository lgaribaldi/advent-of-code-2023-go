package main

import (
	"image"
	"log"
	"strings"
)

func day21b(c chan string) {

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
			position := image.Point{lineNumber - 1, i}
			switch positions[i] {
			case "S":
				grid[position] = 9999
				startingPosition = image.Point{lineNumber - 1, i}
			case ".":
				grid[position] = 9999
			case "#":
				grid[position] = -1
			}
		}
	}

	floodFillWalk2(startingPosition, 0, &grid)

	// gets total walkable tiles on even and odd grids, total rocks and unreachable spots
	var even, odd, evenCorner, oddCorner, rocks, unreachable int
	for i := 0; i < 131; i++ {
		for j := 0; j < 131; j++ {
			if v, ok := grid[image.Point{i, j}]; ok {
				if v > 65 {
					if v%2 == 0 {
						evenCorner++
					} else if v != 9999 {
						oddCorner++
					}
				}
				if v%2 == 0 {
					even++
				} else if v != 9999 && v != -1 {
					odd++
				}
				if v == -1 {
					rocks++
				} else if v == 9999 {
					unreachable++
				}

			}
		}
	}
	// Sums the total of odd and even grids considering a square of 202300 grids (26501365 steps - 65 from middle grid / 131 (width adn height of each grid))
	// Adds even corners and subtracts missing odd corners for the outer layer
	total := (202301)*(202301)*odd + 202300*202300*even - (202301 * oddCorner) + 202300*evenCorner
	log.Printf("\nOdds: %d\nEven: %d\nOdd corners: %d\nEven corners: %d\n", odd, even, oddCorner, evenCorner)
	log.Print("Unreachable: ", unreachable, "\n")
	log.Print("Odd + Even + rocks + unreachable: ", odd+even+rocks+unreachable)

	log.Printf("21B : %d", total)
}

func floodFillWalk2(pos image.Point, dist int, grid *map[image.Point]int) {
	curr, ok := (*grid)[pos]
	if !ok || curr == -1 {
		return
	}
	if dist < curr {
		(*grid)[pos] = dist
		floodFillWalk2(pos.Add(image.Point{0, 1}), dist+1, grid)
		floodFillWalk2(pos.Add(image.Point{0, -1}), dist+1, grid)
		floodFillWalk2(pos.Add(image.Point{1, 0}), dist+1, grid)
		floodFillWalk2(pos.Add(image.Point{-1, 0}), dist+1, grid)
	}
}
