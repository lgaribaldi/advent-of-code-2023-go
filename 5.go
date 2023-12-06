package main

import (
	"log"
	"math"
	"strings"
)

func day5(c chan string) {
	var next, prev []int64
	var nearestLocation int64

	for line := range c {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "seeds: ") {
			next = readLineOfNumbers(line[7:], " ")
			continue
		}
		if strings.Contains(line, "map:") {
			log.Printf(line)
			next = append(next, prev...)
			prev = next
			next = nil
			continue
		}
		nextLine := readLineOfNumbers(line, " ")
		destinationStart := nextLine[0]
		sourceStart := nextLine[1]
		mapRange := nextLine[2]
		for _, prevNumber := range prev {
			if prevNumber >= sourceStart && prevNumber < sourceStart+mapRange {
				next = append(next, destinationStart+prevNumber-sourceStart)
				prev = removeItem(prev, prevNumber)
			}
		}
	}

	next = append(next, prev...)
	nearestLocation = math.MaxInt64
	for _, v := range next {
		log.Printf("location: %d", v)
		if v < nearestLocation {
			nearestLocation = v
		}
	}
	log.Printf("5A Nearest location: %d", nearestLocation)
}
