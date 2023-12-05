package main

import (
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
