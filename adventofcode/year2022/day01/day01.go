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

func part1(filename string) uint {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var calories, maxCalories uint

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			v, _ := strconv.ParseUint(scanner.Text(), 10, 32)
			calories += uint(v)
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

func part2(filename string) uint {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		calories    uint
		topCalories = &maxHeap{}
	)

	heap.Init(topCalories)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			v, _ := strconv.ParseUint(scanner.Text(), 10, 64)
			calories += uint(v)
		} else {
			heap.Push(topCalories, calories)
			calories = 0
		}
	}

	heap.Push(topCalories, calories)

	var top3 uint
	for i := 0; i < 3; i++ {
		top3 += heap.Pop(topCalories).(uint)
	}

	return top3
}

type maxHeap []uint

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(uint))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
