// Package day03 implements Day 3 of Advent of Code, 2022: [Rucksack Reorganization].
//
// [Rucksack Reorganization]: https://adventofcode.com/2022/day/3
package day03

import (
	"bufio"
	"log"
	"os"
)

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		rucksack      string
		first, second RuneSet
		sum           int
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack = scanner.Text()

		first = makeRuneSet(rucksack[:len(rucksack)/2])
		second = makeRuneSet(rucksack[len(rucksack)/2:])

		// `intersect` should return only one item.
		for itemType := range intersect(first, second) {
			sum += toPriority(itemType)
		}
	}

	return sum
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		rucksacksGroup = make([]RuneSet, 0)
		sum            int
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(rucksacksGroup) < 3 {
			rucksacksGroup = append(rucksacksGroup, makeRuneSet(scanner.Text()))
		} else {
			// `intersect` should return only one item.
			for itemType := range intersect(rucksacksGroup...) {
				sum += toPriority(itemType)
			}
			rucksacksGroup = []RuneSet{makeRuneSet(scanner.Text())}
		}
	}

	for itemType := range intersect(rucksacksGroup...) {
		sum += toPriority(itemType)
	}

	return sum
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

func intersect(runeSets ...RuneSet) RuneSet {
	intersection := make(RuneSet)

	for _, runeSet := range runeSets {
		if len(intersection) == 0 {
			intersection = runeSet
		} else {
			innerIntersection := make(RuneSet)
			for r := range runeSet {
				if _, ok := intersection[r]; ok {
					innerIntersection[r] = true
				}
			}
			intersection = innerIntersection
		}
	}

	return intersection
}

func toPriority(r rune) int {
	if r >= 97 {
		return int(r - 96) // [A-Z] -> [27-52]
	}
	return int(r - 38) // [a-z] -> [1-26]
}
