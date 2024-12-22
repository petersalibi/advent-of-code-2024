package main

import (
	"fmt"
	"slices"
	"strings"

	"advent-of-code/utils"
)

var (
	UP    = utils.NewPair(0, -1)
	DOWN  = utils.NewPair(0, 1)
	RIGHT = utils.NewPair(1, 0)
	LEFT  = utils.NewPair(-1, 0)
)

func moveRobotPart1(warehouse [][]byte, direction utils.Pair, pos utils.Pair) utils.Pair {
	newPos := utils.Pair{X: pos.X, Y: pos.Y}
	spacesMoved := 0
	for utils.IndexArray(warehouse, newPos) != '#' && utils.IndexArray(warehouse, newPos) != '.' {
		newPos.X += direction.X
		newPos.Y += direction.Y
		spacesMoved++
	}

	if utils.IndexArray(warehouse, newPos) == '#' {
		return pos
	}

	if spacesMoved == 1 {
		warehouse[pos.Y][pos.X] = '.'
		warehouse[newPos.Y][newPos.X] = '@'
		return newPos
	}

	warehouse[newPos.Y][newPos.X] = 'O'
	warehouse[pos.Y][pos.X] = '.'

	// set newPos to move robot 1 move ahead
	newPos.X = pos.X + direction.X
	newPos.Y = pos.Y + direction.Y
	warehouse[newPos.Y][newPos.X] = '@'
	return newPos
}

func moveRobotPart2(warehouse [][]byte, direction utils.Pair, pos utils.Pair) utils.Pair {
	newPos := utils.Pair{X: pos.X, Y: pos.Y}
	spacesMoved := 0
	for utils.IndexArray(warehouse, newPos) != '#' && utils.IndexArray(warehouse, newPos) != '.' {
		newPos.X += direction.X
		newPos.Y += direction.Y
		spacesMoved++
	}

	if utils.IndexArray(warehouse, newPos) == '#' {
		return pos
	}

	if spacesMoved == 1 {
		warehouse[pos.Y][pos.X] = '.'
		warehouse[newPos.Y][newPos.X] = '@'
		return newPos
	}

	if direction != DOWN && direction != UP {
		for range spacesMoved - 1 {
			warehouse[newPos.Y][newPos.X] = utils.IndexArray(warehouse, utils.SubPair(newPos, direction))
			newPos.X -= direction.X
			newPos.Y -= direction.Y
		}
		warehouse[pos.Y][pos.X] = '.'
		warehouse[newPos.Y][newPos.X] = '@'
		return newPos
	}

	aheadRobot := utils.AddPair(pos, direction)
    boxes := moveBigBox(warehouse, direction, aheadRobot)
    if len(boxes) == 0 {
        return pos
    }
    oldBoard := make([][]byte, len(warehouse))

    for i, line := range warehouse {
        oldBoard[i] = make([]byte, len(line))
        copy(oldBoard[i], line)
    }
    for _, box := range boxes {
        warehouse[box.Y][box.X] = '.'
    }
    for _, box := range boxes {
        warehouse[box.Y + direction.Y][box.X + direction.X] = oldBoard[box.Y][box.X]
    }
    warehouse[pos.Y][pos.X] = '.'
    warehouse[aheadRobot.Y][aheadRobot.X] = '@'
    return aheadRobot
}

func moveBigBox(warehouse [][]byte, direction utils.Pair, pos utils.Pair) []utils.Pair {
    frontier := make([]utils.Pair, 500)
    frontier[0] = pos

    frontierEnd := 1

    foundBoxes := make([]utils.Pair, 0)


    for frontierEnd != 0 {
        checkPos := frontier[frontierEnd - 1]
        if slices.ContainsFunc(foundBoxes, utils.ContainsPair(checkPos)) {
            frontierEnd--
            continue
        }
        warehouseChar := utils.IndexArray(warehouse, checkPos)
        switch warehouseChar {
        case '[':
            foundBoxes = append(foundBoxes, checkPos)
            frontier[frontierEnd] = utils.AddPair(checkPos, direction)
            frontierEnd++
            frontier[frontierEnd] = utils.AddPair(checkPos, RIGHT)
            frontierEnd++
        case ']':
            foundBoxes = append(foundBoxes, checkPos)
            frontier[frontierEnd] = utils.AddPair(checkPos, direction)
            frontierEnd++
            frontier[frontierEnd] = utils.AddPair(checkPos, LEFT)
            frontierEnd++
        case '#':
            return []utils.Pair{}
        default:
            frontierEnd--
        }
    }

    return foundBoxes
}

