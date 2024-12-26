package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

var keypad = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'X', '0', 'A'},
}

var directionPad = [][]byte{
	{'X', '^', 'A'},
	{'<', 'v', '>'},
}

func findByte(keypad [][]byte, aVal byte) utils.Pair {
	for y, line := range keypad {
		for x, letter := range line {
			if aVal == letter {
				return utils.NewPair(x, y)
			}
		}
	}
	return utils.NewPair(0, 0)
}

func optimalKeypad(keypad [][]byte, code string, startPoint byte) []string {
	keyOrders := make([]string, 1)
	keyOrders[0] = ""
	lastPair := findByte(keypad, startPoint)
	for _, number := range code {
		order1 := ""
		order2 := ""
		end := findByte(keypad, byte(number))
		distance := utils.SubPair(lastPair, end)
		if distance.Y > 0 {
			order1 += strings.Repeat("^", distance.Y)
		} else {
			order1 += strings.Repeat("v", -distance.Y)
		}
		if distance.X > 0 {
			order2 = strings.Repeat("<", distance.X) + order1
			order1 += strings.Repeat("<", distance.X)
		} else {
			order2 = strings.Repeat(">", -distance.X) + order1
			order1 += strings.Repeat(">", -distance.X)
		}

		if keypad[end.Y][lastPair.X] == 'X' {
			order1 = order2
		}
		if keypad[lastPair.Y][end.X] == 'X' {
			order2 = order1
		}

		order1 += "A"
		order2 += "A"
		lastPair = end
		extendedCodes := make([]string, 0)
		for i := range keyOrders {
			if order1 == order2 {
				extendedCodes = append(extendedCodes, keyOrders[i]+order1)
			} else {
				extendedCodes = append(extendedCodes, keyOrders[i]+order1, keyOrders[i]+order2)
			}
		}
		keyOrders = extendedCodes
	}

	return keyOrders
}

var cache = make(map[string]int)

func findOptimal(x, y byte, depth int) int {
	inputString := fmt.Sprintf("%d,%d,%d", x, y, depth)
	if val, ok := cache[inputString]; ok {
		return val
	}

	optimalLength := math.MaxInt
	if depth == 1 {
		optimalLength = len(optimalKeypad(directionPad, string(x)+string(y), x)[0]) - 1
		fmt.Println(string(x), string(y), optimalKeypad(directionPad, string(x)+string(y), x), optimalLength)
		cache[inputString] = optimalLength
		return optimalLength
	}

	for _, code := range optimalKeypad(directionPad, string(x)+string(y), x) {
		codeLength := 0
		newCode := "A" + code
		for i, letter := range newCode[:len(newCode)-1] {
			nextLetter := newCode[i+1]
			codeLength += findOptimal(byte(letter), nextLetter, depth-1)
		}
		optimalLength = min(optimalLength, codeLength)
	}
	cache[inputString] = optimalLength - 1
	return optimalLength - 1
}

// func findComplexity(keypad, directions [][]byte, codes []string) int {
//     totalComplexity := 0
//     for _, keyCode := range codes {
//         fmt.Println("Code", keyCode)
//         var next []string
//         next = optimalKeypad(keypad, keyCode)
//         for range 2 {
//             possible_next := []string{}
//             for _, seq := range next {
//                 possible_next = append(possible_next, optimalKeypad(directions, seq)...)
//             }
//             next = possible_next
//         }
//
//         code := slices.MinFunc(next, func (a, b string) int {
//             return len(a) - len(b)
//         })
//
//         codeNum, err := strconv.Atoi(keyCode[:len(keyCode) - 1])
//         if err != nil {
//             fmt.Println(err)
//         }
//
//         totalComplexity += codeNum * len(code)
//     }
//
//     return totalComplexity
// }

func findComplexity(data []string, depth int) int {
    totalComplexity := 0
    for _, line := range data {
        optimalLength := math.MaxInt
        for _, code := range optimalKeypad(keypad, line, 'A') {
            fmt.Println(code)
            codeLength := 0
            newCode := "A" + code
            for i, letter := range newCode[:len(newCode)-1] {
                nextLetter := newCode[i+1]
                codeLength += findOptimal(byte(letter), nextLetter, depth)
            }
            optimalLength = min(optimalLength, codeLength)
        }
        codeNum, _ := strconv.Atoi(line[:len(line)-1])
        totalComplexity += optimalLength * codeNum
    }
    return totalComplexity
}

func main() {
	data := utils.GetDataFromFile()
	splitData := strings.Split(data, "\n")

	fmt.Println(splitData)
	// fmt.Println(findComplexity(keypad, directionPad, splitData))
	// fmt.Println(findOptimal('<','>', 2))
	fmt.Println(findComplexity(splitData, 25))
}
