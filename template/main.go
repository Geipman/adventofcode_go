package main

import (
	"fmt"
	"os"

	"entual.de/adventofcode/utils"
)

//////////// Daily Code here /////////////

func part1(input []string) {
	fmt.Printf("Executing Part 1: \n\n")
}

func part2(input []string) {
	fmt.Printf("Executing Part 2: \n\n")
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
