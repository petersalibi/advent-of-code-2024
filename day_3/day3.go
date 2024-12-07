package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
    x int
    y int
}

func calcMuls(data string, enableConditionals bool) int {
    findMul := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

    muls := findMul.FindAllString(data, -1)

    pairs := make([]Pair, len(muls))

    pairRegex := regexp.MustCompile(`\d+,\d+`)
    pairIdx := 0

    enabled := true

    for _, str := range(muls) {
        switch str {
        case "do()":
            enabled = true
        case "don't()":
            enabled = false
        default:
            if enabled || !enableConditionals {
                pairStr := pairRegex.FindString(str)
                numPair := strings.Split(pairStr, ",")
                left, _ := strconv.Atoi(numPair[0])
                right, _ := strconv.Atoi(numPair[1])

                pairs[pairIdx] = Pair{left, right}
                pairIdx++
            }
        }
    }

    result := 0

    for i := 0; i < pairIdx; i++ {
        result += pairs[i].x * pairs[i].y
    }

    return result
}

func main() {
    data := utils.GetDataFromFile()

    result := calcMuls(data, false)
    fmt.Printf("Part 1: %d\n", result)
    result = calcMuls(data, true)
    fmt.Printf("Part 2: %d\n", result)
}


