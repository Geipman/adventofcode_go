package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"entual.de/adventofcode/utils"
)

func part1(input []string) {
	// Code
	biggest := 0
	local_sum := 0

	for _, line := range input {

		if line == "" {
			if local_sum > biggest {
				biggest = local_sum
			}
			local_sum = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Fehler beim Konvertieren der Zahl '%s': %v\n", line, err)
			continue
		}

		local_sum += num
		fmt.Printf("Line %s, Sum %d\n", line, local_sum)
	}

	if local_sum > biggest {
		biggest = local_sum
	}

	fmt.Printf("Sieger: %d\n", biggest)
}

func part2(input []string) {
	// Code
	groups_sum := []int{}
	local_sum := 0

	for _, line := range input {

		if line == "" {
			groups_sum = append(groups_sum, local_sum)
			local_sum = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Fehler beim Konvertieren der Zahl '%s': %v\n", line, err)
			continue
		}

		local_sum += num
		fmt.Printf("Line %s, Sum %d\n", line, local_sum)
	}

	groups_sum = append(groups_sum, local_sum)

	sort.Sort(sort.Reverse(sort.IntSlice(groups_sum)))

	fmt.Println(groups_sum[:3])
	fmt.Println(groups_sum[0] + groups_sum[1] + groups_sum[2])
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

	if part == "part1" {
		part1(input)
	}

	if part == "part2" {
		part2(input)
	}
}
