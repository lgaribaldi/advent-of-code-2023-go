package main

import (
	"container/heap"
	"fmt"
	"log"
)

type Coordinate struct {
	x, y int
}

type PathNode struct {
	coord, prev, aPrev           Coordinate
	dist, aDist, predictedRemain int
}

const DEFAULT_DISTANCE = 999999

func day17(c chan string) {

	grid := make([][]int, 0)
	lowestSoFar := DEFAULT_DISTANCE

	for line := range c {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, readLineOfInt(line, ""))
		//grid = append(grid, []byte(line))
	}
	endX := len(grid) - 1
	endY := len(grid[0]) - 1
	nodeGrid := make([][]PathNode, 0)
	for i := range grid {
		row := make([]PathNode, 0)
		for j := range grid[0] {
			coord := Coordinate{i, j}
			row = append(row, newNode(coord, endX, endY))
		}
		println("Added line to nodegrid", row)
		nodeGrid = append(nodeGrid, row)
	}
	forkMap := make([][]bool, len(grid))
	for i := range forkMap {
		forkMap[i] = make([]bool, len(grid[0]))
	}

	startNode := nodeGrid[0][0]
	startNode.dist = 0
	startNode.predictedRemain -= DEFAULT_DISTANCE
	openSet := &PathNodeHeap{&startNode}
	heap.Init(openSet)
	findCruciblePath(openSet, &nodeGrid, &grid, endX, endY, &forkMap, &lowestSoFar)

	// log.Printf("17A : %d", nodeGrid[endX][endY].dist)

	// Prints result to make it easier to debug
	printGrid := make([][]rune, endX+1)
	for i := range printGrid {
		printGrid[i] = make([]rune, endY+1)
		for j := range printGrid[i] {
			printGrid[i][j] = []rune(fmt.Sprint(grid[i][j]))[0]
		}
	}
	nextCoord := nodeGrid[endX][endY].coord
	for {
		if nextCoord.x == 0 && nextCoord.y == 0 {
			break
		}
		printGrid[nextCoord.x][nextCoord.y] = '#'
		nextCoord = nodeGrid[nextCoord.x][nextCoord.y].prev
	}
	for i := range printGrid {
		println(string(printGrid[i]))
	}
}

func newNode(coord Coordinate, endX, endY int) PathNode {
	return PathNode{coord, Coordinate{0, 0}, Coordinate{0, 0}, DEFAULT_DISTANCE, DEFAULT_DISTANCE, DEFAULT_DISTANCE + calculatePredictedRemain(coord, endX, endY)}
}

func findCruciblePath(set *PathNodeHeap, nGrid *[][]PathNode, grid *[][]int, endX, endY int, forkMap *[][]bool, lowestSoFar *int) {
	nodeGrid := *nGrid
	for set.Len() > 0 {
		if currentNode, ok := heap.Pop(set).(*PathNode); ok {
			if currentNode.dist > 942 {
				break
			}
			// log.Println(currentNode)
			//neighbors := []Coordinate{{currentNode.coord.x - 1, currentNode.coord.y}, {currentNode.coord.x + 1, currentNode.coord.y}, {currentNode.coord.x, currentNode.coord.y - 1}, {currentNode.coord.x, currentNode.coord.y + 1}}
			neighbors := []Coordinate{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
			// prevPrevNode := nodeGrid[nodeGrid[currentNode.coord.x][currentNode.coord.y].prev.x][nodeGrid[currentNode.coord.x][currentNode.coord.y].prev.y].prev
			prevNode := nodeGrid[currentNode.coord.x][currentNode.coord.y].prev
			for _, neighbor := range neighbors {

				if intAbs(currentNode.coord.x+(neighbor.x)-prevNode.x) > 1 || intAbs(currentNode.coord.y+(neighbor.y)-prevNode.y) > 1 {
					continue
				}

				prev := currentNode
				for offset := 1; offset <= 3; offset++ {
					// out of bounds
					if currentNode.coord.x+(neighbor.x*offset) < 0 || currentNode.coord.y+(neighbor.y*offset) < 0 || currentNode.coord.x+(neighbor.x*offset) > endX || currentNode.coord.y+(neighbor.y*offset) > endY {
						continue
					}

					/*if intAbs(currentNode.coord.x + (neighbor.x * offset)-prevPrevNode.x) > 2 || intAbs(currentNode.coord.y + (neighbor.y * offset)-prevPrevNode.y) > 2 {
						continue
					}*/

					curr := &nodeGrid[currentNode.coord.x+(neighbor.x*offset)][currentNode.coord.y+(neighbor.y*offset)]
					// dist is dist of prev + weight of the current node
					distanceSoFar := prev.dist + (*grid)[currentNode.coord.x+(neighbor.x*offset)][currentNode.coord.y+(neighbor.y*offset)]
					if distanceSoFar < curr.dist {
						curr.aDist = curr.dist
						curr.aPrev = curr.prev

						curr.dist = distanceSoFar
						curr.prev = Coordinate{prev.coord.x, prev.coord.y}
						curr.predictedRemain = distanceSoFar + calculatePredictedRemain(neighbor, endX, endY)
						set.Push(curr)
					} else if distanceSoFar < curr.aDist {
						curr.aDist = distanceSoFar
						curr.aPrev = Coordinate{prev.coord.x, prev.coord.y}
					}
					// check if another path could have been better
					if offset == 1 {
						if curr.aDist+(*grid)[currentNode.coord.x][currentNode.coord.y] < distanceSoFar {
							bifurcationNode := &nodeGrid[curr.prev.x][curr.prev.y].prev
							if bifurcationNode.x+bifurcationNode.y > 0 && !(*forkMap)[currentNode.coord.x][currentNode.coord.y] {
								(*forkMap)[currentNode.coord.x][currentNode.coord.y] = true
								newNodeGrid := make([][]PathNode, len(nodeGrid))
								for i := range nodeGrid {
									row := make([]PathNode, len(nodeGrid[0]))
									copy(row, nodeGrid[i])
									newNodeGrid[i] = row
								}
								newNodeGrid[curr.prev.x][curr.prev.y] = newNode(curr.prev, endX, endY)
								newNodeGrid[curr.coord.x][curr.coord.y] = newNode(curr.coord, endX, endY)

								startNode := newNodeGrid[bifurcationNode.x][bifurcationNode.y]
								startNode.dist = startNode.aDist
								startNode.prev = startNode.aPrev
								openSet := &PathNodeHeap{&startNode}
								heap.Init(openSet)
								findCruciblePath(openSet, &newNodeGrid, grid, endX, endY, forkMap, lowestSoFar)
							}

						}
					}
					prev = curr
				}

			}
		}
	}

	if nodeGrid[endX][endY].dist < *lowestSoFar {
		*lowestSoFar = nodeGrid[endX][endY].dist
		log.Printf("17A : %d", nodeGrid[endX][endY].dist)
	}

}

func calculatePredictedRemain(coord Coordinate, endX, endY int) int {
	result := (endX - coord.x + endY - coord.y) * 9
	return result
}

// A PathNodeHeap is a min-heap of PathNodes
type PathNodeHeap []*PathNode

func (h *PathNodeHeap) Len() int           { return len(*h) }
func (h *PathNodeHeap) Less(i, j int) bool { return (*h)[i].predictedRemain < (*h)[j].predictedRemain }
func (h *PathNodeHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *PathNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*PathNode))
}

func (h *PathNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// < 952
// < 942
// < 940
