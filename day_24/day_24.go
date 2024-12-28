package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getBinExpr(binOp string) func (int, int) int {
    switch binOp {
    case "AND":
        return func(a, b int) int {
            return a & b
        }
    case "OR":
        return func(a, b int) int {
            return a | b
        }
    case "XOR":
        return func(a, b int) int {
            return a ^ b
        }
    }
    return func (int, int) int { return 0 }
}

func calcBinFunc(binExprs map[string]string, wires map[string]int, result string) int {
    if wireVal, ok := wires[result]; ok {
        return wireVal
    }
    wireExpr := strings.Split(binExprs[result], " ")

    binFunc := getBinExpr(wireExpr[1])

    left := calcBinFunc(binExprs, wires, wireExpr[0])
    right := calcBinFunc(binExprs, wires, wireExpr[2])

    wires[result] = binFunc(left, right)
    return binFunc(left, right)
}

func expressionValue(values map[string]string, a, op, b string) (string, bool) {
    res, ok := values[a + " " + op + " " + b]
    if !ok {
        res, ok = values[b + " " + op + " " + a]
    }
    return res, ok
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n\n")

    setWires, calcWires := splitData[0], splitData[1]

    wires := make(map[string]int, 0)

    for _, wire := range strings.Split(setWires, "\n") {
        wireValues := strings.Split(wire, ": ")
        wireBool, _ := strconv.Atoi(wireValues[1])
        wires[wireValues[0]] = wireBool
    }

    wireExpressions := make(map[string]string, 0)
    ExpressionValues := make(map[string]string)

    for _, binExpr := range strings.Split(calcWires, "\n") {
        binExpr := strings.Replace(binExpr, " ->", "", 1)
        wireExpr := strings.Split(binExpr, " ")
        binExpr = binExpr[:len(binExpr) - 4]
        wireExpressions[wireExpr[3]] = binExpr
        ExpressionValues[binExpr] = wireExpr[3]
    }

    for wire := range wireExpressions {
        calcBinFunc(wireExpressions, wires, wire)
    }
    fmt.Println(wireExpressions)

    a := "x00"
    b := "y00"
    s := "z00"
    // prev_axb := ""
    // prev_anb := ""
    // prev_cnres := ""
    cin := ""
    cout := "wbd"


    for i := 1; i <= 44; i++ {
        a = fmt.Sprintf("x%02d", i)
        b = fmt.Sprintf("y%02d", i)
        axb, ok := expressionValue(ExpressionValues, a, "XOR", b)
        if !ok {
            fmt.Println("How?")
            break
        }
        cin = cout

        s, ok = expressionValue(ExpressionValues, axb, "XOR", cin)
        fmt.Println(s)
        if !ok {
            fmt.Printf("res xor cin not found: %s XOR %s instead\n", axb, cin)
            break
        }

        anb, ok := expressionValue(ExpressionValues, a, "AND", b)
        if !ok {
            fmt.Printf("a AND b not found: %s AND %s instead\n", a, b)
            break
        }

        cnres, ok := expressionValue(ExpressionValues, cin, "AND", axb)
        if !ok {
            fmt.Printf("res and cin not found: %s AND %s instead\n", cin, axb)
            break
        }

        cout, ok = expressionValue(ExpressionValues, anb, "OR", cnres)
        if !ok {
            fmt.Printf("anb OR cnres not found: %s OR %s instead\n", anb, cnres)
            break
        }
        // prev_axb, prev_anb, prev_cnres = axb, anb, cnres
    }

    x := 0
    for i := range 99 {
        wireName := fmt.Sprintf("x%02d", i)
        val, ok := wires[wireName]
        if !ok {
            break
        }

        x += val << i
    }
    y := 0
    for i := range 99 {
        wireName := fmt.Sprintf("y%02d", i)
        val, ok := wires[wireName]
        if !ok {
            break
        }

        y += val << i
    }
    z := 0
    for i := range 99 {
        wireName := fmt.Sprintf("z%02d", i)
        val, ok := wires[wireName]
        if !ok {
            break
        }

        z += val << i
    }

    fmt.Println(x, y, x + y, z)
    outString := strings.Split("nbc,svm,kqk,z15,cgq,z23,fnr,z39", ",")
    slices.Sort(outString)
    fmt.Println(strings.Join(outString, ","))
}
