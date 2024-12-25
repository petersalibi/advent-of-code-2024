package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

type item struct {
	cost      int
	node      utils.Pair
	direction utils.Pair
}

func (it item) HashKey() string {
    return it.node.HashKey() + "," + it.direction.HashKey()
}

func (it item) unwrap() (int, int, int, int, int) {
	return it.cost, it.node.X, it.node.Y, it.direction.X, it.direction.Y
}

type PriorityQueue []item

func (pq *PriorityQueue) push(node item) {
	*pq = append(*pq, node)
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

func (pq *PriorityQueue) contains(node utils.Pair, dir utils.Pair) bool {
	for _, pqItem := range *pq {
		if utils.Equal(node, pqItem.node) && utils.Equal(pqItem.direction, dir) {
			return true
		}
	}
	return false
}

type nodeDirPair struct {
	node utils.Pair
	dir  utils.Pair
}

func containNodeDirPair(nodeDir nodeDirPair) func(nodeDirPair) bool {
	return func(compND nodeDirPair) bool {
		return utils.Equal(nodeDir.node, compND.node) && utils.Equal(nodeDir.dir, compND.dir)
	}
}

func findMinCost(nodes []item) int {
    return slices.MinFunc(nodes, func (a, b item) int {
        return a.cost - b.cost
    }).cost
}

func pathfind(dataMap [][]byte, initPos, finalPos utils.Pair) int {
	pq := make(PriorityQueue, 1, 200)
	initDir := utils.NewPair(1, 0)
	pq[0] = item{cost: 0, node: initPos, direction: initDir}
	seen := make([]nodeDirPair, 1)
	seen[0] = nodeDirPair{initPos, initDir}
    lowestCost := make(map[string]int)
    minCost := math.MaxInt

    cameFrom := make(map[string][]item)

	for len(pq) != 0 {
        curItem := pq.pop()
		cost, x, y, dx, dy := curItem.unwrap()


        if lCost, ok := lowestCost[curItem.HashKey()]; ok {
            if lCost < cost { continue }
        }

        lowestCost[curItem.HashKey()] = cost

		seen = append(seen, nodeDirPair{utils.NewPair(x, y), utils.NewPair(dx, dy)})
        if cost > minCost {
            break
        }
        if dataMap[y][x] == 'E' {
            // return cost
            if cost > minCost { break }
            minCost = cost
        }

        neighbors := make([]item, 0)
		if dataMap[y + dy][x + dx] != '#' {
            neighbors = append(neighbors, item{cost + 1, utils.NewPair(x+dx, y+dy), utils.NewPair(dx, dy)})
        }
		if dataMap[y + dx][x - dy] != '#' {
            neighbors = append(neighbors, item{cost + 1000, utils.NewPair(x, y), utils.NewPair(-dy, dx)})
        }
		if dataMap[y - dx][x + dy] != '#' {
            neighbors = append(neighbors, item{cost + 1000, utils.NewPair(x, y), utils.NewPair(dy, -dx)})
        }

		for _, neighbor := range neighbors {
			nCost, nx, ny, ndx, ndy := neighbor.unwrap()

			// not wall
			if dataMap[ny][nx] == '#' {
				continue
			}

            pairString := strconv.Itoa(nx) + "," + strconv.Itoa(ny) + "," + strconv.Itoa(ndx) + "," + strconv.Itoa(ndy)
            cameFrom[pairString] = append(cameFrom[pairString], item{cost: nCost, node: utils.NewPair(x, y), direction: utils.NewPair(dx, dy)})
			// not in seen or pq
			if pq.contains(neighbor.node, neighbor.direction) ||
            slices.ContainsFunc(seen, containNodeDirPair(nodeDirPair{node: neighbor.node, dir: neighbor.direction})) {
				continue
			}

			pq.push(neighbor)
		}
	}

    seenNodes := make([]nodeDirPair, 0)
    nodeQueue := make([]item, 0)
    nodeQueue = append(nodeQueue, cameFrom[finalPos.HashKey() + ",0,-1"]...)

    for len(nodeQueue) != 0 {
        var curNode item
        curNode, nodeQueue = nodeQueue[len(nodeQueue) - 1], nodeQueue[:len(nodeQueue) - 1]
        curPair := nodeDirPair{node: curNode.node, dir: curNode.direction}
        if slices.ContainsFunc(seenNodes, containNodeDirPair(curPair)) {
            continue
        }
        seenNodes = append(seenNodes, curPair)
        prevNodes := cameFrom[curNode.HashKey()]
        minCost := findMinCost(prevNodes)

        for _, node := range prevNodes {
            if node.cost == minCost {
                nodeQueue = append(nodeQueue, node)
            }
        }
    }

    for _, node := range seenNodes {
        dataMap[node.node.Y][node.node.X] = 'O'
    }

    // for _, line := range dataMap {
    //     fmt.Println(string(line))
    // }

    numUnique := 0
    for _, line := range dataMap {
        for _, letter := range line {
            if letter == 'O' {
                numUnique += 1
            }
        }
    }

    fmt.Println("Part 1:", minCost)
	return numUnique + 1
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
	fmt.Println("Part 2:", pathfind(race, initPos, finalPos))
}
