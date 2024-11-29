package main

import (
	"fmt"
	"os"

	"entual.de/adventofcode/utils"
)

//////////// Daily Code here /////////////

func strToMap(str string) map[string]int {
	charMap := make(map[string]int)

	for _, rune := range str {
		char := string(rune)
		charMap[char] += 1
	}

	return charMap
}

func part1(input []string) {
	fmt.Printf("Executing Part 1: \n\n")

	result := 0
	counter := 0
	alphabetMap := make(map[string]int) // Map mit rune (Unicode-Codepunkt) als Schlüssel und int als Wert

	for i := 1; i <= 26; i++ {
		char := string(rune('a' + i - 1))
		alphabetMap[char] = i
	}

	for i := 1; i <= 26; i++ {
		char := string(rune('A' + i - 1))
		alphabetMap[char] = i + 26
	}

	for _, line := range input {
		itemsTotal := len(line)
		itemsPerRucksack := itemsTotal / 2

		rucksack1 := line[0:itemsPerRucksack]
		rucksack2 := line[itemsPerRucksack:itemsTotal]

		r1map := strToMap(rucksack1)
		r2map := strToMap(rucksack2)

		for key, _ := range r1map {
			if _, found := r2map[key]; found {
				result += alphabetMap[key]
				counter += 1
			}
		}

	}
	fmt.Printf("Result: %d\n", result)
}

func part2(input []string) {
	fmt.Printf("Executing Part 2: \n\n")

	result := 0
	alphabetMap := make(map[string]int) // Map mit rune (Unicode-Codepunkt) als Schlüssel und int als Wert

	for i := 1; i <= 26; i++ {
		char := string(rune('a' + i - 1))
		alphabetMap[char] = i
	}

	for i := 1; i <= 26; i++ {
		char := string(rune('A' + i - 1))
		alphabetMap[char] = i + 26
	}

	total := len(input) / 3

	for i := 0; i < total; i++ {

		g1 := input[i*3]
		g2 := input[i*3+1]
		g3 := input[i*3+2]

		g1map := strToMap(g1)
		g2map := strToMap(g2)
		g3map := strToMap(g3)

		for key, _ := range g1map {
			if _, found := g2map[key]; found {
				if _, found := g3map[key]; found {
					result += alphabetMap[key]
				}
			}
		}

	}
	fmt.Printf("Result: %d\n", result)
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
