package main

import (
	// "fmt"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatal("Usage: ./program filename")
    }

    data, err := os.ReadFile(os.Args[1])

    if err != nil {
        log.Fatal(err)
    }

    data_str := string(data)

    data_str = strings.Trim(data_str, "\n")

    data_list := strings.Split(data_str, "\n")

    left_list := make([]int, len(data_list))
    right_list := make([]int, len(data_list))

    left_map := make(map[int]int)

    for i, element := range(data_list) {
        fmt.Printf("line: %q\n", element)
        split_data := strings.Split(element, "   ")

        // fmt.Printf("%q\n", split_data)
        left_val, err := strconv.Atoi(split_data[0])
        if err != nil {
            log.Fatalf("Could not convert %s to int", split_data[0])
        }
        left_list[i] = left_val
        left_map[left_val] = 0
        right_list[i], err = strconv.Atoi(split_data[1])
        if err != nil {
            log.Fatalf("Could not convert %s to int", split_data[1])
        }
    }

    slices.Sort(left_list)
    slices.Sort(right_list)

    var sum_result = 0

    for i, elem := range(left_list) {
        sum_result += int(math.Abs(float64(elem - right_list[i])))
    }

    for _, elem := range(right_list) {
        _, ok := left_map[elem]
        if !ok {
            continue
        }

        left_map[elem]++
    }

    var freq_result = 0

    for key, value := range(left_map) {
        freq_result += key * value
    }

    fmt.Println(sum_result)
    fmt.Println(freq_result)
}
