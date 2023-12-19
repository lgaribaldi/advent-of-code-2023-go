package main

import (
	"image"
	"log"
	"math"
	"strconv"
	"strings"
)

func day18b(c chan string) {

	grid := map[image.Point]struct{}{}
	nextPosition := image.Point{0, 0}
	grid[nextPosition] = struct{}{}
	var perimiter int64
	perimiter = 1
	vertices := make([]image.Point, 0)
	vertices = append(vertices, nextPosition)
	for line := range c {
		if len(line) == 0 {
			continue
		}
		params := strings.Split(line, " ")
		direction := 'R'
		switch params[2][7:8] {
		case "1":
			direction = 'D'
		case "2":
			direction = 'L'
		case "3":
			direction = 'U'
		}
		length, err := strconv.ParseInt(params[2][2:7], 16, 32)
		println(length, string(direction))
		if err != nil {
			log.Fatal("wrong input")
		}

		nextPosition = digTrench(nextPosition, direction, int(length), &grid, &perimiter)
		vertices = append(vertices, nextPosition)
	}

	// use Shoelace formula to calculate the innerArea of the polygon
	innerArea := shoelace(&vertices)

	// uses Pick's theorem with the perimenter and inner area to find total area
	// one way to think is that shoelace give's the area using the midpoint of the boundary
	// no we add the other half of the perimiter
	log.Printf("18B Tiles: %f", innerArea+float64(perimiter)/2+1)
}

func shoelace(vertices *[]image.Point) float64 {
	innerArea := 0
	for i := 0; i < len((*vertices))-1; i++ {
		innerArea += (*vertices)[i].X*(*vertices)[i+1].Y - (*vertices)[i+1].X*(*vertices)[i].Y
	}
	innerArea += (*vertices)[len(*vertices)-1].X*(*vertices)[0].Y - (*vertices)[0].X*(*vertices)[len(*vertices)-1].Y
	return math.Abs(float64(innerArea)) / 2.0
}
