package algorithms

import (
	"container/heap"
	"fmt"
)


func UniformCostSearch(graphData [][]int, cities []string, start, target string) ([]string, error) {
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

	// Priority queue to manage which node to explore next
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PriorityQueueItem{city: startIndex, cost: 0, parent: -1})

	// To keep track of the least cost to reach each city
	costs := make([]int, len(cities))
	for i := range costs {
		costs[i] = int(^uint(0) >> 1) // Initialize with infinity
	}
	costs[startIndex] = 0

	// To reconstruct the path
	parent := make([]int, len(cities))
	for i := range parent {
		parent[i] = -1
	}

	visited := make([]bool, len(cities))

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityQueueItem)
		current := item.city

		if visited[current] {
			continue
		}

		visited[current] = true

		if current == targetIndex {
			path, err := reconstructPath(parent, cities, startIndex, targetIndex)
			if err != nil {
				return nil, err
			}
			return path, nil
		}

		for neighbor, distance := range graphData[current] {
			if distance > 0 && !visited[neighbor] {
				newCost := costs[current] + distance
				if newCost < costs[neighbor] {
					costs[neighbor] = newCost
					parent[neighbor] = current
					heap.Push(&pq, &PriorityQueueItem{city: neighbor, cost: newCost, parent: current})
				}
			}
		}
	}

	return nil, fmt.Errorf("no path found from %s to %s", start, target)
}

func reconstructPath(parents []int, cities []string, start, target int) ([]string, error) {
	if parents[target] == -1 {
		return nil, fmt.Errorf("no path to target found")
	}
	
	// Start from the target and work backwards to the start
	path := []string{}
	for current := target; current != -1; current = parents[current] {
		path = append(path, cities[current])
		if current == start {
			break
		}
	}

	// The path is currently from target to start, so we reverse it
	reversePath(path)
	return path, nil
}

// Helper function to reverse a slice of strings.
func reversePath(path []string) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}