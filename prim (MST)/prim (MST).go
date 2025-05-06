package prim__MST_

import (
	"math"
)

type Graph map[string]map[string]int

type Edge struct {
	Source      string
	Destination string
	Weight      int
}

// Prim finds the Minimum Spanning Tree (MST) of a graph using Prim's algorithm
func Prim(graph Graph, startNode string) []Edge {
	var mst []Edge
	visited := make(map[string]bool)
	visited[startNode] = true

	for len(visited) < len(graph) {
		minEdge := Edge{Weight: math.MaxInt64}

		for src := range visited {
			for dest, weight := range graph[src] {
				if !visited[dest] && weight < minEdge.Weight {
					minEdge = Edge{src, dest, weight}
				}
			}
		}

		mst = append(mst, minEdge)
		visited[minEdge.Destination] = true
	}

	return mst
}

// calculatePrimMSTCost calculates the total cost of the MST
func calculatePrimMSTCost(mst []Edge) int {
	var totalCost int
	for _, edge := range mst {
		totalCost += edge.Weight
	}
	return totalCost
}
