// Package day06 implements Day 6 of Advent of Code, 2022: [Tuning Trouble].
//
// [Tuning Trouble]: https://adventofcode.com/2022/day/6
package day06

import (
	"log"
	"os"
)

func part1(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var marker int

	datastream := string(data)
	for i := range datastream {
		if len(makeRuneSet(datastream[i:i+4])) == 4 {
			marker = i + 4
			break
		}
	}

	return marker
}

func part2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	var marker int

	datastream := string(data)
	for i := range datastream {
		if len(makeRuneSet(datastream[i:i+14])) == 14 {
			marker = i + 14
			break
		}
	}

	return marker
}

// -----------------------------------------------------------------------------

type RuneSet map[rune]bool

func makeRuneSet(s string) RuneSet {
	rs := make(RuneSet)
	for _, r := range s {
		rs[r] = true
	}
	return rs
}
