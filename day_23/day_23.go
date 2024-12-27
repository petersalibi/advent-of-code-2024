package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strings"
)

func containsSlice(as []string) func([]string) bool {
    return func (bs []string) bool {
        for _, a := range as {
            for _, b := range bs {
                if a == b { return true }
            }
        }
        return false
    }
}

func findAllCliques(network map[string][]string, prevNodes []string, lookNode string) []string {
    cliques := make([]string, 0)

    neighbors :=  network[lookNode]
    if slices.Contains(neighbors, lookNode) { return cliques }
    fmt.Println(neighbors)

    for _, neighbor := range neighbors {
        if slices.Contains(prevNodes, neighbor) { continue }
        newClique := findAllCliques(network, append(prevNodes, neighbor), lookNode)
        if len(newClique) < 3 { continue }
        cliques = append(cliques, newClique...)
    }

    return cliques
}

// func allCliques(network map[string][]string) []string {
//     cliques := make([]string, 0)
//
//     for node := range network {
//         cliques = append(cliques, findAllCliques(network, []string{node}, node)...)
//     }
//
//     return cliques
// }

var cliques = make([]string, 0)
const MAXDEPTH = 7

func allCliques(network map[string][]string, node string, req []string, depth int) {
    slices.Sort(req)
    key := strings.Join(req, ",")
    if depth == 0 {
        return
    }
    if slices.Contains(cliques, key) {
        return
    }
    cliques = append(cliques, key)

    outer:
    for _, neighbor := range network[node] {
        if slices.Contains(req, neighbor) { continue }
        for _, query := range req {
            if !slices.Contains(network[query], neighbor) {
                continue outer
            }
        }

        allCliques(network, neighbor, append(req, neighbor), depth - 1)
    }
}

func main() {
    data := utils.GetDataFromFile()
    splitData := strings.Split(data, "\n")

    network := make(map[string][]string, 0)

    for _, conNodes := range splitData {
        nodePair := strings.Split(conNodes, "-")
        network[nodePair[0]] = append(network[nodePair[0]], nodePair[1])
        network[nodePair[1]] = append(network[nodePair[1]], nodePair[0])
    }

    // cliques := make([]string, 0)
    for x := range network {
        if x[0] != 't' { continue }
        fmt.Println("finding", x)
        allCliques(network, x, []string{x}, MAXDEPTH)
    }

    var maxClique []string
    for _, clique := range cliques {
        splitClique := strings.Split(clique, ",")
        if len(splitClique) > len(maxClique) {
            maxClique = splitClique
        }
    }

    // fmt.Println(network)
    fmt.Println(len(cliques))
    slices.Sort(maxClique)
    fmt.Println(strings.Join(maxClique, ","))
}
