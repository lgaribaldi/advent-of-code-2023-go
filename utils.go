package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func isNumeric(b byte) bool {
	return unicode.IsDigit(rune(b))
}

func splitString(s, sep string) (string, string) {
	x := strings.Split(s, sep)
	return x[0], x[1]
}

func unique[T comparable](s []T) []T {
	var unique []T
	m := map[T]bool{}

	for _, v := range s {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

func nonZeroValues(s []int64) []int64 {
	var nonZero []int64
	for _, v := range s {
		if v != 0 {
			nonZero = append(nonZero, v)
		}
	}
	return nonZero
}

func nonEmptyValues(s []string) []string {
	var nonEmpty []string
	for _, v := range s {
		if v != "" {
			nonEmpty = append(nonEmpty, v)
		}
	}
	return nonEmpty
}

func intersection[T comparable](s1, s2 []T) []T {
	var intersect []T

block1:
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v2 == v1 {
				intersect = append(intersect, v1)
				continue block1
			}
		}
	}
	return intersect
}

func readLineOfNumbers(line string, sep string) []int64 {
	items := strings.Split(line, sep)
	var result []int64
	for _, item := range items {
		num, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, num)
	}
	return result
}

func readLineOfInt(line string, sep string) []int {
	items := strings.Split(line, sep)
	var result []int
	for _, item := range items {
		num, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, int(num))
	}
	return result
}

func removeItem[T comparable](list []T, item T) []T {
	result := make([]T, 0)
	for _, element := range list {
		if element != item {
			result = append(result, element)
		}
	}
	return result
}

func intAbs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}
