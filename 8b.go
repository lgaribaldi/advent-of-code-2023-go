package main

import (
	"log"
)

func day8b(c chan string) {
	var steps int
	path, nodes, ghosts := initializeDay8b(c)
	var ghostSteps []int

	log.Printf("Num of ghosts: %d", len(ghosts))
	for j := range ghosts {
		steps = 0
	repeatPath:
		for {
			for i := range path {
				steps++
				if path[i] == 'L' {
					ghosts[j] = nodes[ghosts[j]].left

				} else {
					ghosts[j] = nodes[ghosts[j]].right
				}
				if ghosts[j][2:3] == "Z" {
					log.Printf("Ghost %d Steps: %d", j+1, steps)
					ghostSteps = append(ghostSteps, steps)
					break repeatPath
				}
			}
		}
	}

	total := LCM(ghostSteps[0], ghostSteps[1], ghostSteps[2:]...)
	log.Printf("8B total: %d", total)
}

func initializeDay8b(c chan string) (string, map[string]Node, []string) {
	startingNodes := make([]string, 0, 0)
	nodes := make(map[string]Node)
	path := <-c
	for line := range c {
		if len(line) == 0 {
			continue
		}
		if line[2:3] == "A" {
			startingNodes = append(startingNodes, line[:3])
		}
		nodes[line[:3]] = Node{line[7:10], line[12:15]}
	}
	return path, nodes, startingNodes
}

// Got these from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
