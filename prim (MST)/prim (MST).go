package prim__MST_

import (
	"math"
)

type Node struct {
	Source      string
	Destination string
	Weight      int
}

func findPrimMST(graph map[string]map[string]int) []Node {
	var mst []Node

	vertices := make([]string, 0, len(graph))
	for vertex, _ := range graph {
		vertices = append(vertices, vertex)
	}

	visited := make(map[string]bool)
	visited[vertices[0]] = true

	for len(visited) < len(vertices) {
		minEdge := Node{"", "", math.MaxInt64}

		for vertex := range visited {
			for neighbor, weight := range graph[vertex] {
				if !visited[neighbor] && weight < minEdge.Weight {
					minEdge = Node{vertex, neighbor, weight}
				}
			}
		}

		mst = append(mst, minEdge)
		visited[minEdge.Destination] = true
	}

	return mst
}

func calculatePrimMSTCost(mst []Node) int {
	var totalCost int
	for _, edge := range mst {
		totalCost += edge.Weight
	}
	return totalCost
}
