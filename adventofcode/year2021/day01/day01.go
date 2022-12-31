// Package day01 implements Day 1 of Advent of Code, 2021: [Sonar Sweep].
//
// [Sonar Sweep]: https://adventofcode.com/2021/day/1
package day01

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) int {
	var (
		measurements = extractMeasurements(filename)
		increases    int
	)

	for i, n := range measurements[1:] {
		if n > measurements[i] {
			increases++
		}
	}

	return increases
}

func part2(filename string) int {
	var (
		measurements = extractMeasurements(filename)
		increases    int
	)

	for i, n := range measurements[3:] {
		if n > measurements[i] {
			increases++
		}
	}

	return increases
}

// -----------------------------------------------------------------------------

func extractMeasurements(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	ss := strings.Fields(string(data))
	measurements := make([]int, len(ss))

	for i, s := range ss {
		measurements[i], _ = strconv.Atoi(s)
	}

	return measurements
}
