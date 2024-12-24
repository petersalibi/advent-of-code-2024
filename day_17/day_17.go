package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

type vm struct {
    regA int64
    regB int64
    regC int64

    instructions []string
    output []string

    ip int
}

func newVM(a, b, c int64, instructions []string) vm {
    return vm{regA: a, regB: b, regC: c, instructions: instructions, ip: 0}
}

func readComboOperand(vm vm, operand int64) int64 {
    if operand <= 3 {
        return operand
    }

    switch operand {
    case 4: return vm.regA
    case 5: return vm.regB
    case 6: return vm.regC
    default: return 0
    }
}

func runVM(vm *vm) {
    for vm.ip < len(vm.instructions) {
        ins, _ := strconv.Atoi(vm.instructions[vm.ip])
        operandi32, _ := strconv.Atoi(vm.instructions[vm.ip + 1])
        operand := int64(operandi32)
        switch ins {
        case 0:
            vm.regA = vm.regA >> readComboOperand(*vm, operand)
            vm.ip += 2
        case 1:
            vm.regB = vm.regB ^ operand
            vm.ip += 2
        case 2:
            vm.regB = readComboOperand(*vm, operand) % 8
            vm.ip += 2
        case 3:
            if vm.regA != 0 {
                vm.ip = int(operand)
            } else {
                vm.ip += 2
            }
        case 4:
            vm.regB = vm.regB ^ vm.regC
            vm.ip += 2
        case 5:
            vm.output = append(vm.output, strconv.Itoa(int(readComboOperand(*vm, operand) % 8)))
            vm.ip += 2
        case 6:
            vm.regB = vm.regA >> readComboOperand(*vm, operand)
            vm.ip += 2
        case 7:
            vm.regC = vm.regA >> readComboOperand(*vm, operand)
            vm.ip += 2
        }
    }
    vm.ip = 0
}

func findQuine(vm *vm, regA int64, index int, output []string) int64 {
    if index < 0 {
        return regA
    }

    for j := 0; j <= 7; j++ {
        newRegister := (regA << 3) + int64(j)
        vm.regA = newRegister
        vm.output = []string{}
        runVM(vm)
        if vm.output[0] == output[index] {
            result := findQuine(vm, newRegister, index - 1, output)

            if result >= 0 {
                return result
            }
            continue
        }
    }
    return -1
}

func main() {
    _ = utils.GetDataFromFile()
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n\n")

    programString := strings.Split(splitData[1], " ")
    program := strings.Split(programString[1], ",")

    registers := strings.Split(splitData[0], "\n")
    regA, _ := strconv.ParseInt(strings.Split(registers[0], " ")[2], 10, 64)
    regB, _ := strconv.ParseInt(strings.Split(registers[1], " ")[2], 10, 64)
    regC, _ := strconv.ParseInt(strings.Split(registers[2], " ")[2], 10, 64)


    // regA, regB, regC := 4, 2024, 43690
    // fmt.Println(regA, regB, regC)
    //
    // programString := "0,2"
    // program := strings.Split(programString, ",")

    problemVM := newVM(regA, regB, regC, program)
    // 1: 4, 10: 6, 11: 7, 100: 0, 101: 1, 110: 2, 111: 3
    // problemVM.regA = 0b110
    problemVM.output = []string{}
    problemVM.regA = 164278764924605
    runVM(&problemVM)
    fmt.Println("Part 1:", strings.Join(problemVM.output, ","))
    fmt.Println("Part 2:", findQuine(&problemVM, 0, len(program) - 1, program))
}
