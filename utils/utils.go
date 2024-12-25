package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
    X int
    Y int
}

func NewPair(x, y int) Pair {
    return Pair{X: x, Y: y}
}

func AddPair(p1, p2 Pair) Pair {
    return Pair{X: p1.X + p2.X, Y: p1.Y + p2.Y}
}

func SubPair(p1, p2 Pair) Pair {
    return Pair{X: p1.X - p2.X, Y: p1.Y - p2.Y}
}

func Equal(p1, p2 Pair) bool {
    return p1.X == p2.X && p1.Y == p2.Y
}

func InArrayBounds(xLimit, yLimit int, point Pair) bool {
    if point.X < 0 || point.Y < 0 {
        return false
    }

    if point.X >= xLimit || point.Y >= yLimit {
        return false
    }

    return true
}

func (p *Pair) HashKey() string {
    return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

func IndexString(input []string, p Pair) byte {
    return input[p.Y][p.X]
}

func IndexArray[s [][]E, E any](arr s, p Pair) E {
    return arr[p.Y][p.X]
}

func ContainsPair(p1 Pair) func(Pair) bool {
    return func (p2 Pair) bool {
        return Equal(p1, p2)
    }
}

func SortPair(p1, p2 Pair) int {
    if p1.X > p2.X {
        return 1
    } else if p1.X < p2.X {
        return -1
    }
    if p1.Y > p2.Y {
        return 1
    } else if p1.Y < p2.Y {
        return -1
    }
    return 0
}

func Abs(a int) int {
    if a < 0 {
        return -a
    }

    return a
}

func ReplaceStringAtIndex(str []string, letter rune, index Pair) string {
	newStr := []rune(str[index.Y])
	newStr[index.X] = letter
	return string(newStr)
}

func ByteArrayToString(b [][]byte) string {
	newStr := make([]string, len(b))
	for i, line := range b {
		newStr[i] = string(line)
	}

    return strings.Join(newStr, "\n")
}

func GetDataFromFile() string {
    if (len(os.Args) != 2) {
        fmt.Fprintf(os.Stderr, "Usage: ./program file")
        return ""
    }

    data, err := os.ReadFile(os.Args[1])

    if err != nil {
        fmt.Fprint(os.Stderr, err)
        return ""
    }

    dataStr := string(data)
    dataStr = strings.Trim(dataStr, "\n")

    return dataStr
}

func GetDataBytesFromFile() []byte {
    if (len(os.Args) != 2) {
        fmt.Fprintf(os.Stderr, "Usage: ./program file")
        return []byte{0}
    }

    data, err := os.ReadFile(os.Args[1])

    if err != nil {
        fmt.Fprint(os.Stderr, err)
        return []byte{0}
    }

    return data
}
