package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

const ABUTTON = 3
const BBUTTON = 1

var coords = regexp.MustCompile(`(\d+).*\+(\d+)`)
var prizeCoordRegex = regexp.MustCompile(`(\d+).*=(\d+)`)

func findPrizeTokens(machine string, part1 bool) int {
    splitMachine := strings.Split(machine, "\n")

    aCoords := coords.FindStringSubmatch(splitMachine[0])
    a, _ := strconv.Atoi(aCoords[1])
    c, _ := strconv.Atoi(aCoords[2])

    bCoords := coords.FindStringSubmatch(splitMachine[1])
    b, _ := strconv.Atoi(bCoords[1])
    d, _ := strconv.Atoi(bCoords[2])

    prizeCoord := prizeCoordRegex.FindStringSubmatch(splitMachine[2])
    px, _ := strconv.Atoi(prizeCoord[1])
    py, _ := strconv.Atoi(prizeCoord[2])

    if !part1 {
        px = px + 10000000000000
        py = py + 10000000000000
    }

    det := a*d - b*c
    if det == 0 {
        return 0
    }

    x := px * d - py * b
    if x % det != 0 {
        return 0
    }
    aPresses := x / det

    y := px * -c + py * a
    if y % det != 0 {
        return 0
    }
    bPresses := y / det

    // fmt.Println(machine)
    // fmt.Println("", a, b, "\n", c, d)
    // fmt.Println("", px, "\n", py)
    // fmt.Println(aPresses, bPresses)

    return aPresses * ABUTTON + bPresses * BBUTTON
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n\n")
    totalTokens := 0
    totalTokensPart2 := 0

    for _, machine := range splitData {
        totalTokens += findPrizeTokens(machine, true)
    }

    for _, machine := range splitData {
        totalTokensPart2 += findPrizeTokens(machine, false)
    }

    fmt.Println("Part 1:", totalTokens)
    fmt.Println("Part 2:", totalTokensPart2)
}
