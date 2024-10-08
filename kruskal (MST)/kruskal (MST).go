package kruskal

import (
	"sort"
)

type GraphEdge struct {
	Source      string
	Destination string
	Weight      int
}

/**
 * Finds the Minimum Spanning Tree (MST) of a weighted, connected, undirected graph using Kruskal's Algorithm.
 * @param graph map[string]map[string]int - The graph represented as an adjacency list.
 * @returns []GraphEdge - A slice representing the edges of the Minimum Spanning Tree.
 */
func findMinimumSpanningTree(graph map[string]map[string]int) []GraphEdge {
	// Step 1: Sort all edges
	edges := getSortedEdges(graph)

	// Step 2: Initialize disjoint set
	parent := initializeDisjointSet(graph)

	// Step 3: Find MST edges
	return findMST(edges, parent)
}

/**
 * Sorts all edges of the graph in increasing order of their weight.
 * @param graph map[string]map[string]int - The graph represented as an adjacency list.
 * @returns []GraphEdge - A slice representing the sorted edges.
 */
func getSortedEdges(graph map[string]map[string]int) []GraphEdge {
	var edges []GraphEdge

	for src, neighbors := range graph {
		for dest, weight := range neighbors {
			edges = append(edges, GraphEdge{src, dest, weight})
		}
	}

	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	return edges
}

/**
 * Initializes the disjoint set data structure.
 * @param graph map[string]map[string]int - The graph represented as an adjacency list.
 * @returns map[string]string - A map representing the disjoint set.
 */
func initializeDisjointSet(graph map[string]map[string]int) map[string]string {
	parent := make(map[string]string)

	for vertex, _ := range graph {
		// Each vertex is its own parent initially
		parent[vertex] = vertex
	}

	return parent
}

/**
 * Finds the Minimum Spanning Tree (MST) edges using Kruskal's Algorithm.
 * @param edges []GraphEdge - A slice representing all the edges of the graph sorted in increasing order of their weight.
 * @param parent map[string]string - The disjoint set data structure.
 * @returns []GraphEdge - A slice representing the edges of the Minimum Spanning Tree.
 */
func findMST(edges []GraphEdge, parent map[string]string) []GraphEdge {
	var mstEdges []GraphEdge

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

/**
 * Finds the parent of a vertex in the disjoint set.
 * @param parent map[string]string - The disjoint set data structure.
 * @param vertex string - The vertex whose parent needs to be found.
 * @returns string - The parent vertex.
 */
func findParent(parent map[string]string, vertex string) string {
	// Base case: vertex is its own parent
	if parent[vertex] == vertex {
		return vertex
	}
	// Recursively find parent
	return findParent(parent, parent[vertex])
}

/**
 * Performs the union operation in the disjoint set.
 * @param parent map[string]string - The disjoint set data structure.
 * @param x string - The first vertex.
 * @param y string - The second vertex.
 */
func union(parent map[string]string, x, y string) {
	xRoot := findParent(parent, x)
	yRoot := findParent(parent, y)
	parent[xRoot] = yRoot // Set parent of xRoot to yRoot
}

/**
 * Calculates the total cost of the Minimum Spanning Tree (MST).
 * @param mst []GraphEdge - The Minimum Spanning Tree represented as a slice of edges.
 * @returns int - The total cost of the Minimum Spanning Tree.
 */
func calculateMSTCost(mst []GraphEdge) int {
	var totalCost int
	for _, edge := range mst {
		totalCost += edge.Weight
	}
	return totalCost
}
