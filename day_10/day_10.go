package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

var directions = []utils.Pair{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: -1},
}

func walkTrail(trail []string, curNum int, location utils.Pair) []utils.Pair {
    if curNum == 9 {
        return []utils.Pair{location}
    }

    trailHeads := []utils.Pair{}
    yLimit := len(trail)
    xLimit := len(trail[0])

    for _, direction := range directions {
        newLocation := utils.AddPair(location, direction)
        if !utils.InArrayBounds(xLimit, yLimit, newLocation) {
            continue
        }
        nextTopology, _ := strconv.Atoi(string(trail[newLocation.Y][newLocation.X]))

        if nextTopology == curNum + 1 {
            newPoints := walkTrail(trail, nextTopology, newLocation)
            for _, point := range newPoints {
                checkPoint := func (p utils.Pair) bool { return utils.Equal(p, point) }
                if !slices.ContainsFunc(trailHeads, checkPoint) {
                    trailHeads = append(trailHeads, point)
                }
            }
        }
    }

    return trailHeads
}

func walkOverAll(trail []string) int {
    allTrailHeads := 0
    for y, line := range trail {
        for x, num := range line {
            if num != '0' {
                continue
            }
            height, _ := strconv.Atoi(string(num))

            allTrailHeads += len(walkTrail(trail, height, utils.Pair{X: x, Y: y}))
        }
    }
    return allTrailHeads
}

func main() {
	data := utils.GetDataFromFile()
	splitData := strings.Split(data, "\n")

    heads := walkOverAll(splitData)

    fmt.Println("Part 1:", heads)
}
