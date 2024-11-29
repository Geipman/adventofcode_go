package main

import (
	"fmt"
	"os"

	"entual.de/adventofcode/utils"
)

//////////// Daily Code here /////////////

func part1(input []string) {
	fmt.Printf("Executing Part 1: \n\n")

	datastream := input[0]
	charMap := make(map[string]int)
	index := 0
	marker := 4

	// mjqjpqmgbljsphdztnvjfqwrcgsmlb

	for i := 0; i < len(datastream); i++ {
		index = i

		char := string(datastream[i])
		charMap[char] += 1

		if i > marker-1 {
			old_char := string(datastream[i-marker])
			charMap[old_char] -= 1

			if charMap[old_char] == 0 {
				delete(charMap, old_char)
			}
		}

		if len(charMap) == marker {
			break
		}

	}

	fmt.Printf("Result: %d\n", index+1)

}

func part2(input []string) {
	fmt.Printf("Executing Part 2: \n\n")

	datastream := input[0]
	charMap := make(map[string]int)
	index := 0
	marker := 14

	// mjqjpqmgbljsphdztnvjfqwrcgsmlb

	for i := 0; i < len(datastream); i++ {
		index = i

		char := string(datastream[i])
		charMap[char] += 1

		if i > marker-1 {
			old_char := string(datastream[i-marker])
			charMap[old_char] -= 1

			if charMap[old_char] == 0 {
				delete(charMap, old_char)
			}
		}

		if len(charMap) == marker {
			break
		}

	}

	fmt.Printf("Result: %d\n", index+1)
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
