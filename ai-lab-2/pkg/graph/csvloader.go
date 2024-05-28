package graph

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// LoadGraphFromCSV loads a graph and city names from a CSV file.
func LoadGraphFromCSV(filename string) ([][]int, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	cities, err := reader.Read()
	if err != nil {
		return nil, nil, err
	}

	var graph [][]int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		var row []int
		for _, value := range record {
			if intVal, err := strconv.Atoi(value); err == nil {
				row = append(row, intVal)
			} else {
				return nil, nil, fmt.Errorf("conversion error: %v", err)
			}
		}
		graph = append(graph, row)
	}

	return graph, cities, nil
}
