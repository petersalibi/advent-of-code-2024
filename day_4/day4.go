package main

import (
	"fmt"
	"strings"
	"sync"

	"advent-of-code/utils"
)

const searchWord string = "XMAS"

type pair struct {
	x int
	y int
}

// top left is (0, 0)
var directions []pair = []pair{
	{1, 0},   // right
	{1, -1},  // up right
	{0, -1},  // up
	{-1, -1}, // up left
	{-1, 0},  // left
	{-1, 1},  // down left
	{0, 1},   // down
	{1, 1},   // down right
}

var diagonals []pair = []pair{
	{1, -1},  // up right
	{-1, -1}, // up left
	{-1, 1},  // down left
	{1, 1},   // down right
}

func readDirection(data []string, x, y int, direction pair) bool {
	for i := range searchWord {
		newY := y + i*direction.y
		newX := x + i*direction.x

		// Bounds checking
		if newY >= len(data) || newY < 0 {
			return false
		}
		if newX >= len(data[newY]) || newX < 0 {
			return false
		}

		if data[newY][newX] != searchWord[i] {
			return false
		}
	}

	return true
}

func readXMas(data []string, x, y int) bool {
	numM := 0
	numS := 0
	diagLetters := make([]byte, 4)
	for i, direction := range diagonals {
		newY := y + direction.y
		newX := x + direction.x
		// Bounds checking
		if newY >= len(data) || newY < 0 {
			return false
		}
		if newX >= len(data[newY]) || newX < 0 {
			return false
		}

		diagLetters[i] = data[newY][newX]
		switch data[newY][newX] {
		case 'M':
			numM++
		case 'S':
			numS++
		default:
			return false
		}
	}

	if numM == 2 && numS == 2 {
		if diagLetters[0] != diagLetters[2] && diagLetters[1] != diagLetters[3] {
			// Check if letters on the same diagonal are different
			return true
		}
	}
	return false
}

func findMatches(input []string, start, end int, part1 bool, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for y, line := range input[start:end] {
		for x, letter := range line {
			if letter == 'X' && part1 {
				for _, direction := range directions {
					if readDirection(input, x, y+start, direction) {
						result += 1
					}
				}
			}
			if letter == 'A' && !part1 {
				if readXMas(input, x, y+start) {
					result += 1
				}
			}
		}
	}
	ch <- result
}

func dispatchFindMatches(input []string, part1 bool, numThreads int) int {
	ch := make(chan int)
	var wg sync.WaitGroup

	chunkSize := len(input) / numThreads

	for i := 0; i < numThreads; i++ {
		start := i * chunkSize
		var end int

		if i+1 == numThreads {
			end = len(input)
		} else {
			end = (i + 1) * chunkSize
		}

		wg.Add(1)
		go findMatches(input, start, end, part1, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := 0
	for val := range ch {
		result += val
	}

	return result
}

func main() {
	data := utils.GetDataFromFile()
	dataLines := strings.Split(data, "\n")
	numThreads := 8

	fmt.Printf("Part 1: %d, Part 2: %d\n", dispatchFindMatches(dataLines, true, numThreads), dispatchFindMatches(dataLines, false, numThreads))
}
