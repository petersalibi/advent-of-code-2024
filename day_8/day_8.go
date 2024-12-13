package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strings"
)

func manhattanDistance(p1, p2 utils.Pair) utils.Pair {
    x := p1.X - p2.X
    y := p1.Y - p2.Y
    return utils.Pair{X: x, Y: y}
}

func mapAntennaLocations(input []string) map[rune][]utils.Pair {
    antennaMap := make(map[rune][]utils.Pair)

    for y, line := range input {
        for x, letter := range line {
            if letter == '.' {
                continue
            }
            antennaMap[letter] = append(antennaMap[letter], utils.Pair{X: x, Y: y})
        }
    }

    return antennaMap
}

func inBounds(xBound, yBound int, point utils.Pair) bool {
    if point.X < 0 || point.Y < 0 {
        return false
    }

    if point.X >= xBound || point.Y >= yBound {
        return false
    }

    return true
}

func findAntinodes(xBound, yBound int, antennaMap map[rune][]utils.Pair, part1 bool) []utils.Pair {
    antinodes := make([]utils.Pair, 0)

    for _, antenna := range antennaMap {
        for _, firstAntenna := range antenna {
            for _, secondAntenna := range antenna {
                if utils.Equal(firstAntenna, secondAntenna) {
                    continue
                }

                dist := manhattanDistance(secondAntenna, firstAntenna)
                antinode := utils.Pair{}
                if part1 {
                    antinode.X = firstAntenna.X + 2 * dist.X
                    antinode.Y = firstAntenna.Y + 2 * dist.Y
                } else {
                    antinode.X = firstAntenna.X + 1 * dist.X
                    antinode.Y = firstAntenna.Y + 1 * dist.Y
                }

                if part1 && slices.ContainsFunc(antinodes, func (p utils.Pair) bool {
                    return utils.Equal(p, antinode)
                }) {
                    continue
                }

                if !inBounds(xBound, yBound, antinode) {
                    continue
                }
                if part1 {
                    antinodes = append(antinodes, antinode)
                    continue
                }
                for inBounds(xBound, yBound, antinode) {
                    for slices.ContainsFunc(antinodes, func (p utils.Pair) bool {
                        return utils.Equal(p, antinode)
                    }) {
                        antinode.X += dist.X
                        antinode.Y += dist.Y
                    }
                    if !inBounds(xBound, yBound, antinode) {
                        break
                    }
                    antinodes = append(antinodes, antinode)
                    antinode.X += dist.X
                    antinode.Y += dist.Y
                }
            }
        }
    }

    return antinodes
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n")
    antennaMap := mapAntennaLocations(splitData)
    yBound := len(splitData)
    xBound := len(splitData[0])

    antinodes := findAntinodes(xBound, yBound, antennaMap, true)
    fmt.Printf("Part 1: %d\n", len(antinodes))
    antinodesResonance := findAntinodes(xBound, yBound, antennaMap, false)
    fmt.Printf("Part 2: %d\n", len(antinodesResonance))
}
