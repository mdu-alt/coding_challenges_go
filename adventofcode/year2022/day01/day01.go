// Package day01 implements Day 1 of Advent of Code, 2022: [Calorie Counting].
//
// [Calorie Counting]: https://adventofcode.com/2022/day/1
package day01

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
)

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var calories, maxCalories int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			v, _ := strconv.Atoi(scanner.Text())
			calories += v
		} else {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
		}
	}

	if calories > maxCalories {
		maxCalories = calories
	}

	return maxCalories
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		calories, top3 int
		topCalories    = new(MaxHeap)
	)

	heap.Init(topCalories)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			v, _ := strconv.Atoi(scanner.Text())
			calories += v
		} else {
			heap.Push(topCalories, calories)
			calories = 0
		}
	}

	heap.Push(topCalories, calories)

	for i := 0; i < 3; i++ {
		top3 += heap.Pop(topCalories).(int)
	}

	return top3
}

// -----------------------------------------------------------------------------

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
