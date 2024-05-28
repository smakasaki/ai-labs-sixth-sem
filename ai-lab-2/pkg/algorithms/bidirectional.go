package algorithms

import (
	"fmt"
)

func BidirectionalSearch(graphData [][]int, cities []string, start, target string) []string {
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
		fmt.Println("Invalid start or target city.")
		return nil
	}

	// Initialize front and back queues, visited arrays and parent maps
	queueFront := []int{startIndex}
	queueBack := []int{targetIndex}
	visitedFront := make([]bool, len(cities))
	visitedBack := make([]bool, len(cities))
	parentFront := make(map[int]int)
	parentBack := make(map[int]int)

	visitedFront[startIndex] = true
	visitedBack[targetIndex] = true

	// Explore nodes from both ends until queues meet
	for len(queueFront) > 0 && len(queueBack) > 0 {
		// Explore from the front
		currentFront := queueFront[0]
		queueFront = queueFront[1:]

		for neighbor, distance := range graphData[currentFront] {
			if distance > 0 && !visitedFront[neighbor] {
				visitedFront[neighbor] = true
				parentFront[neighbor] = currentFront
				queueFront = append(queueFront, neighbor)

				// Check if visited by the backward search
				if visitedBack[neighbor] {
					return constructPath(parentFront, parentBack, neighbor, cities, startIndex, targetIndex)
				}
			}
		}

		// Explore from the back
		currentBack := queueBack[0]
		queueBack = queueBack[1:]

		for neighbor, distance := range graphData[currentBack] {
			if distance > 0 && !visitedBack[neighbor] {
				visitedBack[neighbor] = true
				parentBack[neighbor] = currentBack
				queueBack = append(queueBack, neighbor)

				// Check if visited by the front search
				if visitedFront[neighbor] {
					return constructPath(parentFront, parentBack, neighbor, cities, startIndex, targetIndex)
				}
			}
		}
	}
	return nil
}

func constructPath(parentFront, parentBack map[int]int, meetingPoint int, cities []string, start, target int) []string {
	// Construct the path from the start to the meeting point
	path := []string{cities[meetingPoint]}
	for i := meetingPoint; i != start; i = parentFront[i] {
		path = append([]string{cities[parentFront[i]]}, path...)
	}

	// Extend the path from the meeting point to the target
	for i := meetingPoint; i != target; i = parentBack[i] {
		path = append(path, cities[parentBack[i]])
	}

	return path
}
