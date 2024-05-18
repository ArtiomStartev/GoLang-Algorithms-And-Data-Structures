package main

import (
	"math"
)

type Edge struct {
	Source      string
	Destination string
	Weight      int
}

func findPrimMST(graph map[string]map[string]int) []Edge {
	var mst []Edge

	vertices := make([]string, 0, len(graph))
	for vertex, _ := range graph {
		vertices = append(vertices, vertex)
	}

	visited := make(map[string]bool)
	visited[vertices[0]] = true

	for len(visited) < len(vertices) {
		minEdge := Edge{"", "", math.MaxInt64}

		for vertex := range visited {
			for neighbor, weight := range graph[vertex] {
				if !visited[neighbor] && weight < minEdge.Weight {
					minEdge = Edge{vertex, neighbor, weight}
				}
			}
		}

		mst = append(mst, minEdge)
		visited[minEdge.Destination] = true
	}

	return mst
}

func calculatePrimMSTCost(mst []Edge) int {
	var totalCost int
	for _, edge := range mst {
		totalCost += edge.Weight
	}
	return totalCost
}
