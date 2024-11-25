package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(fileName string) ([]string, error) {

	input := []string{}

	// Datei öffnen
	file, err := os.Open(fileName)
	if err != nil {
		return input, fmt.Errorf("Fehler beim Öffnen der Datei: %s", fileName)
	}
	defer file.Close()

	// Scanner für das Einlesen zeilenweise
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Gelesene Zeile holen
		line := strings.TrimSpace(scanner.Text())
		input = append(input, line)
	}

	return input, nil
}

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

func main() {

	fmt.Println(os.Args)
	// Prüfen, ob ein Argument übergeben wurde
	if len(os.Args) != 3 {
		fmt.Println("Bitte einen Modus angeben: train oder run + part1 oder part2")
		return
	}

	// Parameter lesen
	mode := os.Args[1]
	part := os.Args[2]
	var fileName string

	// Datei basierend auf dem Modus auswählen
	if mode == "train" {
		fileName = "data_small.dat"
	} else if mode == "run" {
		fileName = "data_big.dat"
	} else {
		fmt.Println("Ungültiger Modus. Verwenden Sie 'train' oder 'run'.")
		return
	}

	// Datei einlesen
	input, err := readFile(fileName)
	if err != nil {
		fmt.Printf("Fehler: %s\n", err)
	}

	if part == "part1" {
		part1(input)
	}

	if part == "part2" {
		part2(input)
	}
}
