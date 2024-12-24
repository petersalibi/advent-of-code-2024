package main

import (
	"fmt"
	"slices"
	"strings"

	"advent-of-code/utils"
)

// An Item is something we manage in a priority queue.
// type Item struct {
// 	x int
//     y int
//     dx int
//     dy int
// 	priority int    // The priority of the item in the queue.
// 	// The index is needed by update and is maintained by the heap.Interface methods.
// 	index int // The index of the item in the heap.
// }

// A PriorityQueue implements heap.Interface and holds Items.
// type PriorityQueue []*Item
//
// func (pq PriorityQueue) Len() int { return len(pq) }
//
// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].priority < pq[j].priority
// }
//
// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }
//
// func (pq *PriorityQueue) Push(x any) {
// 	n := len(*pq)
// 	item := x.(*Item)
// 	item.index = n
// 	*pq = append(*pq, item)
// }
//
// func (pq *PriorityQueue) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
// 	item.index = -1 // for safety
// 	*pq = old[0 : n-1]
// 	return item
// }
//
// var directions = []utils.Pair{
// 	{X: 1, Y: 0},
// 	{X: 0, Y: 1},
// 	{X: -1, Y: 0},
// 	{X: 0, Y: -1},
// }
//
// const (
// 	RIGHT = 0
// 	DOWN  = 1
// 	LEFT  = 2
// 	UP    = 3
// )
//
// type raceCoords struct {
//     x int
//     y int
//     dx int
//     dy int
// }
//
// func containsCoords(rc raceCoords) func(raceCoords) bool {
//     return func (cmpRC raceCoords) bool {
//         return rc.x == cmpRC.x && rc.dx == cmpRC.dx && rc.dy == cmpRC.dy && rc.y == cmpRC.y
//     }
// }
//
// func copyMatrix(s [][]byte) [][]byte {
// 	newMatrix := make([][]byte, len(s))
//
// 	for i, line := range s {
// 		newMatrix[i] = make([]byte, len(line))
// 		copy(newMatrix[i], line)
// 	}
//
// 	return newMatrix
// }
//
// func findScore(race [][]byte, pos utils.Pair) int {
//     pq := make(PriorityQueue, 0)
//     heap.Init(&pq)
//     start := &Item{x: pos.X, y: pos.Y, dx: 0, dy: 1, priority: 0}
//     heap.Push(&pq, start)
//     seen := make([]raceCoords, 0)
//     seen = append(seen, raceCoords{x: pos.X, y: pos.Y, dx: 0, dy: 1})
//
//     for pq.Len() > 0 {
//         pqNode := heap.Pop(&pq).(*Item)
//         x, y, dx, dy, cost := pqNode.x, pqNode.y, pqNode.dx, pqNode.dy, pqNode.priority
//         seen = append(seen, raceCoords{x: x, y: y, dx: dx, dy: dy})
//         if race[pqNode.y][pqNode.x] == 'E' {
//             fmt.Println(cost)
//             return cost
//         }
//
//         newNodes := []Item{
//             {x: x + dx, y: y + dy, dx: dx, dy: dy, priority: 1 + cost},
//             {x: x, y: y, dx: -dy, dy: dx, priority: 1000 + cost},
//             {x: x, y: y, dx: dy, dy: -dx, priority: 1000 + cost},
//         }
//
//         for _, newNode := range newNodes {
//             if race[newNode.y][newNode.x] == '#' { continue }
//             newCoords := raceCoords{newNode.x, newNode.y, newNode.dx, newNode.dy}
//             if slices.ContainsFunc(seen, containsCoords(newCoords)) { continue }
//
//             heap.Push(&pq, &newNode)
//         }
//     }
//
//     return -1
// }

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

func pathfind(dataMap [][]byte, endPointX, endPointY int, initPos utils.Pair) int {
	pq := make(PriorityQueue, 1, 200)
	pq[0] = item{cost: 0, node: initPos}
	seen := make([]utils.Pair, 1)
	seen[0] = utils.NewPair(0, 0)

	for len(pq) != 0 {
        fmt.Println(pq)
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
			if neighbor.X == endPointX && neighbor.Y == endPointY {
				return cost
			}
            pq.push(neighbor, cost)
		}
	}

	return -1
}

func findInitPos(race [][]byte, target byte) utils.Pair {
	for y, line := range race {
		for x, char := range line {
			if char == target {
				return utils.NewPair(x, y)
			}
		}
	}
	return utils.NewPair(0, 0)
}

func main() {
	data := utils.GetDataFromFile()
	splitData := strings.Split(data, "\n")
	fmt.Println(data)
	race := make([][]byte, len(splitData))

	for i, line := range splitData {
		race[i] = []byte(line)
	}

	initPos := findInitPos(race, 'S')
    finalPos := findInitPos(race, 'E')
	fmt.Println(initPos)
	fmt.Println("Part 1:", pathfind(race, finalPos.X, finalPos.Y, initPos))

	// raceStr := make([]string, len(race))
	// for i, line := range race {
	// 	raceStr[i] = string(line)
	// }
}
