package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"

	// "strconv"
	"strings"
)

const pruneop = 16777216

func genSecretNumber(secret, iterations int) []int {
    secretSeq := make([]int, 1)
    secretSeq[0] = secret % 10
    for range iterations {
        newNum := secret * 64
        // mix
        secret = newNum ^ secret
        // prune
        secret = secret % pruneop

        newNum = secret / 32
        // mix
        secret = newNum ^ secret
        // prune
        secret = secret % pruneop

        newNum = secret * 2048
        // mix
        secret = newNum ^ secret
        // prune
        secret = secret % pruneop
        secretSeq = append(secretSeq, secret % 10)
    }
    return secretSeq
}

func generateSequences() [][]int {
    seqs := make([][]int, 0)

    for x := range 19 {
        for y := range 19 {
            if utils.Abs(x - 9 + y - 9) > 9 { continue }
            for z := range 19 {
                if utils.Abs(x - 9 + y - 9 + z - 9) > 9 { continue }
                if utils.Abs(y - 9 + z - 9) > 9 { continue }
                for w := range 10 {
                    if utils.Abs(x - 9 + y - 9 + z - 9 + w - 9) > 9 { continue }
                    if utils.Abs(y - 9 + z - 9 + w - 9) > 9 { continue }
                    if utils.Abs(z - 9 + w - 9) > 9 { continue }
                    seqs = append(seqs, []int{x - 9, y - 9, z - 9, w})
                }
            }
        }
    }

    return seqs
}

func findSellPrice(priceDifSeq, sellSeq []int) int {
    for i, x := range priceDifSeq[:len(priceDifSeq) - 4] {
        y, z, w := priceDifSeq[i + 1], priceDifSeq[i + 2], priceDifSeq[i + 3]

        if x == sellSeq[0] && y == sellSeq[1] && z == sellSeq[2] && w == sellSeq[3] {
            // if i + 4 > 2000 { return - 1 }
            return i + 4
        }
    }

    return -1
}

func getPriceDif(prices []int) []int {
    prevCost := prices[0]
    priceDif := make([]int, 0)
    for _, cost := range prices[1:] {
        priceDif = append(priceDif, cost - prevCost)
        prevCost = cost
    }

    return priceDif
}

func findTotalSell(sequences, priceDifs [][]int, sellSeq []int) int {
    totalSell := 0
    for i, seq := range sequences {
        seqDif := priceDifs[i]
        sellIndex := findSellPrice(seqDif, sellSeq)
        if sellIndex == -1 { continue }
        totalSell += seq[sellIndex]
    }

    return totalSell
}

func findBestSequenceSell(sequences, priceDifs, sellSeqs [][]int) int {
    best := 0
    // var bestSeq []int

    for i, sellSeq := range sellSeqs {
        fmt.Println("left:", len(sellSeqs) - i)
        numBananas := findTotalSell(sequences, priceDifs, sellSeq)


        // if numBananas > best {
        //     best = numBananas
        //     bestSeq = sellSeq
        // }
        best = max(best, numBananas)
    }
    // fmt.Println(bestSeq)
    return best
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n")

    priceSeqs := make([][]int, 0)

    for _, secretNumber := range splitData {
        num, _ := strconv.Atoi(secretNumber)
        priceSeqs = append(priceSeqs, genSecretNumber(num, 2000))
    }

    priceDifs := make([][]int, 0)

    for _, seq := range priceSeqs {
        priceDifs = append(priceDifs, getPriceDif(seq))
    }

    // fmt.Println(getPriceDif(genSecretNumber(1, 2000)))
    // fmt.Println(sequenceSell(genSecretNumber(1, 2000), []int{-2, -6, 3, -1}))
    // testSeq := [][]int {
    //     {0, -1, 2, -3},
    //     {-2, 1, -1, 3},
    // }
    sellSeq := generateSequences()
    fmt.Println(len(sellSeq))
    fmt.Println(findBestSequenceSell(priceSeqs, priceDifs, sellSeq))
}
