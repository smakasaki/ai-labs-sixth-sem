package main

import (
	"ai-lab-2/pkg/algorithms"
	"ai-lab-2/pkg/graph"
	"fmt"
)

func main() {
	filename := "map.csv"

	graphData, cities, err := graph.LoadGraphFromCSV(filename)
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}

	startCity := "Arad"
	targetCity := "Bucharest"
	
	heuristics := map[string]int{"Arad": 366, "Bucharest": 0, "Craiova": 160, "Drobita": 242, "Eforie": 161, "Fagaras": 176, "Giurgiu": 77, "Hirsova": 151, "Iasi": 226, "Lugoj": 244, "Mehedia": 241, "Neamt": 234, "Oradea": 380, "Pitesti": 100, "RM": 193, "Sibiu": 253, "Timisoara": 329, "Urziceni": 80, "Vaslui": 199, "Zerind": 374}

	path := algorithms.AStar(graphData, cities, startCity, targetCity, heuristics)
	fmt.Printf("\nPath from %s to %s (A*): %v\n", startCity, targetCity, path)
	distance, err := graph.CalculatePathDistance(path, cities, graphData)
    if err != nil {
        fmt.Printf("Error calculating path distance: %v\n", err)
    } else {
        fmt.Printf("Total distance for the path %v is %d km\n", path, distance)
    }

	path = algorithms.BreadthFirstSearch(graphData, cities, startCity, targetCity)
	fmt.Printf("\nPath from %s to %s (BFS): %v\n", startCity, targetCity, path)
	distance, err = graph.CalculatePathDistance(path, cities, graphData)
    if err != nil {
        fmt.Printf("Error calculating path distance: %v\n", err)
    } else {
        fmt.Printf("Total distance for the path %v is %d km\n", path, distance)
    }

	path = algorithms.BidirectionalSearch(graphData, cities, startCity, targetCity)
	fmt.Printf("\nPath from %s to %s (Bidirectional): %v\n", startCity, targetCity, path)
	distance, err = graph.CalculatePathDistance(path, cities, graphData)
    if err != nil {
        fmt.Printf("Error calculating path distance: %v\n", err)
    } else {
        fmt.Printf("Total distance for the path %v is %d km\n", path, distance)
    }

	path, err = algorithms.DepthFirstSearch(graphData, cities, startCity, targetCity)
	if err != nil {
		fmt.Printf("Error in DFS: %v\n", err)
	} else {
		fmt.Printf("\nPath from %s to %s (DFS): %v\n", startCity, targetCity, path)
		distance, err = graph.CalculatePathDistance(path, cities, graphData)
		if err != nil {
			fmt.Printf("Error calculating path distance: %v\n", err)
		} else {
			fmt.Printf("Total distance for the path %v is %d km\n", path, distance)
		}
	}

	path, err = algorithms.UniformCostSearch(graphData, cities, startCity, targetCity)
	if err != nil {
		fmt.Printf("Error in UCS: %v\n", err)
	} else {
		fmt.Printf("\nPath from %s to %s (UCS): %v\n", startCity, targetCity, path)
		distance, err = graph.CalculatePathDistance(path, cities, graphData)
		if err != nil {
			fmt.Printf("Error calculating path distance: %v\n", err)
		} else {
			fmt.Printf("Total distance for the path %v is %d km\n", path, distance)
		}
	}

	path, err = algorithms.GreedyBestFirstSearch(graphData, cities, startCity, targetCity, heuristics)
	if err != nil {
		fmt.Printf("Error in Greedy Best-First: %v\n", err)
	} else {
		fmt.Printf("\nPath from %s to %s (Greedy Best-First): %v\n", startCity, targetCity, path)
		distance, err = graph.CalculatePathDistance(path, cities, graphData)
		if err != nil {
			fmt.Printf("Error calculating path distance: %v\n", err)
		} else {
			fmt.Printf("Total distance for the path %v is %d km\n", path, distance)
		}
	}

}
