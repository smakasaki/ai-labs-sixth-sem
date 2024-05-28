package algorithms

import (
	"ai-lab-2/pkg/graph"
	"fmt"
)

// BreadthFirstSearch finds the shortest path in an unweighted graph using the BFS algorithm.
func BreadthFirstSearch(adjMatrix [][]int, cities []string, start, target string) []string {
	startIndex, targetIndex := -1, -1
	// Find the indices of the start and target cities.
	for i, city := range cities {
		if city == start {
			startIndex = i
		}
		if city == target {
			targetIndex = i
		}
	}

	if startIndex == -1 || targetIndex == -1 {
		fmt.Println("Invalid start or target city.")
		return nil
	}

	// Queue for BFS
	queue := []int{startIndex}
	// To keep track of visited cities
	visited := make([]bool, len(cities))
	visited[startIndex] = true
	// To reconstruct the path
	parent := make(map[int]int)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == targetIndex {
			return graph.ReconstructPath(parent, cities, startIndex, targetIndex)
		}

		for neighbor, distance := range adjMatrix[current] {
			if distance > 0 && !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}
	return nil
}
