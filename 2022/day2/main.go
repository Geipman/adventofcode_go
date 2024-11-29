package main

import (
	"fmt"
	"os"
	"strings"

	"entual.de/adventofcode/utils"
)

//////////// Daily Code here /////////////

type t_round struct {
	opponent string
	me       string
}

func part1(input []string) {
	// Code

	rounds := []t_round{}

	for _, line := range input {

		splitted := strings.Split(line, " ")

		opponent := "x"
		me := "x"

		switch splitted[0] {
		case "A":
			opponent = "R"
		case "B":
			opponent = "P"
		case "C":
			opponent = "S"
		}

		switch splitted[1] {
		case "Y":
			me = "P"
		case "X":
			me = "R"
		case "Z":
			me = "S"
		}

		round := t_round{opponent: opponent, me: me}
		rounds = append(rounds, round)

	}

	results := []int{}

	for _, round := range rounds {

		type_score := 0
		result_score := 0
		total_score := 0

		switch round.me {
		case "R":
			type_score = 1
		case "P":
			type_score = 2
		case "S":
			type_score = 3
		}

		if round.me == round.opponent {
			result_score = 3
		}

		if round.me == "R" && round.opponent == "S" || round.me == "S" && round.opponent == "P" || round.me == "P" && round.opponent == "R" {
			result_score = 6
		}

		if round.me == round.opponent {
			result_score = 3
		}

		total_score = type_score + result_score

		results = append(results, total_score)
	}

	fmt.Println(rounds)
	fmt.Println(results)

	sum := 0
	for _, result := range results {
		sum += result
	}

	fmt.Println(sum)
}

func part2(input []string) {
	rounds := []t_round{}

	for _, line := range input {

		splitted := strings.Split(line, " ")

		opponent := "x"
		me := "x"

		switch splitted[0] {
		case "A":
			opponent = "R"
		case "B":
			opponent = "P"
		case "C":
			opponent = "S"
		}

		if opponent == "R" && splitted[1] == "Y" || opponent == "P" && splitted[1] == "X" || opponent == "S" && splitted[1] == "Z" {
			me = "R"
		}

		if opponent == "P" && splitted[1] == "Y" || opponent == "S" && splitted[1] == "X" || opponent == "R" && splitted[1] == "Z" {
			me = "P"
		}

		if opponent == "S" && splitted[1] == "Y" || opponent == "R" && splitted[1] == "X" || opponent == "P" && splitted[1] == "Z" {
			me = "S"
		}

		round := t_round{opponent: opponent, me: me}
		rounds = append(rounds, round)

	}

	results := []int{}

	for _, round := range rounds {

		type_score := 0
		result_score := 0
		total_score := 0

		switch round.me {
		case "R":
			type_score = 1
		case "P":
			type_score = 2
		case "S":
			type_score = 3
		}

		if round.me == round.opponent {
			result_score = 3
		}

		if round.me == "R" && round.opponent == "S" || round.me == "S" && round.opponent == "P" || round.me == "P" && round.opponent == "R" {
			result_score = 6
		}

		if round.me == round.opponent {
			result_score = 3
		}

		total_score = type_score + result_score

		results = append(results, total_score)
	}

	fmt.Println(rounds)
	fmt.Println(results)

	sum := 0
	for _, result := range results {
		sum += result
	}

	fmt.Println(sum)
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
