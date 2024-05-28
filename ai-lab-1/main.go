package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type World [][]int
var log []string

func createWorld(length, width int) World {
	world := make(World, length)
	for i := range world {
		world[i] = make([]int, width)
	}
	return world
}

func addBlock(world World, x, y int) bool {
	if x >= len(world) || y >= len(world[0]) || x < 0 || y < 0 {
		log = append(log, fmt.Sprintf("Failed to place block at (%d, %d): Out of bounds.", x, y))
		return false
	}
	if world[x][y] != 0 {
		log = append(log, fmt.Sprintf("Failed to place block at (%d, %d): Position already occupied.", x, y))
		return false
	}
	if x != 0 && world[x-1][y] == 0 {
		log = append(log, fmt.Sprintf("Failed to place block at (%d, %d): No support underneath.", x, y))
		return false
	}
	world[x][y] = 1
	log = append(log, fmt.Sprintf("Place block at (%d, %d)", x, y))
	return true
}

func grasp(blockX, blockY int) {
    log = append(log, fmt.Sprintf("Grasp block at (%d, %d)", blockX, blockY))
}


func move(world World, block1X, block1Y, block2X, block2Y int, getRidOf bool) {
	if getRidOf {
		world[block1X][block1Y] = 0 // Удаляем блок с первой позиции
		log = append(log, fmt.Sprintf("Get rid of block at (%d, %d)", block1X, block1Y))
	}
	world[block2X][block2Y] = 1 // Помещаем блок на новую позицию
	log = append(log, fmt.Sprintf("Move block from (%d, %d) to (%d, %d)", block1X, block1Y, block2X, block2Y))
}

func putOn(world World, block1X, block1Y, block2X, block2Y int) {
    if world[block1X][block1Y] == 0 {
        fmt.Println("Cannot put on: No block to grasp.")
        return
    }

    if world[block2X][block2Y] != 0 {
        fmt.Println("Cannot put on: Position already occupied.")
        return
    }

    // Захватываем блок перед перемещением
    grasp(block1X, block1Y)

    // Временно удаляем блок из его текущей позиции для корректной проверки поддержки
    world[block1X][block1Y] = 0

    // Проверяем наличие поддержки перед перемещением
    if block2X > 0 && world[block2X-1][block2Y] == 0 {
        fmt.Println("Cannot put on: No support underneath.")
        // Возвращаем блок на исходное место, если перемещение не возможно
        world[block1X][block1Y] = 1
        return
    }

    // Выполняем перемещение блока
    move(world, block1X, block1Y, block2X, block2Y, true)
}

func printWorld(world World) {
	for i := len(world) - 1; i >= 0; i-- {
		for _, val := range world[i] {
			if val == 0 {
				fmt.Print(" . ")
			} else {
				fmt.Print(" # ")
			}
		}
		fmt.Println()
	}
}

func cmdHandler(world World) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Current log entries:")
		for _, entry := range log {
			fmt.Println(entry)
		}

		fmt.Print("Enter command: ")
		if !scanner.Scan() {
			continue
		}
		command := scanner.Text()
		args := strings.Split(command, " ")
		switch args[0] {
		case "Why", "How":
			handleQuestion(command)
		case "show":
			printWorld(world)
		case "put_on":
			if len(args) < 5 {
				fmt.Println("Usage: put_on [block1X] [block1Y] [block2X] [block2Y]")
				continue
			}
			block1X, err1 := strconv.Atoi(args[1])
			block1Y, err2 := strconv.Atoi(args[2])
			block2X, err3 := strconv.Atoi(args[3])
			block2Y, err4 := strconv.Atoi(args[4])
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				fmt.Println("Invalid coordinates. Please enter valid integers.")
				continue
			}
			putOn(world, block1X, block1Y, block2X, block2Y)
		case "kill", "quit":
			fmt.Println("Quitting program.")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

func handleQuestion(question string) {
    parts := strings.Fields(question)
    if len(parts) < 6 {
        fmt.Println("Invalid question format. Please use correct syntax like 'How did you grasp block at (0, 0)'.")
        return
    }

    // Формируем запрос к логу, приводя все к нижнему регистру для универсальности сравнения
    action := strings.ToLower(parts[3])  // действие: grasp, got rid of, moved
    details := strings.ToLower(strings.Join(parts[4:], " "))  // детализация действия

    query := fmt.Sprintf("%s %s", action, details)
    fmt.Printf("Formed query: '%s'\n", query)  // Дебаг: что формируем для поиска

    // Поиск в логе, также приводим лог к нижнему регистру перед сравнением
    found := false
    for _, entry := range log {
        fmt.Printf("Checking against log entry: '%s'\n", entry)  // Дебаг: сравнение с логом
        if strings.ToLower(entry) == query {  // Использование точного совпадения для проверки
            fmt.Println(entry)
            found = true
            break
        }
    }
    if !found {
        fmt.Println("No log entry found for this question.")
    }
}

func main() {
	var length, width int
	fmt.Print("Enter world length and width: ")
	fmt.Scan(&length, &width)
	world := createWorld(length, width)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter block coordinates (x y), or 'done' to finish: ")
		if !scanner.Scan() {
			continue
		}
		input := scanner.Text()
		if input == "done" {
			break
		}
		coords := strings.Split(input, " ")
		if len(coords) != 2 {
			fmt.Println("Invalid input, please enter two integers separated by space.")
			continue
		}
		x, errX := strconv.Atoi(coords[0])
		y, errY := strconv.Atoi(coords[1])
		if errX != nil || errY != nil {
			fmt.Println("Invalid coordinates, please enter valid integers.")
			continue
		}
		addBlock(world, x, y)
	}

	printWorld(world)

	cmdHandler(world)
}
