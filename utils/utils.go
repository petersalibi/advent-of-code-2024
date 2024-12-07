package utils

import (
    "os"
    "fmt"
    "strings"
)

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
