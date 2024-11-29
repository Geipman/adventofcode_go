package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"entual.de/adventofcode/utils"
)

//////////// Daily Code here ////////////

type t_move struct {
	count int
	from  int
	to    int
}

func part1(input []string) {
	fmt.Printf("Executing Part 1: \n\n")

	// find number of stacks
	var stackConfig []string
	var moveConfig []string
	var moves []t_move
	numOfStacks := (len(input[0]) / 4) + 1

	var stacks []*utils.Stack[string]

	// init stacks
	for i := 0; i < numOfStacks; i++ {
		stacks = append(stacks, new(utils.Stack[string]))
	}

	// get configs
	splitter := 0
	for _, line := range input {

		if strings.TrimSpace(line) == "" {
			splitter = 1
			continue
		}

		if splitter == 0 {
			stackConfig = append(stackConfig, line)
		} else {
			moveConfig = append(moveConfig, line)
		}
	}

	for _, line := range stackConfig {
		fmt.Println((line))
	}

	// create stacks
	for i := len(stackConfig) - 2; i >= 0; i-- {
		line := stackConfig[i]

		for j := 0; j < numOfStacks; j++ {
			c := string(line[4*j+1])
			if c != " " {
				stacks[j].Push(c)
			}
		}
	}

	// parse moves
	for _, line := range moveConfig {

		//move 1 from 2 to 1
		re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

		matches := re.FindAllStringSubmatch(line, -1)

		count, _ := strconv.Atoi(matches[0][1])
		from, _ := strconv.Atoi(matches[0][2])
		to, _ := strconv.Atoi(matches[0][3])

		moves = append(moves, t_move{count: count, from: from, to: to})
	}

	for _, m := range moves {
		for i := 1; i <= m.count; i++ {
			stacks[m.to-1].Push(stacks[m.from-1].Pop())
		}
	}

	fmt.Print("\nResult: ")
	for _, stack := range stacks {
		fmt.Print(stack.Pop())
	}
	fmt.Print("\n")
}

func part2(input []string) {
	fmt.Printf("Executing Part 2: \n\n")

	// find number of stacks
	var stackConfig []string
	var moveConfig []string
	var moves []t_move
	numOfStacks := (len(input[0]) / 4) + 1

	var stacks []*utils.Stack[string]

	// init stacks
	for i := 0; i < numOfStacks; i++ {
		stacks = append(stacks, new(utils.Stack[string]))
	}

	// get configs
	splitter := 0
	for _, line := range input {

		if strings.TrimSpace(line) == "" {
			splitter = 1
			continue
		}

		if splitter == 0 {
			stackConfig = append(stackConfig, line)
		} else {
			moveConfig = append(moveConfig, line)
		}
	}

	for _, line := range stackConfig {
		fmt.Println((line))
	}

	// create stacks
	for i := len(stackConfig) - 2; i >= 0; i-- {
		line := stackConfig[i]

		for j := 0; j < numOfStacks; j++ {
			c := string(line[4*j+1])
			if c != " " {
				stacks[j].Push(c)
			}
		}
	}

	// parse moves
	for _, line := range moveConfig {

		//move 1 from 2 to 1
		re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

		matches := re.FindAllStringSubmatch(line, -1)

		count, _ := strconv.Atoi(matches[0][1])
		from, _ := strconv.Atoi(matches[0][2])
		to, _ := strconv.Atoi(matches[0][3])

		moves = append(moves, t_move{count: count, from: from, to: to})
	}

	for _, m := range moves {

		var tempStack utils.Stack[string]

		for i := 1; i <= m.count; i++ {
			tempStack.Push(stacks[m.from-1].Pop())
		}

		for i := 1; i <= m.count; i++ {
			stacks[m.to-1].Push(tempStack.Pop())
		}
	}

	fmt.Print("\nResult: ")
	for _, stack := range stacks {
		fmt.Print(stack.Pop())
	}
	fmt.Print("\n")
}

////////// Main Function ///////////

func main() {

	// Abfrage Inputparameter
	fileName, part, err := utils.Initializer(os.Args)
	if err != nil {
		fmt.Printf("Fehler: %s\n", err)
		return
	}

	// Datei einlesen
	input, err := utils.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Fehler: %s\n", err)
		return
	}

	// Ausführung des Rätselcodes
	if part == "part1" {
		part1(input)
	}

	if part == "part2" {
		part2(input)
	}
}
