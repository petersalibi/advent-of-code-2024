package main

import (
	"fmt"
	"slices"
	"strings"

	"advent-of-code/utils"
)

type pair struct {
	x int
	y int
}

type guard struct {
	position  pair
	direction pair
}

func rotateDirection(vec pair) pair {
	// rotation by 90 degrees cw in 2d
	return pair{-vec.y, vec.x}
}

func inArrayBounds(arrLen, idx int) bool {
	return idx >= 0 && idx < arrLen
}

func equalPair(a, b pair) bool {
	return a.x == b.x && a.y == b.y
}

func hasPair(a pair) func(pair) bool {
	return func(b pair) bool {
		return equalPair(a, b)
	}
}

// Walks until it hits a '#'
func walkUntilObstacle(guardMap []string, curGuard guard, replaceTiles bool) (int, bool, guard) {
	numLocations := 0
	nextX := curGuard.position.x + curGuard.direction.x
	nextY := curGuard.position.y + curGuard.direction.y

	inBounds := inArrayBounds(len(guardMap), nextY) && inArrayBounds(len(guardMap[nextY]), nextX)

	for inBounds {
		switch guardMap[nextY][nextX] {
		case '.':
			if replaceTiles {
				guardMap[nextY] = replaceStringAtIndex(guardMap[nextY], 'X', nextX)
				numLocations++
			}
		case '#':
			curGuard.direction = rotateDirection(curGuard.direction)
	 		return numLocations, inBounds, curGuard
		}

		curGuard.position.x = nextX
		curGuard.position.y = nextY

		nextX = nextX + curGuard.direction.x
		nextY = nextY + curGuard.direction.y
		inBounds = inArrayBounds(len(guardMap), nextY) && inArrayBounds(len(guardMap[nextY]), nextX)
	}
	return numLocations, inBounds, curGuard
}

func checkIfLooping(guardMap []string, patrolGuard guard) bool {
	seenPos := make([]pair, 1)
	var inBounds bool
	var newGuard guard
	_, inBounds, newGuard = walkUntilObstacle(guardMap, patrolGuard, false)

	seenPos[0] = newGuard.position

	for inBounds {
		_, inBounds, newGuard = walkUntilObstacle(guardMap, newGuard, false)
        if slices.ContainsFunc(seenPos[:len(seenPos) - 1], hasPair(newGuard.position)) {
            // We don't check the last position in seenPos to avoid cases where we end up rotating in place
			return true
		}
		seenPos = append(seenPos, newGuard.position)
	}

	return false
}

func bruteForceObstacles(guardMap []string, startGuard guard) int {
	result := 0
	for y, line := range guardMap {
		for x, letter := range line {
			if letter != '.' {
				continue
			}
			guardMap[y] = replaceStringAtIndex(guardMap[y], '#', x)
			if checkIfLooping(guardMap, startGuard) {
				result++
			}
			guardMap[y] = replaceStringAtIndex(guardMap[y], '.', x)
		}
	}
	return result
}

func walkMap(guardMap []string, patrolGuard guard) int {
	// intialized at 1 to account for starting square
	result := 1
	var numLocations int
	var inBounds bool
	for {
		numLocations, inBounds, patrolGuard = walkUntilObstacle(guardMap, patrolGuard, true)
		result += numLocations
		if !inBounds {
			break
		}
	}
	return result
}

func replaceStringAtIndex(str []string, letter rune, index Pair) string {
	newStr := []rune(str)
	newStr[index] = letter
	return string(newStr)
}

func findGuard(input []string) pair {
	for y, line := range input {
		for x, letter := range line {
			if letter == '^' {
				return pair{x, y}
			}
		}
	}
	return pair{}
}

func main() {
	data := utils.GetDataFromFile()
	splitData := strings.Split(data, "\n")
	guardPos := findGuard(splitData)
	guardDirection := pair{0, -1}
	newGuard := guard{guardPos, guardDirection}
	part2 := bruteForceObstacles(splitData, newGuard)

	splitData[guardPos.y] = replaceStringAtIndex(splitData[guardPos.y], 'X', guardPos.x)

	fmt.Printf("Part 1: %d\n", walkMap(splitData, newGuard))
	fmt.Printf("Part 2: %d\n", part2)

	// for _, line := range splitData {
	//     fmt.Println(line)
	// }
}
