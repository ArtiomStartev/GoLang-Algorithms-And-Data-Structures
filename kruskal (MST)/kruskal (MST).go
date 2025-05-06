package kruskal

import (
	"sort"
)

type Graph map[string]map[string]int

type Edge struct {
	Source      string
	Destination string
	Weight      int
}

type DisjointSet map[string]string

// Kruskal finds the Minimum Spanning Tree (MST) of a graph using Kruskal's algorithm
func Kruskal(graph Graph) []Edge {
	edges := getSortedEdges(graph)

	parent := initDisjointSet(graph)

	return findMST(edges, parent)
}

// getSortedEdges retrieves and sorts the edges of the graph
func getSortedEdges(graph Graph) []Edge {
	var edges []Edge

	for src, neighbors := range graph {
		for dest, weight := range neighbors {
			edges = append(edges, Edge{src, dest, weight})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	return edges
}

// initDisjointSet initializes the disjoint set for Kruskal's algorithm
func initDisjointSet(graph Graph) DisjointSet {
	parent := make(DisjointSet)
	for node := range graph {
		parent[node] = node
	}
	return parent
}

// findMST finds the Minimum Spanning Tree (MST) using Kruskal's algorithm
func findMST(edges []Edge, parent DisjointSet) []Edge {
	var mstEdges []Edge

	for _, edge := range edges {
		srcParent := findParent(parent, edge.Source)
		destParent := findParent(parent, edge.Destination)

		if srcParent != destParent {
			mstEdges = append(mstEdges, edge)
			union(parent, srcParent, destParent)
		}
	}

	return mstEdges
}

// findParent finds the root of the disjoint set containing the node
func findParent(parent DisjointSet, node string) string {
	if parent[node] == node {
		return node
	}
	return findParent(parent, parent[node])
}

// union merges two sets in the disjoint set
func union(parent DisjointSet, src, dest string) {
	srcRoot := findParent(parent, src)
	destRoot := findParent(parent, dest)
	parent[srcRoot] = destRoot
}

// calculateMSTCost calculates the total cost of the MST
func calculateMSTCost(mst []Edge) int {
	var totalCost int
	for _, edge := range mst {
		totalCost += edge.Weight
	}
	return totalCost
}