func decodeMove(warehouse [][]byte, dir byte, pos utils.Pair, part1 bool) utils.Pair {
	var direction utils.Pair
	switch dir {
	case '>':
		direction = RIGHT
	case '<':
		direction = LEFT
	case '^':
		direction = UP
	case 'v':
		direction = DOWN
	default:
		return pos
	}
	if part1 {
		return moveRobotPart1(warehouse, direction, pos)
	}
	return moveRobotPart2(warehouse, direction, pos)
}

func doRobotMoves(warehouse [][]byte, directions string, pos utils.Pair, part1 bool) {
	for _, dir := range directions {
		pos = decodeMove(warehouse, byte(dir), pos, part1)
	}
}

func findInitPos(warehouse [][]byte) utils.Pair {
	for y, line := range warehouse {
		for x, char := range line {
			if char == '@' {
				return utils.NewPair(x, y)
			}
		}
	}
	return utils.NewPair(0, 0)
}

func countGPS(warehouse [][]byte) int {
	sum := 0
	for y, line := range warehouse {
		for x, char := range line {
			if char == 'O' || char == '[' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func main() {
	data := utils.GetDataFromFile()
	dataSplit := strings.Split(data, "\n\n")
	warehouseString := strings.Split(dataSplit[0], "\n")
	warehouse := make([][]byte, len(warehouseString))
	warehousePart2 := make([][]byte, len(warehouseString))


	for i, line := range warehouseString {
		warehouse[i] = []byte(line)
	}

	lineLen := len(warehouse[0]) * 2
	for y, line := range warehouse {
		warehousePart2[y] = make([]byte, lineLen)
		for x, char := range line {
			switch char {
			case '#':
				warehousePart2[y][x*2] = '#'
				warehousePart2[y][x*2+1] = '#'
			case '.':
				warehousePart2[y][x*2] = '.'
				warehousePart2[y][x*2+1] = '.'
			case 'O':
				warehousePart2[y][x*2] = '['
				warehousePart2[y][x*2+1] = ']'
			case '@':
				warehousePart2[y][x*2] = '@'
				warehousePart2[y][x*2+1] = '.'
			}
		}
	}

    moves := dataSplit[1]//[:210]
    // moves := "^<>vv<v<^^<^<<^^^^>>>><<vv"

    initPos := findInitPos(warehouse)
	// // fmt.Println(initPos)
	doRobotMoves(warehouse, moves, initPos, true)
	// //
	for i, line := range warehouse {
		warehouseString[i] = string(line)
	}
	// //
	// // fmt.Println(strings.Join(warehouseString, "\n"))
	fmt.Println("Part 1:", countGPS(warehouse))

	warehouseString2 := strings.Split(dataSplit[0], "\n")
	for i, line := range warehousePart2 {
		warehouseString2[i] = string(line)
	}

	// fmt.Println(strings.Join(warehouseString2, "\n"))
    // fmt.Println(moves, string(dataSplit[1][316]))
    initPos = findInitPos(warehousePart2)
	// fmt.Println(initPos)
	doRobotMoves(warehousePart2, moves, initPos, false)

	for i, line := range warehousePart2 {
		warehouseString2[i] = string(line)
	}

	fmt.Println(strings.Join(warehouseString2, "\n"))
	fmt.Println("Part 2:", countGPS(warehousePart2))
}
