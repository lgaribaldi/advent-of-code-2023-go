package main

import (
	"log"
)

func day14b(c chan string) {
	var total int
	var platform [][]rune
	patternMap := make(map[string]int)

	for line := range c {
		if len(line) == 0 {
			continue
		}
		platform = append(platform, []rune(line))
	}

	for cycle := 0; cycle < 139; cycle++ {
		//N
		for row := range platform {
			for col := range platform[row] {
				if platform[row][col] == '.' {
				rollNorth:
					for i := row + 1; i < len(platform); i++ {
						switch platform[i][col] {
						case '.':
							continue rollNorth
						case '#':
							break rollNorth
						case 'O':
							platform[row][col] = 'O'
							platform[i][col] = '.'
							break rollNorth
						}
					}
				}
			}
		}
		//W
		for row := range platform {
			for col := range platform[row] {
				if platform[row][col] == '.' {
				rollWest:
					for j := col + 1; j < len(platform[0]); j++ {
						switch platform[row][j] {
						case '.':
							continue rollWest
						case '#':
							break rollWest
						case 'O':
							platform[row][col] = 'O'
							platform[row][j] = '.'
							break rollWest
						}
					}
				}
			}
		}
		//S
		for row := len(platform) - 1; row >= 0; row-- {
			for col := range platform[row] {
				if platform[row][col] == '.' {
				rollSouth:
					for i := row - 1; i >= 0; i-- {
						switch platform[i][col] {
						case '.':
							continue rollSouth
						case '#':
							break rollSouth
						case 'O':
							platform[row][col] = 'O'
							platform[i][col] = '.'
							break rollSouth
						}
					}
				}
			}
		}
		//E
		for row := range platform {
			for col := len(platform[row]) - 1; col >= 0; col-- {
				if platform[row][col] == '.' {
				rollEast:
					for j := col - 1; j >= 0; j-- {
						switch platform[row][j] {
						case '.':
							continue rollEast
						case '#':
							break rollEast
						case 'O':
							platform[row][col] = 'O'
							platform[row][j] = '.'
							break rollEast
						}
					}
				}
			}
		}
		pattern := ""
		for row := range platform {
			pattern += string(platform[row])
		}
		// To solve this I looked at the list to see the pattern
		// It repeated every 21 cycles (after cycle 121)
		// Then I used 1000000000 mod 21 to find which cycle would be exactly the same
		// Cycle 139 is the same as 1000000000
		if prev, ok := patternMap[pattern]; ok {
			println("Cycle ", prev, "Repeated on", cycle)
		} else {
			patternMap[pattern] = cycle
		}
	}

	for row := range platform {
		println(string(platform[row]))
		for col := range platform[row] {
			if platform[row][col] == 'O' {
				total += len(platform) - row
			}
		}
	}
	log.Printf("14B Total: %d", total)
}
