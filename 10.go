package main

import (
	"bytes"
	"log"
)

func day10(c chan string) {

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
	step := 1
	for {
		step++
		switch direction {
		case 'E':
			switch pipeMap[nextPipeX][nextPipeY] {
			case '-':
				nextPipeY += 1
			case '7':
				nextPipeX += 1
				direction = 'S'
			case 'J':
				nextPipeX -= 1
				direction = 'N'
			}
		case 'N':
			switch pipeMap[nextPipeX][nextPipeY] {
			case '|':
				nextPipeX -= 1
			case 'F':
				nextPipeY += 1
				direction = 'E'
			case '7':
				nextPipeY -= 1
				direction = 'W'
			}
		case 'W':
			switch pipeMap[nextPipeX][nextPipeY] {
			case '-':
				nextPipeY -= 1
			case 'F':
				nextPipeX += 1
				direction = 'S'
			case 'L':
				nextPipeX -= 1
				direction = 'N'
			}
		case 'S':
			switch pipeMap[nextPipeX][nextPipeY] {
			case '|':
				nextPipeX += 1
			case 'L':
				nextPipeY += 1
				direction = 'E'
			case 'J':
				nextPipeY -= 1
				direction = 'W'
			}
		}
		if pipeMap[nextPipeX][nextPipeY] == 'S' {
			break
		}
	}

	log.Printf("10A Halfway point: %d", step/2)
}
