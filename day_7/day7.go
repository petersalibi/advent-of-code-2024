package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

func first[S ~[]E, E any](s S) E {
    return s[0]
}

func rest[S ~[]E, E any](s S) []E {
    return s[1:]
}

func concatInts(a, b int) int {
    aStr := strconv.Itoa(a)
    bStr := strconv.Itoa(b)

    concatStr := aStr + bStr

    concatInt, err := strconv.Atoi(concatStr)

    if err != nil {
        return 0
    }

    return concatInt
}

func checkAllOperations(want, runningTotal int, rawValues []string, part1 bool) bool {
    if len(rawValues) == 0 {
        return want == runningTotal
    }
    val, err := strconv.Atoi(first(rawValues))

    if err != nil {
        return false
    }

    valuesRest := rest(rawValues)

    result :=  checkAllOperations(want, runningTotal + val, valuesRest, part1) ||
               checkAllOperations(want, runningTotal * val, valuesRest, part1)
    if part1 {
        return result
    } else {
        return result || checkAllOperations(want, concatInts(runningTotal, val), valuesRest, part1)
    }
}

func checkLine(line string, part1 bool) int {
    calibrations := strings.Split(line, ": ")
    if len(calibrations) != 2 {
        return 0
    }

    calibrationValue, err := strconv.Atoi(calibrations[0])

    if err != nil {
        return 0
    }

    operands := strings.Split(calibrations[1], " ")

    firstVal, err := strconv.Atoi(first(operands))

    if err != nil {
        return 0
    }

    if checkAllOperations(calibrationValue, firstVal, rest(operands), part1) {
        return calibrationValue
    } else {
        return 0
    }
}

func checkAllLines(input []string, part1 bool) int {
    result := 0
    for _, line := range input {
        result += checkLine(line, part1)
    }

    return result
}

func main() {
	data := utils.GetDataFromFile()

	splitData := strings.Split(data, "\n")

    fmt.Printf("Part 1: %d\n", checkAllLines(splitData, true))
    fmt.Printf("Part 2: %d\n", checkAllLines(splitData, false))
}
