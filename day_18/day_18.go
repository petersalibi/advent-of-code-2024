package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	_ "time"

	"advent-of-code/utils"
)

type item struct {
	cost int
	node utils.Pair
}

type PriorityQueue []item

func (pq *PriorityQueue) push(node utils.Pair, cost int) {
	*pq = append(*pq, item{cost: cost, node: node})
}

func (pq *PriorityQueue) pop() item {
	cmpFunc := func(a, b item) int {
		return a.cost - b.cost
	}
	lowestCost := slices.MinFunc(*pq, cmpFunc)

	for i := range *pq {
		if (*pq)[i].cost == lowestCost.cost {
			*pq = slices.Delete(*pq, i, i+1)
			break
		}
	}

	return lowestCost
}

func (pq *PriorityQueue) contains(node utils.Pair) bool {
	for _, pqItem := range *pq {
		if utils.Equal(node, pqItem.node) {
			return true
		}
	}
	return false
}

func pathfind(dataMap [][]byte, endPoint int) int {
	pq := make(PriorityQueue, 1, 200)
	pq[0] = item{cost: 0, node: utils.NewPair(0, 0)}
	seen := make([]utils.Pair, 1)
	seen[0] = utils.NewPair(0, 0)

	for len(pq) != 0 {
        // fmt.Println(pq)
		currentItem := pq.pop()
		current := currentItem.node
		seen = append(seen, currentItem.node)

		neighbors := []utils.Pair{
			{X: current.X, Y: current.Y + 1},
			{X: current.X, Y: current.Y - 1},
			{X: current.X + 1, Y: current.Y},
			{X: current.X - 1, Y: current.Y},
		}

		for _, neighbor := range neighbors {
			if pq.contains(neighbor) || slices.ContainsFunc(seen, utils.ContainsPair(neighbor)) ||
				!utils.InArrayBounds(len(dataMap), len(dataMap[0]), neighbor) {
				continue
			}
            if dataMap[neighbor.Y][neighbor.X] == '#' {
                continue
            }

			cost := currentItem.cost + 1
			if neighbor.X == endPoint && neighbor.Y == endPoint {
				return cost
			}
            pq.push(neighbor, cost)
		}
	}

	return -1
}

func main() {
	data := utils.GetDataFromFile()
	splitData := strings.Split(data, "\n")
    endPoint := 71

    dataMap := make([][]byte, endPoint)

    for i := range endPoint {
        dataMap[i] = make([]byte, endPoint)
        dataMap[i] = []byte(strings.Repeat(".", endPoint))
    }

    for i := range splitData {
        corruptData := strings.Split(splitData[i], ",")
        corruptX, _ := strconv.Atoi(corruptData[0])
        corruptY, _ := strconv.Atoi(corruptData[1])
        dataMap[corruptY][corruptX] = '#'
    }

    i := len(splitData) - 1
    for {
        corruptData := strings.Split(splitData[i], ",")
        corruptX, _ := strconv.Atoi(corruptData[0])
        corruptY, _ := strconv.Atoi(corruptData[1])
        dataMap[corruptY][corruptX] = '.'
        // for _, line := range dataMap {
        //     fmt.Println(string(line))
        // }
        // fmt.Println(i)

        shortestDistance := pathfind(dataMap, endPoint - 1)
        // fmt.Println(shortestDistance)
        // time.Sleep(50 * time.Millisecond)
        if shortestDistance > 0 {
            break
        }
        i--
    }

    fmt.Println("Part 2:", splitData[i])
}
