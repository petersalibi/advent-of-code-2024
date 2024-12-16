package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

const COLS = 101
const ROWS = 103

// const COLS = 11
// const ROWS = 7

var digits = regexp.MustCompile(`(-?\d+),(-?\d+)`)

func moveRobot(robot string, seconds int) (int, int, int) {
    coords := digits.FindAllStringSubmatch(robot, -1)
    rx, _ := strconv.Atoi(coords[0][1])
    ry, _ := strconv.Atoi(coords[0][2])

    vx, _ := strconv.Atoi(coords[1][1])
    if vx < 0 {
        vx = vx + COLS
    }
    vy, _ := strconv.Atoi(coords[1][2])
    if vy < 0 {
        vy = vy + ROWS
    }
    rx = (rx + seconds * vx) % COLS
    ry = (ry + seconds * vy) % ROWS
    // fmt.Println(rx, ry, vx, vy)
    if rx > COLS / 2 && ry < ROWS / 2 {
        return 0, rx, ry
    }

    if rx < COLS / 2 && ry < ROWS / 2 {
        return 1, rx, ry
    }

    if rx < COLS / 2 && ry > ROWS / 2 {
        return 2, rx, ry
    }

    if rx > COLS / 2 && ry > ROWS / 2 {
        return 3, rx, ry
    }

    return 4, rx, ry
}

func makeBlankString() []string {
    outputString := make([]string, 0)

    for range ROWS {
        outputString = append(outputString, strings.Repeat(".", COLS))
    }

    return outputString
}

func outputRobots(robots []string) {
    outputString := makeBlankString()
    minValue := 2147000000
    minIndex := 0
    robotPositions := make([]utils.Pair, len(robots))
    minPositions := make([]utils.Pair, len(robots))

    for i := range ROWS*COLS {
        quadrants := make([]int, 5)

        for j, robot := range robots {
            quadIndex, rx, ry := moveRobot(robot, i)
            quadrants[quadIndex]++
            robotPositions[j] = utils.NewPair(rx, ry)
        }
        safety := 1

        for _, numRob := range quadrants[:4] {
            safety *= numRob
        }
        if safety < minValue {
            minValue = safety
            minIndex = i
            // for k := range robotPositions {
            //     minPositions[k] = robotPositions[k]
            // }
            copy(minPositions, robotPositions)
        }
    }

    for _, robot := range minPositions {
        outputString[robot.Y] = utils.ReplaceStringAtIndex(outputString, 'X', robot)
    }
    fileName := "output/out.txt"
    fileString := strings.Join(outputString, "\n")
    newFile, err := os.Create(fileName)

    if err != nil {
        fmt.Fprint(os.Stderr, err)
        return
    }

    newFile.WriteString(fileString)
    newFile.Close()
    fmt.Println("Part 2:", minIndex)
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n")
    quadrants := make([]int, 5)

    // fmt.Println(splitData[0])
    // fmt.Println(moveRobot(splitData[0], 100))

    for _, robot := range splitData {
        index, _, _ := moveRobot(robot, 100)
        quadrants[index]++
    }

    safety := 1

    for _, numRob := range quadrants[:4] {
        safety *= numRob
    }

    outputRobots(splitData)

    fmt.Println("Part 1:", safety)
}
