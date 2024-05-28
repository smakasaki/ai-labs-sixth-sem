package algorithms

import (
	"container/heap"
	"fmt"
)

// Define a struct for the priority queue.
type GreedyPriorityQueueItem struct {
	city    int
	heuristic int
}

// Implement a priority queue.
type GreedyPriorityQueue []*GreedyPriorityQueueItem

func (pq GreedyPriorityQueue) Len() int { return len(pq) }

func (pq GreedyPriorityQueue) Less(i, j int) bool {
	// Min-Heap based on heuristic value
	return pq[i].heuristic < pq[j].heuristic
}

func (pq GreedyPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *GreedyPriorityQueue) Push(x interface{}) {
	item := x.(*GreedyPriorityQueueItem)
	*pq = append(*pq, item)
}

func (pq *GreedyPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// GreedyBestFirstSearch finds a path using the greedy best-first algorithm.
func GreedyBestFirstSearch(graphData [][]int, cities []string, start, target string, heuristic map[string]int) ([]string, error) {
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

	visited := make([]bool, len(cities))
	parent := make([]int, len(cities))
	for i := range parent {
		parent[i] = -1
	}

	pq := make(GreedyPriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &GreedyPriorityQueueItem{city: startIndex, heuristic: heuristic[cities[startIndex]]})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*GreedyPriorityQueueItem)
		current := item.city

		if current == targetIndex {
			return reconstructPath(parent, cities, startIndex, targetIndex) // Remove the ", nil" from the return statement
		}

		if visited[current] {
			continue
		}
		visited[current] = true

		for neighbor, distance := range graphData[current] {
			if distance > 0 && !visited[neighbor] {
				heap.Push(&pq, &GreedyPriorityQueueItem{
					city: neighbor,
					heuristic: heuristic[cities[neighbor]],
				})
				parent[neighbor] = current
			}
		}
	}

	return nil, fmt.Errorf("no path found from %s to %s", start, target)
}
