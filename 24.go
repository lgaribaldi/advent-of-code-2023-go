package main

import (
	"log"
	"math"
)

type hailstone struct {
	x, y, z, vx, vy, vz float64
}

func day24(c chan string) {
	hailstones := make([]hailstone, 0)
	var comparisons, total, outside, past, parallel int

	for line := range c {
		if len(line) == 0 {
			continue
		}
		pos, vel := splitString(line, " @ ")
		posInt := readLineOfNumbers(pos, ", ")
		velInt := readLineOfNumbers(vel, ", ")
		stone := hailstone{float64(posInt[0]), float64(posInt[1]), float64(posInt[2]), float64(velInt[0]), float64(velInt[1]), float64(velInt[2])}
		hailstones = append(hailstones, stone)
	}

	for i := range hailstones {
		for j := i + 1; j < len(hailstones); j++ {
			comparisons++
			var mA, mB, x, y float64
			A := hailstones[i]
			B := hailstones[j]
			mA = A.vy / A.vx
			mB = B.vy / B.vx
			x = ((mA * A.x) - A.y - (mB * B.x) + B.y) / (mA - mB)
			y = A.y + mA*(x-A.x)
			// y2 = B.y + mB*(x-B.x) dont really need
			if !math.IsInf(math.Abs(y), 0) {
				if (x >= 200000000000000 && x <= 400000000000000) && (y >= 200000000000000 && y <= 400000000000000) {
					if (x > A.x && A.vx < 0) || (y > A.y && A.vy < 0) || (y < A.y && A.vy > 0) || (x < A.x && A.vx > 0) ||
						(x > B.x && B.vx < 0) || (y > B.y && B.vy < 0) || (y < B.y && B.vy > 0) || (x < B.x && B.vx > 0) {
						past++
					} else {
						total++
					}
				} else {
					outside++
				}
			} else {
				parallel++
				// paralell lines should return NaN
				println("parallel", i, j)
			}
		}
	}

	log.Printf("24A total: %d", total)
	log.Printf("comparisons: %d, parallel: %d, outside: %d, past: %d,", comparisons, parallel, outside, past)
}
