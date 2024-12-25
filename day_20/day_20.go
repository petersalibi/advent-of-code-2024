package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

var directions = []utils.Pair{
    utils.UP,
    utils.DOWN,
    utils.LEFT,
    utils.RIGHT,
}

var circlePoints = generateCircularPoints(20)

func generateCircularPoints(radius int) []utils.Pair {
    points := make([]utils.Pair, 0)
    // for off by 1 error
    radius = radius + 1

    for x := range radius {
        for y := range radius {
            if x + y >= radius || (x <= 1 && y <= 1) {
                continue
            }
            if x == 0 {
                points = append(points, utils.NewPair(x, y), utils.NewPair(x, -y))
            } else if y == 0 {
                points = append(points, utils.NewPair(x, y), utils.NewPair(-x, y))
            } else {
                points = append(points, utils.NewPair(x, y), utils.NewPair(x, -y), utils.NewPair(-x, y), utils.NewPair(-x, -y))
            }
        }
    }

    return points
}

func findPath(raceMap []string, initPos utils.Pair, finalPos utils.Pair) ([][]int, []utils.Pair) {
    intMapping := make([][]int, len(raceMap))
    path := make([]utils.Pair, 0)
    curPos := initPos

    for i := range raceMap {
        intMapping[i] = make([]int, len(raceMap[0]))
    }

    raceTime := 0
    for !utils.Equal(curPos, finalPos) {
        intMapping[curPos.Y][curPos.X] = raceTime
        path = append(path, curPos)
        raceTime++
        for _, dir := range directions {
            newPos := utils.AddPair(curPos, dir)
            canMove := raceMap[newPos.Y][newPos.X] == '.' || raceMap[newPos.Y][newPos.X] == 'E'
            if canMove && intMapping[newPos.Y][newPos.X] == 0 {
                curPos = newPos
                break
            }
        }
    }
    intMapping[curPos.Y][curPos.X] = raceTime
    path = append(path, finalPos)

    return intMapping, path
}

type cheatPair struct {
    start utils.Pair
    end utils.Pair
}

func containsCheatPair(a cheatPair) func (cheatPair) bool {
    return func (b cheatPair) bool {
        return utils.Equal(a.start, b.start) && utils.Equal(a.end, b.end)
    }
}

func cheatRace(raceMap []string, intMap [][]int, threshold int, path []utils.Pair) int {
    overThreshold := 0
    // seen := make([]cheatPair, 0)

    for _, curPos := range path {
        var nextPos utils.Pair
        for _, dir := range circlePoints {
            cheatPos := utils.AddPair(curPos, dir)
            inBounds := utils.InArrayBounds(len(raceMap), len(raceMap[0]), cheatPos)
            if inBounds {
                if intMap[cheatPos.Y][cheatPos.X] - intMap[curPos.Y][curPos.X] - (utils.Abs(dir.X) + utils.Abs(dir.Y)) >= threshold {
                    // fmt.Println("Start, end", curPos, cheatPos)
                    // fmt.Println(intMap[cheatPos.Y][cheatPos.X], intMap[curPos.Y][curPos.X], intMap[cheatPos.Y][cheatPos.X] - intMap[curPos.Y][curPos.X] - 2)
                    // seen = append(seen, cheatPair{start: curPos, end: cheatPos})
                    overThreshold++
                }
            }
        }
        curPos = nextPos
    }

    return overThreshold
}

func findPos(raceMap []string, findLetter byte) utils.Pair {
    for y, line := range raceMap {
        for x, letter := range line {
            if byte(letter) == findLetter {
                return utils.NewPair(x, y)
            }
        }
    }
    return utils.NewPair(0, 0)
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n")
    startPos := findPos(splitData, 'S')
    endPos := findPos(splitData, 'E')

    intMapping, path := findPath(splitData, startPos, endPos)
    // for _, line := range intMapping {
    //     fmt.Println(line)
    // }

    circle := make([][]byte, 43)

    for i := range circle {
        circle[i] = make([]byte, 43)
        circle[i] = []byte(strings.Repeat(".", 43))
    }

    for _, point := range circlePoints {
        circle[21 + point.Y][21 + point.X] = 'O'
    }

    for _, line := range circle {
        fmt.Println(string(line))
    }

    // fmt.Println(circlePoints)

    fmt.Println("Part 2:", cheatRace(splitData, intMapping, 100, path))
}
