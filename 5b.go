package main

import (
	"log"
	"math"
	"sort"
	"strings"
)

func day5b(c chan string) {
	var next, nextRange, prev, prevRange []int64
	var nearestLocation int64

	for line := range c {
		if len(line) == 0 {
			continue
		}
		//log.Printf("------------------------------------------")
		//log.Printf(line)

		if strings.Contains(line, "seeds: ") {
			seedNumbers := readLineOfNumbers(line[7:], " ")
			for i := 0; i < len(seedNumbers)/2; i++ {
				next = append(next, seedNumbers[i*2])
				nextRange = append(nextRange, seedNumbers[i*2+1])
			}
			continue
		}
		if strings.Contains(line, "map:") {
			next = append(next, prev...)
			nextRange = append(nextRange, prevRange...)
			prev = next
			prevRange = nextRange

			next = nil
			nextRange = nil
			continue
		}
		nextLine := readLineOfNumbers(line, " ")
		destinationStart := nextLine[0]
		sourceStart := nextLine[1]
		mapRange := nextLine[2]

		/*for i := range prev {
			log.Printf("prev: %d range: %d", prev[i], prevRange[i])
		}*/
		for i := range prev {
			if prev[i] >= sourceStart && prev[i] < sourceStart+mapRange {
				//log.Printf("start within")
				if prev[i]+prevRange[i] < sourceStart+mapRange {
					//log.Printf("end within")
					// ends within the range
					next = append(next, destinationStart+prev[i]-sourceStart)
					nextRange = append(nextRange, prevRange[i])
				} else {
					//log.Printf("end after")
					// ends after range
					next = append(next, destinationStart+prev[i]-sourceStart)
					nextRange = append(nextRange, prev[i]-sourceStart+mapRange)
					next = append(next, sourceStart+mapRange)
					nextRange = append(nextRange, prev[i]+prevRange[i]-sourceStart-mapRange)
				}
				prev[i] = -1
				prevRange[i] = -1
			}
			if prev[i] < sourceStart && prev[i]+prevRange[i] >= sourceStart {
				//log.Printf("start before prev+range: %d sourceStart: %d", prev[i]+prevRange[i], sourceStart)
				if prev[i]+prevRange[i] < sourceStart+mapRange {
					//log.Printf("end within")
					// ends within the range
					next = append(next, destinationStart)
					nextRange = append(nextRange, prev[i]+prevRange[i]-sourceStart)
				} else {
					//log.Printf("end after")
					// ends after range
					next = append(next, destinationStart)
					nextRange = append(nextRange, mapRange)
					next = append(next, sourceStart+mapRange)
					nextRange = append(nextRange, prev[i]+prevRange[i]-sourceStart-mapRange)
				}
				// adjust range before start
				next = append(next, prev[i])
				nextRange = append(nextRange, sourceStart-prev[i])
				prev[i] = -1
				prevRange[i] = -1
			}
			/*for i := range next {
				log.Printf("next: %d range: %d", next[i], nextRange[i])
			}*/
		}
		prev = removeItem(prev, -1)
		prevRange = removeItem(prevRange, -1)
	}
	prev = removeItem(prev, 0)
	prevRange = removeItem(prevRange, 0)

	next = append(next, prev...)
	next = unique(next)
	sort.Slice(next, func(i, j int) bool {
		return next[i] < next[j]
	})
	nearestLocation = math.MaxInt64
	for _, v := range next {
		log.Printf("location: %d", v)
		if v < nearestLocation {
			nearestLocation = v
		}
	}
	log.Printf("5B Nearest location: %d", nearestLocation)
}

func removeItem2[T comparable](list []T, item T) ([]T, int) {
	var index int
	for i, element := range list {
		if element == item {
			index = i
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list, index
}

func removeItemByIndex[T comparable](list []T, index int) []T {
	return append(list[:index], list[index+1:]...)
}

/*
should be 219529182
*/
