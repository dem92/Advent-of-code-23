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
		total += getGamePower(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	log.Println(total)
}

func getGamePower(line string) int {
	_, allResults, _ := strings.Cut(line, ": ")
	results := strings.Split(allResults, "; ")
	var minAmountMap = map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}

	for _, result := range results {
		cubeResults := strings.Split(result, ", ")

		for _, cubeRes := range cubeResults {
			colour, amount := getCubeAmount(cubeRes)

			if minAmountMap[colour] < amount {
				minAmountMap[colour] = amount
			}
		}
	}

	power := 1

	for _, value := range minAmountMap {
		power = power * value
	}

	return power
}

func getCubeAmount(cubeRes string) (string, int) {
	amountString, colour, _ := strings.Cut(cubeRes, " ")

	amount, err := strconv.Atoi(amountString)

	if err != nil {
		log.Fatal(err)
	}

	return colour, amount
}
