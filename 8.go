package main

import (
	"log"
)

type Node struct {
	left  string
	right string
}

func day8(c chan string) {
	var total int64
	path, nodes := initializeDay8(c)
	nextNode := "AAA"

repeatPath:
	for {
		for i := range path {
			log.Printf("Next node: %s", nextNode)
			total++
			if path[i] == 'L' {
				nextNode = nodes[nextNode].left
			} else {
				nextNode = nodes[nextNode].right
			}
			if nextNode == "ZZZ" {
				break repeatPath
			}
		}
	}

	log.Printf("8A Total: %d", total)

}

func initializeDay8(c chan string) (string, map[string]Node) {
	nodes := make(map[string]Node)
	path := <-c
	for line := range c {
		if len(line) == 0 {
			continue
		}
		nodes[line[:3]] = Node{line[7:10], line[12:15]}
	}
	return path, nodes
}
