package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		id, possible := parseLine(text)

		if possible {
			total += id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	log.Println(total)
}

func parseLine(line string) (int, bool) {
	prefix, results, _ := strings.Cut(line, ": ")
	isGameValid := validateGame(strings.Split(results, "; "))

	if !isGameValid {
		return 0, false
	}

	idString := strings.Split(prefix, " ")[1]
	gameId, err := strconv.Atoi(idString)

	if err != nil {
		log.Fatal(err)
	}

	return gameId, true
}

func validateGame(results []string) bool {
	for _, result := range results {
		cubeResults := strings.Split(result, ", ")

		for _, cubeRes := range cubeResults {
			valid := validateCubes(cubeRes)

			if !valid {
				return false
			}
		}
	}

	return true
}

var cubeLimitMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func validateCubes(cubeRes string) bool {
	amountString, colour, _ := strings.Cut(cubeRes, " ")

	amount, err := strconv.Atoi(amountString)

	if err != nil {
		log.Fatal(err)
	}

	limit := cubeLimitMap[colour]

	return amount <= limit
}
