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
		"1":  day1,
		"1b": day1b,
		"2":  day2,
		"2b": day2b,
		"3":  day3,
		"3b": day3b,
		"4":  day4,
		"4b": day4b,
		"5":  day5,
		"5b": day5b,
		"6":  day6,
		"6b": day6b,
		"7":  day7,
		"7b": day7b,
		"8":  day8,
		"8b": day8b,
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
