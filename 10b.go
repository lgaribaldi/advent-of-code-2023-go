package main

import (
	"bytes"
	"log"
)

/*
To solve this I choose to raycast from each tile until the bottom of the map
If the loop is crossed an even number of times, it is outside the loop, otherwise it is inside
Each loop tile is getting a new letter after mapping the loop.
-: b as it block the ray
|: p as the ray passes
7J: l as it is a left turn
FL: r as it is a right turn
turning left and then right, or vice versa, means it blocks the ray
turning to the same side twice means the ray passes by the side of the wall
*/

func day10b(c chan string) {

	total := 0
	pipeMap := make([][]byte, 0, 0)
	var startX, startY int
	for line := range c {
		if len(line) == 0 {
			continue
		}
		pipeLine := []byte(line)
		pipeMap = append(pipeMap, pipeLine)
		if idx := bytes.IndexByte(pipeLine, 'S'); idx >= 0 {
			startX = len(pipeMap) - 1
			startY = idx
		}
	}

	nextPipeX := startX
	nextPipeY := startY + 1
	direction := 'E'
	for {
		switch direction {
		case 'E':
			switch pipeMap[nextPipeX][nextPipeY] {
			case '-':
				pipeMap[nextPipeX][nextPipeY] = 'b'
				nextPipeY += 1
			case '7':
				pipeMap[nextPipeX][nextPipeY] = 'l'
				nextPipeX += 1
				direction = 'S'
			case 'J':
				pipeMap[nextPipeX][nextPipeY] = 'l'
				nextPipeX -= 1
				direction = 'N'
			}
		case 'N':
			switch pipeMap[nextPipeX][nextPipeY] {
			case '|':
				pipeMap[nextPipeX][nextPipeY] = 'p'
				nextPipeX -= 1
			case 'F':
				pipeMap[nextPipeX][nextPipeY] = 'r'
				nextPipeY += 1
				direction = 'E'
			case '7':
				pipeMap[nextPipeX][nextPipeY] = 'l'
				nextPipeY -= 1
				direction = 'W'
			}
		case 'W':
			switch pipeMap[nextPipeX][nextPipeY] {
			case '-':
				pipeMap[nextPipeX][nextPipeY] = 'b'
				nextPipeY -= 1
			case 'F':
				pipeMap[nextPipeX][nextPipeY] = 'r'
				nextPipeX += 1
				direction = 'S'
			case 'L':
				pipeMap[nextPipeX][nextPipeY] = 'r'
				nextPipeX -= 1
				direction = 'N'
			}
		case 'S':
			switch pipeMap[nextPipeX][nextPipeY] {
			case '|':
				pipeMap[nextPipeX][nextPipeY] = 'p'
				nextPipeX += 1
			case 'L':
				pipeMap[nextPipeX][nextPipeY] = 'r'
				nextPipeY += 1
				direction = 'E'
			case 'J':
				pipeMap[nextPipeX][nextPipeY] = 'l'
				nextPipeY -= 1
				direction = 'W'
			}
		}
		if pipeMap[nextPipeX][nextPipeY] == 'S' {
			break
		}
	}
	log.Printf("read map")

	lastRow := len(pipeMap) - 1
	for x := range pipeMap {
		for y := range pipeMap[x] {

			// part of the loop
			if pipeMap[x][y] == 'S' || pipeMap[x][y] == 'b' || pipeMap[x][y] == 'p' || pipeMap[x][y] == 'l' || pipeMap[x][y] == 'r' {
				continue
			}
			crossedLoop := 0
			var lastTurn byte
			for x2 := x + 1; x2 <= lastRow; x2++ {
				if pipeMap[x2][y] == 'b' || pipeMap[x2][y] == 'S' {
					crossedLoop++
					continue
				}
				if pipeMap[x2][y] == 'l' || pipeMap[x2][y] == 'r' {
					// first turn
					if lastTurn == 0 {
						lastTurn = pipeMap[x2][y]
						continue
					}
					// second turn to other side is considered one loop wall
					if lastTurn != pipeMap[x2][y] {
						lastTurn = 0
						crossedLoop++
						continue
					}
					// turn to the same side is not loop wall
					lastTurn = 0
				}
			}
			// in the loop
			if crossedLoop%2 == 1 {
				log.Printf("In loop %d %d", x, y)
				total++
			}
		}
	}

	log.Printf("10B tiles inside loop: %d", total)
}
