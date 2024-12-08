package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func addAfter(mapping []string, afterMap map[string][]string) {
    if len(mapping) != 2 {
        return
    }

    origList := afterMap[mapping[0]]
    newElem := mapping[1]

    afterMap[mapping[0]] = append(origList, newElem)
}

func addAllToAfter(input []string, afterMap map[string][]string) {
    for _, line := range input {
        mapping := strings.Split(line, "|")
        addAfter(mapping, afterMap)
    }
}

func isValidSequence(pageNumbers []string, afterMap map[string][]string) bool {
    readNumbers := make([]string, len(pageNumbers))
    for i, number := range pageNumbers {
        afterNumbers, ok := afterMap[number]
        if !ok {
            readNumbers[i] = number
            continue
        }
        for _, afterNum := range afterNumbers {
            if slices.Contains(readNumbers, afterNum) {
                return false
            }
        }
        readNumbers[i] = number
    }
    return true
}

func findMedian(pageSequence []string) int {
    index := len(pageSequence) / 2
    median, _ := strconv.Atoi(pageSequence[index])
    return median
}

func checkAllValid(pageNumbersList []string, afterMap map[string][]string, part1 bool) int {
    // a should be ahead of b if a is in the afterlist of b
    sortFunction := func (a, b string) int {
        afterList, ok := afterMap[b]
        if !ok {
            return -1
        }

        if slices.Contains(afterList, a) {
            return 1
        } else {
            return -1
        }
    }
    result := 0
    for _, pageNumbers := range pageNumbersList {
        pageSequence := strings.Split(pageNumbers, ",")
        if isValidSequence(pageSequence, afterMap) {
            if !part1 {
                continue
            }
            result += findMedian(pageSequence)
        } else if !part1 {
            slices.SortFunc(pageSequence, sortFunction)
            result += findMedian(pageSequence)
        }
    }
    return result
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n\n")
    rules := splitData[0]
    splitRules := strings.Split(rules, "\n")
    afterMap := make(map[string][]string)
    addAllToAfter(splitRules, afterMap)

    pageNumbers := splitData[1]
    splitPageNumbers := strings.Split(pageNumbers, "\n")

    fmt.Printf("Part 1: %d\n", checkAllValid(splitPageNumbers, afterMap, true))
    fmt.Printf("Part 2: %d\n", checkAllValid(splitPageNumbers, afterMap, false))

    // for key, value := range afterMap {
    //     fmt.Printf("%s: %q\n", key, value)
    // }
}
