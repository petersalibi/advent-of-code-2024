package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func blink(initialStones []int, blinks int, cacheTable map[utils.Pair]int) int {
    stones := make([]int, 0, 500000)

    stones = append(stones, initialStones...)
    numStones := 0

    for _, stone := range stones {
        numStones += computeStoneNums(stone, blinks, cacheTable)
    }

    return numStones
}

func computeStoneNums(stone, stepsLeft int, cacheTable map[utils.Pair]int) int {
    if stepsLeft == 0 {
        return 1
    }

    curStone := utils.Pair{X: stone, Y: stepsLeft}
    cachedVal, ok := cacheTable[curStone]
    if ok {
        return cachedVal
    }

    numDigits := int(math.Log10(float64(stone))) + 1
    switch {
    case stone == 0:
        newStonePair := utils.Pair{X: 1, Y: stepsLeft - 1}
        val, ok := cacheTable[newStonePair]

        if ok {
            cachedVal = val
            break
        }

        val = computeStoneNums(1, stepsLeft - 1, cacheTable)
        cacheTable[newStonePair] = val
        cachedVal = val
    case numDigits % 2 == 0:
        leftStone, rightStone := splitInt(stone, numDigits / 2)
        newLeftStonePair := utils.Pair{X: leftStone, Y: stepsLeft - 1}
        newRightStonePair := utils.Pair{X: rightStone, Y: stepsLeft - 1}
        leftVal, ok := cacheTable[newLeftStonePair]

        if !ok {
            leftVal = computeStoneNums(leftStone, stepsLeft - 1, cacheTable)
            cacheTable[newLeftStonePair] = leftVal
        }

        rightVal, ok := cacheTable[newRightStonePair]
        if !ok {
            rightVal = computeStoneNums(rightStone, stepsLeft - 1, cacheTable)
            cacheTable[newRightStonePair] = rightVal
        }
        cachedVal = leftVal + rightVal
    default:
        newStonePair := utils.Pair{X: stone*2024, Y: stepsLeft - 1}
        val, ok := cacheTable[newStonePair]

        if ok {
            cachedVal = val
            break
        }

        val = computeStoneNums(stone*2024, stepsLeft - 1, cacheTable)
        cacheTable[newStonePair] = val
        cachedVal = val
    }
    cacheTable[curStone] = cachedVal
    return cachedVal
}

func splitInt(number, splitIndex int) (int, int) {
    splitPow := int(math.Pow10(splitIndex))
    leftSplit := number / splitPow
    rightSplit := number - (leftSplit * splitPow)

    return leftSplit, rightSplit
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, " ")
    dataInts := make([]int, len(splitData))
    var err error

    for i, num := range splitData {
        dataInts[i], err = strconv.Atoi(num)

        if err != nil {
            fmt.Fprint(os.Stderr, err)
            return
        }
    }

    cacheTable := make(map[utils.Pair]int)

    fmt.Println("Part 1:", blink(dataInts, 25, cacheTable))
    fmt.Println("Part 2:", blink(dataInts, 75, cacheTable))
}
