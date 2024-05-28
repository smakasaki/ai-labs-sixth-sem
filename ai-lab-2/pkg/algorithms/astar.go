package algorithms

import (
	"container/heap"
	"fmt"
)

func AStar(graphData [][]int, cities []string, start, target string, heuristics map[string]int) []string {
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

    // Priority queue to select the next city to visit.
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)

    // Add the start city to the priority queue.
    startItem := &PriorityQueueItem{
        city: startIndex,
        cost: heuristics[start],
        parent: -1,
    }
    heap.Push(&pq, startItem)

    // To keep track of the shortest paths from the start.
    cameFrom := make(map[int]int)
    costSoFar := make(map[int]int)
    costSoFar[startIndex] = 0

    for pq.Len() > 0 {
        currentItem := heap.Pop(&pq).(*PriorityQueueItem)
        current := currentItem.city

        // If we reach the target city, reconstruct the path.
        if current == targetIndex {
            return reconstructPathAstar(cameFrom, cities, startIndex, targetIndex)
        }

        for neighbor, cost := range graphData[current] {
            if cost > 0 { // There's a road between the cities.
                newCost := costSoFar[current] + cost
                if _, found := costSoFar[neighbor]; !found || newCost < costSoFar[neighbor] {
                    costSoFar[neighbor] = newCost
                    cameFrom[neighbor] = current
                    heap.Push(&pq, &PriorityQueueItem{
                        city: neighbor,
                        cost: newCost,
                        parent: current,
                    })
                }
            }
        }
    }

    return nil
}

func reconstructPathAstar(cameFrom map[int]int, cities []string, start, target int) []string {
    path := []string{}
    for at := target; at != start; at = cameFrom[at] {
        path = append([]string{cities[at]}, path...)
    }
    path = append([]string{cities[start]}, path...)
    return path
}
