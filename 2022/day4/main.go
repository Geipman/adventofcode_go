package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"entual.de/adventofcode/utils"
)

//////////// Daily Code here /////////////

func part1(input []string) {
	fmt.Printf("Executing Part 1: \n\n")

	result := 0

	for _, line := range input {
		pairs := strings.Split(line, ",")

		min1, _ := strconv.Atoi(strings.Split(pairs[0], "-")[0])
		max1, _ := strconv.Atoi(strings.Split(pairs[0], "-")[1])

		min2, _ := strconv.Atoi(strings.Split(pairs[1], "-")[0])
		max2, _ := strconv.Atoi(strings.Split(pairs[1], "-")[1])

		if (min1 <= min2 && max1 >= max2) || (min2 <= min1 && max2 >= max1) {
			result += 1
		}
	}

	println(result)
}

func part2(input []string) {
	fmt.Printf("Executing Part 2: \n\n")

	result := 0

	for _, line := range input {
		pairs := strings.Split(line, ",")

		min1, _ := strconv.Atoi(strings.Split(pairs[0], "-")[0])
		max1, _ := strconv.Atoi(strings.Split(pairs[0], "-")[1])

		min2, _ := strconv.Atoi(strings.Split(pairs[1], "-")[0])
		max2, _ := strconv.Atoi(strings.Split(pairs[1], "-")[1])

		// search for cases without overlap and negate, to get all with any overlap
		if !((max1 < min2) || (min1 > max2)) {
			result += 1
		}
	}

	println(result)
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
