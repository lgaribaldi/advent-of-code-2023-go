package main

import (
	"log"
)

func day14(c chan string) {
	var total int
	var platform []string

	for line := range c {
		if len(line) == 0 {
			continue
		}
		platform = append(platform, line)
	}

	for row := range platform {
		for col := range platform[row] {

			if platform[row][col] == '.' {
			rock:
				for i := row + 1; i < len(platform); i++ {
					switch platform[i][col] {
					case '.':
						continue rock
					case '#':
						break rock
					case 'O':
						northRow := []rune(platform[row])
						northRow[col] = 'O'
						platform[row] = string(northRow)

						southRow := []rune(platform[i])
						southRow[col] = '.'
						platform[i] = string(southRow)
						total += len(platform) - row
						break rock
					}
				}
			} else if platform[row][col] == 'O' {
				total += len(platform) - row
			}
		}
	}

	for row := range platform {
		println(platform[row])
	}
	log.Printf("14A Total: %d", total)
}
