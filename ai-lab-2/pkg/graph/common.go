package graph

import "fmt"

// City represents each node in the graph.
type City struct {
	Name      string
	Heuristic int // Estimated cost to the target for A*.
}

// Reconstructs a path from a parent map.
func ReconstructPath(parent map[int]int, cities []string, start, target int) []string {
	path := []string{}
	for at := target; at != start; at = parent[at] {
		path = append([]string{cities[at]}, path...)
	}
	path = append([]string{cities[start]}, path...)
	return path
}

func CalculatePathDistance(path []string, cities []string, graphData [][]int) (int, error) {
    totalDistance := 0

    if len(path) < 2 {
        return 0, fmt.Errorf("path must contain at least two cities to calculate distance")
    }

    // Translate city names to indices
    cityIndices := make(map[string]int)
    for index, city := range cities {
        cityIndices[city] = index
    }

    // Sum distances between consecutive cities in the path
    for i := 0; i < len(path)-1; i++ {
        startCityIndex, ok1 := cityIndices[path[i]]
        endCityIndex, ok2 := cityIndices[path[i+1]]
        if !ok1 || !ok2 {
            return 0, fmt.Errorf("one or more cities in the path are not found in the adjacency matrix")
        }
        distance := graphData[startCityIndex][endCityIndex]
        if distance == 0 && startCityIndex != endCityIndex {
            return 0, fmt.Errorf("no direct connection between %s and %s in the graph", path[i], path[i+1])
        }
        totalDistance += distance
    }

    return totalDistance, nil
}
