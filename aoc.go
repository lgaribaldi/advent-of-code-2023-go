package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	args := os.Args
	c := make(chan string)
	go readInputLineByLine(args[1]+".txt", c)

	functions := map[string]func(c chan string){
		"1":   day1,
		"1b":  day1b,
		"2":   day2,
		"2b":  day2b,
		"3":   day3,
		"3b":  day3b,
		"4":   day4,
		"4b":  day4b,
		"5":   day5,
		"5b":  day5b,
		"6":   day6,
		"6b":  day6b,
		"7":   day7,
		"7b":  day7b,
		"8":   day8,
		"8b":  day8b,
		"9":   day9,
		"9b":  day9b,
		"10":  day10,
		"10b": day10b,
		"11":  day11,
		"11b": day11b,
		"12":  day12,
		"12b": day12b,
		"13":  day13,
		"13b": day13b,
		"14":  day14,
		"14b": day14b,
		"15":  day15,
		"15b": day15b,
		"16":  day16,
		"16b": day16b,
		"17":  day17,
		"18":  day18,
		"18b": day18b,
		"19":  day19,
		"19b": day19b,
		"20":  day20,
		"20b": day20b,
		"21":  day21,
		"21b": day21b,
	}

	if fn, ok := functions[args[1]]; ok {
		fn(c)
	} else {
		log.Fatal("Function not found.")
	}

}

func readInputLineByLine(fileName string, c chan string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	close(c)
}
