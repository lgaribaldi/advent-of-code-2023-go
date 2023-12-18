package main

import (
	"image"
	"log"
	"strconv"
	"strings"
)

func day18(c chan string) {

	grid := map[image.Point]struct{}{}
	nextPosition := image.Point{0, 0}
	grid[nextPosition] = struct{}{}
	var size int64
	size = 1
	for line := range c {
		if len(line) == 0 {
			continue
		}
		params := strings.Split(line, " ")
		direction := rune(params[0][0])
		length, err := strconv.ParseInt(params[1], 10, 32)
		if err != nil {
			log.Fatal("wrong input")
		}
		//color is useless so far
		nextPosition = digTrench(nextPosition, direction, int(length), &grid, &size)
	}

	floodFill(image.Point{1, 1}, &grid, &size)

	for i := -5; i < 120; i++ {
		println()
		for j := -5; j < 150; j++ {
			if _, ok := grid[image.Point{i, j}]; ok {
				print("#")
			} else {
				print(".")
			}
		}
	}
	log.Printf("18A Tiles: %d", size)
}

func floodFill(pos image.Point, grid *map[image.Point]struct{}, size *int64) {
	if _, ok := (*grid)[pos]; ok {
		return
	}
	(*grid)[pos] = struct{}{}
	(*size)++
	floodFill(pos.Add(image.Point{0, 1}), grid, size)
	floodFill(pos.Add(image.Point{0, -1}), grid, size)
	floodFill(pos.Add(image.Point{1, 0}), grid, size)
	floodFill(pos.Add(image.Point{-1, 0}), grid, size)
}

func digTrench(start image.Point, direction rune, length int, grid *map[image.Point]struct{}, size *int64) image.Point {

	var directionPoint, lastPoint image.Point
	switch direction {
	case 'R':
		directionPoint = image.Point{0, 1}
	case 'L':
		directionPoint = image.Point{0, -1}
	case 'U':
		directionPoint = image.Point{-1, 0}
	case 'D':
		directionPoint = image.Point{1, 0}
	}

	for i := 1; i <= length; i++ {
		vector := directionPoint.Mul(i)
		lastPoint = start.Add(vector)
		if _, ok := (*grid)[lastPoint]; ok {
			continue
		}
		(*grid)[lastPoint] = struct{}{}
		(*size)++
	}
	return lastPoint
}
