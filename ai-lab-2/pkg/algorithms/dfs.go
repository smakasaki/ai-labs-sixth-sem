package algorithms

import (
	"fmt"
)

func DepthFirstSearch(graphData [][]int, cities []string, start, target string) ([]string, error) {
	startIndex, targetIndex := -1, -1
	for i, city := range cities {
		if city == start {
			startIndex = i
		}
		if city == target {
			targetIndex = i
		}
	}

	if startIndex == -1 || targetIndex == -1 {
		return nil, fmt.Errorf("invalid start or target city")
	}

	// Visited array to keep track of visited nodes
	visited := make([]bool, len(cities))
	var path []string

	// Helper function to perform DFS
	var dfs func(current int) bool
	dfs = func(current int) bool {
		if current == targetIndex {
			path = append(path, cities[current])
			return true
		}
		visited[current] = true
		path = append(path, cities[current])

		for neighbor, distance := range graphData[current] {
			if distance > 0 && !visited[neighbor] {
				if dfs(neighbor) {
					return true
				}
			}
		}

		// Backtrack
		path = path[:len(path)-1]
		return false
	}

	if !dfs(startIndex) {
		return nil, fmt.Errorf("no path found from %s to %s", start, target)
	}
	return path, nil
}
