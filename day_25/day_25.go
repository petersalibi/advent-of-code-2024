package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func getHeights(schem []string) []int {
    colHeights := make([]int, len(schem[0]))

    for x := range(len(schem[0])) {
        curCol := 0
        for y := range(len(schem)) {
            if schem[y][x] == '#' {
                curCol++
            }
        }
        colHeights[x] = curCol - 1
    }
    return colHeights
}

func compareLockKey(lock, key []int) int {
    if len(lock) != len(key) {
        return 0
    }

    for i := range lock {
        if lock[i] + key[i] > 5 {
            return 0
        }
    }

    return 1
}

func main() {
    data := utils.GetDataFromFile()
    locksAndKeys := strings.Split(data, "\n\n")
    locks := make([][]int, 0)
    keys := make([][]int, 0)

    for _, schem := range locksAndKeys {
        schem := strings.Split(schem, "\n")
        if schem[0][0] == '#' {
            locks = append(locks, getHeights(schem))
        } else {
            keys = append(keys, getHeights(schem))
        }
    }


    total := 0

    for _, lock := range locks {
        for _, key := range keys {
            total += compareLockKey(lock, key)
        }
    }

    fmt.Println(locks)
    fmt.Println(keys)
    fmt.Println(total)
}
