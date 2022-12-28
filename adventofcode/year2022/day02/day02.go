// Package day02 implements Day 2 of Advent of Code, 2022: [Rock Paper Scissors].
//
// [Rock Paper Scissors]: https://adventofcode.com/2022/day/2
package day02

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		score  int
		scores = map[string]int{
			"A": Rock, "X": Rock,
			"B": Paper, "Y": Paper,
			"C": Scissors, "Z": Scissors,
		}
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moves := strings.Fields(scanner.Text())
		opponent, player := scores[moves[0]], scores[moves[1]]

		switch player {
		case opponent%3 + 1:
			score += 6
		case opponent:
			score += 3
		}

		score += player
	}

	return score
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		score  int
		scores = map[string]int{
			"A": Rock, "X": Lose,
			"B": Paper, "Y": Draw,
			"C": Scissors, "Z": Win,
		}
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moves := strings.Fields(scanner.Text())
		opponent, outcome := scores[moves[0]], scores[moves[1]]

		switch outcome {
		case Lose:
			score += (opponent+1)%3 + 1
		case Draw:
			score += 3 + opponent
		case Win:
			score += 6 + opponent%3 + 1
		}
	}

	return score
}

// -----------------------------------------------------------------------------

const (
	Rock = iota + 1
	Paper
	Scissors
)

const (
	Lose = iota + 1
	Draw
	Win
)
