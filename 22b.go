package main

import (
	"image"
	"log"
	"sort"
)

func day22b(c chan string) {

	total := 0
	supporters := make(map[int][]int)
	supporteds := make(map[int][]int)
	bricks := make(map[int]sandRectangle)
	falls := make(map[int][]int)
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
		supporteds[grid[i].id] = make([]int, 0)
		current := grid[i]
		bricks[current.id] = grid[i]
		landed := false
		// in air
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
					supporteds[current.id] = append(supporteds[current.id], set.id)
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

	fragilized := make(map[int][]int, 0)
	for _, brick := range settled {
		println(brick.id)
		falls[brick.id] = make([]int, 0)
		fragilized[brick.id] = make([]int, 0)

		for _, v := range supporters[brick.id] {
			if bricks[v].supports == 1 {
				falls[brick.id] = append(falls[brick.id], falls[bricks[v].id]...)
				falls[brick.id] = append(falls[brick.id], bricks[v].id)
				fragilized[brick.id] = append(fragilized[brick.id], fragilized[bricks[v].id]...)
			} else {
				fragilized[brick.id] = append(fragilized[brick.id], bricks[v].id)
			}
		}
		falls[brick.id] = append(falls[brick.id], brick.id)
		falls[brick.id] = unique(falls[brick.id])
		fragilized[brick.id] = unique(fragilized[brick.id])
		for {
			fallsSoFar := len(falls[brick.id])
			for _, v := range fragilized[brick.id] {
				if len(intersection(supporteds[v], falls[brick.id])) >= len(supporteds[v]) {
					falls[brick.id] = append(falls[brick.id], v)
					falls[brick.id] = append(falls[brick.id], falls[v]...)
					fragilized[brick.id] = append(fragilized[brick.id], fragilized[v]...)
				}
			}
			fragilized[brick.id] = unique(fragilized[brick.id])
			falls[brick.id] = unique(falls[brick.id])

			if len(falls[brick.id]) <= fallsSoFar {
				break
			}
		}
		falls[brick.id] = unique(falls[brick.id])
		// 1 means only brick itself
		if len(falls[brick.id]) > 1 {
			total += len(falls[brick.id]) - 1
		}
	}

	log.Printf("22B : %d", total)
}
