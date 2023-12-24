package main

import (
	"image"
	"log"
)

func day23b(c chan string) {

	forestMap := make(map[image.Point]crosspath)
	input := make([]string, 0)

	for line := range c {
		if len(line) == 0 {
			continue
		}
		input = append(input, line)
	}

	for i := range input {
		for j := range input[i] {
			if input[i][j] == '.' {
				if i == 0 {
					forestMap[image.Point{i, j}] = crosspath{make([]path, 0), true, false}
					break
				}
				if i == len(input)-1 {
					forestMap[image.Point{i, j}] = crosspath{make([]path, 0), false, true}
					break
				}
				walls := 0
				if input[i+1][j] == '#' {
					walls++
				}
				if input[i-1][j] == '#' {
					walls++
				}
				if input[i][j+1] == '#' {
					walls++
				}
				if input[i][j-1] == '#' {
					walls++
				}
				if walls < 2 {
					forestMap[image.Point{i, j}] = crosspath{make([]path, 0), false, false}
				}
			}
		}
	}

	dir := []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for pos := range forestMap {
		for i := range dir {
			if (forestMap[pos].start && i != 0) || (forestMap[pos].end && i != 1) {
				continue
			}

			direction := dir[i]
			dist := 1
			position := pos.Add(direction)

			if input[position.X][position.Y] != '#' {
			walkUntilNextNode:
				for {
					if _, ok := forestMap[position]; ok {
						node := forestMap[pos]
						// reached other node
						node.paths = append(node.paths, path{position, dist})
						forestMap[pos] = node
						break
					}

					next := position.Add(direction)
					if input[next.X][next.Y] != '#' {
						position = next
						dist++
						continue
					} else {
						var testDirections []image.Point
						if direction.X != 0 {
							testDirections = dir[2:]
						} else {
							testDirections = dir[:2]
						}
						for _, tDir := range testDirections {
							next := position.Add(tDir)
							if input[next.X][next.Y] != '#' {
								direction = tDir
								position = next
								dist++
								continue walkUntilNextNode
							}
						}

					}
				}
			}
		}
	}
	var start image.Point
	for p := range forestMap {
		if forestMap[p].start {
			start = p
			break
		}
	}
	test, ok := longestPath(start, make(map[image.Point]struct{}), &forestMap)
	log.Printf("23B : %d %v", test, ok)
}
