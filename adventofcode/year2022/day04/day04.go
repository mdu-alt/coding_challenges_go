// Package day04 implements Day 4 of Advent of Code, 2022: [Camp Cleanup].
//
// [Camp Cleanup]: https://adventofcode.com/2022/day/4
package day04

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		pairs          []string
		l1, r1, l2, r2 int
		containOther   int
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pairs = strings.Split(scanner.Text(), ",")

		l1, r1 = getSections(strings.Split(pairs[0], "-"))
		l2, r2 = getSections(strings.Split(pairs[1], "-"))

		if (l1 <= l2 && r1 >= r2) || (l1 >= l2 && r1 <= r2) {
			containOther++
		}
	}

	return containOther
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		pairs          []string
		l1, r1, l2, r2 int
		overlapping    int
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pairs = strings.Split(scanner.Text(), ",")

		l1, r1 = getSections(strings.Split(pairs[0], "-"))
		l2, r2 = getSections(strings.Split(pairs[1], "-"))

		if !(r1 < l2 || l1 > r2) {
			overlapping++
		}
	}

	return overlapping
}

// -----------------------------------------------------------------------------

func getSections(ss []string) (int, int) {
	l, _ := strconv.Atoi(ss[0])
	r, _ := strconv.Atoi(ss[1])
	return l, r
}
