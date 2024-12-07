package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type compareFunc func(left, right int) bool

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func convertStringToIntSlice(input string) ([]int, error) {
    valuesStr := strings.Split(input, " ")

    values := make([]int, len(valuesStr))

    for i, val := range(valuesStr) {
        valInt, err := strconv.ParseInt(val, 10, 32)

        if err != nil {
            return nil, err
        }

        values[i] = int(valInt)
    }

    return values, nil
}

func determineSafe(values []int, dampener bool) bool {
    isDecreasing := func (left, right int) bool { return left > right }
    isIncreasing := func (left, right int) bool { return left < right }

    var compare compareFunc

    {
        curVal := values[1]
        prevVal := values[0]
        if isIncreasing(prevVal, curVal) {
            compare = isIncreasing
        } else {
            compare = isDecreasing
        }
    }

    for i, val := range(values[1:]) { // i is offset by 1 because we access the slice from index 1
        curVal := val
        prevVal := values[i]
        if !compare(prevVal, curVal) || abs(prevVal - curVal) > 3 {
            if dampener {
                for i := range(values) {
                    var valCopy = make([]int, len(values))

                    copy(valCopy, values)

                    valCopy = slices.Delete(valCopy, i, i + 1)
                    if determineSafe(valCopy, false) {
                        return true
                    }
                }
            }
            return false
        }
    }

    return true
}

func getNumSafe(data []string, ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    result := 0
    for _, line := range(data) {
        values, err := convertStringToIntSlice(line)
        if err != nil {
            break
        }
        if determineSafe(values, true) {
            result++
        }
    }
    ch <- result
}

func dispatchSafe(data []string, numThreads int, ch chan int) {
    var wg sync.WaitGroup
    interval := len(data) / numThreads
    for i := 0; i < numThreads; i++ {
        start := i * interval
        var end int
        if i == numThreads - 1 {
            end = len(data)
        } else {
            end = (i + 1) * interval
        }

        wg.Add(1)
        go getNumSafe(data[start:end], ch, &wg)
    }

    wg.Wait()
    close(ch)
}

func main() {
    if (len(os.Args) != 2) {
        fmt.Fprintf(os.Stderr, "Usage: ./program file")
        return
    }

    data, err := os.ReadFile(os.Args[1])

    if err != nil {
        fmt.Fprint(os.Stderr, err)
        return
    }

    dataStr := string(data)
    dataStr = strings.Trim(dataStr, "\n")

    dataLines := strings.Split(dataStr, "\n")

    ch := make(chan int)

    go dispatchSafe(dataLines, 4, ch)
    result := 0

    for val := range(ch) {
        result += val
    }

    fmt.Println(result)
}
