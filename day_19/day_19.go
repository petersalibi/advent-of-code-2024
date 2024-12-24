package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

var cache = make(map[string]int)

func checkDesign(design string, towels []string) int {
    if design == "" { return 1 }

    if permPossible, ok := cache[design]; ok {
        return permPossible
    }

    possibleDesign := 0
    for _, towel := range towels {
        if len(design) < len(towel) { continue }
        if strings.Contains(design[:len(towel)], towel) {
            possibleDesign += checkDesign(design[len(towel):], towels)
            continue
        }
    }

    cache[design] = possibleDesign

    return possibleDesign
}

func checkAll(designs, towels []string) int {
    numPossible := 0
    for _, design := range designs {
        numPossible += checkDesign(design, towels)
    }
    return numPossible
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n\n")

    substrings := strings.Split(splitData[0], ", ")
    designs := strings.Split(splitData[1], "\n")

    fmt.Println(checkAll(designs, substrings))
}
