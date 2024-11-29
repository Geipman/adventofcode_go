package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Initializer(args []string) (string, string, error) {
	// Prüfen, ob ein Argument übergeben wurde
	if len(os.Args) != 3 {
		fmt.Println("Bitte einen Modus angeben: 'train' oder 'run' + 'part1' oder 'part2'")
		return "", "", errors.New("x")
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
		return "", "", errors.New("x")
	}

	return fileName, part, nil
}

func ReadFile(fileName string) ([]string, error) {

	input := []string{}

	// Datei öffnen
	file, err := os.Open(fileName)
	if err != nil {
		return input, fmt.Errorf("error opening file: %s", fileName)
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
