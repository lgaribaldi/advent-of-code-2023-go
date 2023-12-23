package main

import (
	"image"
	"log"
	"sort"
)

type sandRectangle struct {
	rect                image.Rectangle
	id, z, z2, supports int
}

func day22(c chan string) {

	total := 0
	supporters := make(map[int][]int)
	bricks := make(map[int]sandRectangle)
	grid := make([]sandRectangle, 0)
	settled := make([]sandRectangle, 0)
	lineNumber := 0
	for line := range c {
		if len(line) == 0 {
			continue
		}
		lineNumber++
		start, end := splitString(line, "~")
		startPos := readLineOfNumbers(start, ",")
		endPos := readLineOfNumbers(end, ",")
		grid = append(grid, sandRectangle{image.Rectangle{image.Point{int(startPos[0]), int(startPos[1])}, image.Point{int(endPos[0] + 1), int(endPos[1] + 1)}}, lineNumber, int(startPos[2]), int(endPos[2]), 0})
	}
	sort.SliceStable(grid, func(i, j int) bool {
		return grid[i].z < grid[j].z
	})

	for i := range grid {
		supporters[grid[i].id] = make([]int, 0)
		current := grid[i]
		bricks[current.id] = grid[i]
		landed := false
		if grid[i].z != 1 {
			for _, set := range settled {
				if set.z2 >= current.z {
					continue
				}
				if set.rect.Overlaps(current.rect) {
					if !landed {
						landed = true
						current.z2 = current.z2 - current.z + set.z2 + 1
						current.z = set.z2 + 1
					}
					if set.z2+1 < current.z {
						continue
					}
					current.supports++
					supporters[set.id] = append(supporters[set.id], current.id)
					grid[i] = current
					bricks[current.id] = current
				}
			}
		} else {
			landed = true // ground level
		}
		if !landed {
			current.z2 = current.z2 - current.z + 1
			current.z = 1
			grid[i] = current
			bricks[current.id] = current
		}
		settled = append(settled, current)
		sort.SliceStable(settled, func(i, j int) bool {
			return settled[i].z2 > settled[j].z2
		})
	}

	for _, brick := range bricks {
		canDisintegrate := true
		for _, v := range supporters[brick.id] {
			if bricks[v].supports == 1 {
				canDisintegrate = false
				break
			}
		}
		if canDisintegrate {
			println("Disintegrating ", brick.id)
			total++
		}
	}

	log.Printf("22A : %d", total)
}
