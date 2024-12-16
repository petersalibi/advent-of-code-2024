package main

import (
	"fmt"
	"slices"
	"strings"

	"advent-of-code/utils"
)

var directions = []utils.Pair{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: -1},
}

func findRegion(regionMap []string, code byte, location utils.Pair, seenLocations *[]utils.Pair) int {
    if slices.ContainsFunc(*seenLocations, utils.ContainsPair(location)) {
        return 0
    }

    *seenLocations = append(*seenLocations, location)
    yBound := len(regionMap)
    xBound := len(regionMap[0])
    perimiter := 0
    for _, direction := range directions {
        newLoc := utils.AddPair(location, direction)
        if !utils.InArrayBounds(xBound, yBound, newLoc) || utils.IndexString(regionMap, newLoc) != code {
            perimiter++
            continue
        }
        perimiter += findRegion(regionMap, code, newLoc, seenLocations)
    }

    return perimiter
}

type region struct {
    perimiter int
    seenLocations *[]utils.Pair
}

func findAllRegions(regionMap []string) []region {
    regions := make([]region, 0)
    yLen := len(regionMap)
    xLen := len(regionMap[0])

    for y := range yLen {
        for x := range xLen {
            letter := regionMap[y][x]
            if letter == '.' {
                continue
            }
            newRegion := region{}
            newRegion.seenLocations = new([]utils.Pair)
            newRegion.perimiter = findRegion(regionMap, byte(letter), utils.NewPair(x, y), newRegion.seenLocations)
            for _, location := range *newRegion.seenLocations {
                regionMap[location.Y] = utils.ReplaceStringAtIndex(regionMap, '.', location)
            }
            regions = append(regions, newRegion)
        }
    }
    return regions
}

func regionEdges(reg region, strLen int) int {
    regionStr := make([]string, 0)
    for range strLen {
        newStr := strings.Repeat(".", strLen)
        regionStr = append(regionStr, newStr)
    }

    for _, point := range *reg.seenLocations {
        regionStr[point.Y] = utils.ReplaceStringAtIndex(regionStr, 'A', point)
    }

    // count top and bottom edges
    topIsEdge := false
    bottomIsEdge := false

    numTopEdges := 0
    numBottomEdges := 0

    for y, line := range regionStr {
        for x, letter := range line {
            if letter == '.' {
                topIsEdge = false
                bottomIsEdge = false
                continue
            }
            topLocation := utils.NewPair(x, y - 1)
            if !utils.InArrayBounds(strLen, strLen, topLocation) || utils.IndexString(regionStr, topLocation) != 'A' {
                if !topIsEdge {
                    numTopEdges++
                    topIsEdge = true
                }
            } else {
                topIsEdge = false
            }

            bottomLocation := utils.NewPair(x, y + 1)
            if !utils.InArrayBounds(strLen, strLen, bottomLocation) || utils.IndexString(regionStr, bottomLocation) != 'A' {
                if !bottomIsEdge {
                    numBottomEdges++
                    bottomIsEdge = true
                }
            } else {
                bottomIsEdge = false
            }
        }
        topIsEdge = false
        bottomIsEdge = false
    }

    // count side edges
    leftIsEdge := false
    rightIsEdge := false

    numLeftEdges := 0
    numRightEdges := 0


    for x, line := range regionStr {
        for y := range line {
            letter := regionStr[y][x]
            if letter == '.' {
                leftIsEdge = false
                rightIsEdge = false
                continue
            }
            leftLocation := utils.NewPair(x - 1, y)
            if !utils.InArrayBounds(strLen, strLen, leftLocation) || utils.IndexString(regionStr, leftLocation) != 'A' {
                if !leftIsEdge {
                    numLeftEdges++
                    leftIsEdge = true
                }
            } else {
                leftIsEdge = false
            }

            rightLocation := utils.NewPair(x + 1, y)
            if !utils.InArrayBounds(strLen, strLen, rightLocation) || utils.IndexString(regionStr, rightLocation) != 'A' {
                if !rightIsEdge {
                    numRightEdges++
                    rightIsEdge = true
                }
            } else {
                rightIsEdge = false
            }
        }
        leftIsEdge = false
        rightIsEdge = false
    }

    // fmt.Println(regionStr)
    // for _, line := range regionStr {
    //     fmt.Println(line)
    // }

    return len(*reg.seenLocations) * (numTopEdges + numBottomEdges + numLeftEdges + numRightEdges)
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n")

    regions := findAllRegions(splitData)
    price := 0
    bulkDiscount := 0

    for _, region := range regions {
        bulkDiscount += regionEdges(region, len(splitData))
    }

    // for _, region := range regions {
    //    fmt.Println(regionEdges(region, len(splitData)))
    // }

    for _, region := range regions {
        price += len(*region.seenLocations) * region.perimiter
    }

    fmt.Println("Part 1:", price)
    fmt.Println("Part 2:", bulkDiscount)
}
