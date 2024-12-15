package main

import (
	"fmt"
	"slices"

	"advent-of-code/utils"
)

type fileBlock struct {
	id int

	indices []int
}

func (b *fileBlock) add(index int) {
    b.indices = append(b.indices, index)
}

// byte char to number
func bton(b byte) byte {
    return b - 0x30
}

func calcCheckSum(blocks []fileBlock) int {
    checkSum := 0

    for _, block := range blocks {
        for _, idx := range block.indices {
            checkSum += block.id * idx
        }
    }

    return checkSum
}

func part1(data []byte) int {
    blocks := make([]fileBlock, len(data) / 2 + 1)
    idFile := 0
    fileSizeLeft := 0
    var freeChunkSizeLeft byte = 0
    writingFree := false

    for i, j := 0, len(data); i < j; {
        leftBlock := &blocks[i / 2]
        leftBlock.id = i / 2
        rightBlock := &blocks[j / 2]
        rightBlock.id = j / 2

        if fileSizeLeft == 0 {
            fileSizeLeft = int(bton(data[j - 1]))
        }

        if i / 2 == j / 2 && !writingFree {
            for range fileSizeLeft {
                leftBlock.add(idFile)
                idFile--
            }
            break
        }

        if !writingFree {
            for range bton(data[i]) {
                leftBlock.add(idFile)
                idFile++
            }
            writingFree = true
            i++
            freeChunkSizeLeft = bton(data[i])
            i++
        }

        for freeChunkSizeLeft > 0 && writingFree {
            rightBlock.add(idFile)
            idFile++
            fileSizeLeft--
            freeChunkSizeLeft--
            if fileSizeLeft == 0 {
                j -= 2
                break
            }

        }

        if freeChunkSizeLeft == 0 {
            writingFree = false
        }
    }

    return calcCheckSum(blocks)
}

func splitSlice(splitIndex int, freeIndices []int) ([]int, []int) {
    return freeIndices[:splitIndex], freeIndices[splitIndex:]
}

func joinAdjacentBlocks(freeBlocks []fileBlock) {
    for i := 0; i + 1 < len(freeBlocks); i++ {
        firstBlock := &freeBlocks[i]
        nextBlock := &freeBlocks[i+1]
        if len(firstBlock.indices) == 0 || len(nextBlock.indices) == 0 {
            continue
        }

        firstBlockEnd := firstBlock.indices[len(firstBlock.indices) - 1]
        nextBlockStart := nextBlock.indices[0]

        if firstBlockEnd + 1 == nextBlockStart {
            firstBlock.indices = slices.Concat(firstBlock.indices, nextBlock.indices)
            nextBlock.indices = []int{}
        }
    }
}

func part2(data []byte) int {
    files := make([]fileBlock, len(data) / 2 + 1)
    freeSpace := make([]fileBlock, len(data) / 2 + 1)

    idFile := len(data) / 2
    idFree := 0

    index := 0

    for i, size := range data {
        var block *fileBlock
        if i % 2 == 0 {
            block = &files[idFile]
            block.id = i / 2
            idFile--
        } else {
            block = &freeSpace[idFree]
            idFree++
        }

        for range bton(size) {
            block.indices = append(block.indices, index)
            index++
        }
    }

    outer:
    for i := range files {
        block := &files[i]
        if len(block.indices) == 0 {
            continue
        }
        for j := range freeSpace {
            freeBlock := &freeSpace[j]
            if len(block.indices) <= len(freeBlock.indices) {
                if block.indices[0] < freeBlock.indices[0] {
                    break
                }
                block.indices, freeBlock.indices = splitSlice(len(block.indices), freeBlock.indices)
                continue outer
            }
        }
    }

    return calcCheckSum(files)
}

func main() {
	data := utils.GetDataBytesFromFile()
    // remove endline
    data = data[:len(data) - 1]

    checkSumPart1 := part1(data)
    checkSumPart2 := part2(data)

    fmt.Println("Part 1:", checkSumPart1)
    fmt.Println("Part 2:", checkSumPart2)

}
