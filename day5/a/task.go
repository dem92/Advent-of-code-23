package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var seeds = []int{}
var maps = [][][]int{}

func main() {
	file, err := os.Open("./day5/input.txt")
	// file, err := os.Open("./day5/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	res := getClosestLocation(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	log.Println(res)
}

func getClosestLocation(scanner *bufio.Scanner) int {
	parseSeeds(scanner)

	for scanner.Scan() {
		parseMap(scanner, [][]int{})
	}

	closestLocation := 1000000000000000000

	for _, seedNumber := range seeds {
		currentNumber := seedNumber

		for _, m := range maps {
			for _, row := range m {
				sourceStart := row[1]
				sourceEnd := sourceStart + row[2]

				if currentNumber >= sourceStart && currentNumber < sourceEnd {
					currentNumber = row[0] + currentNumber - sourceStart
					break
				}
			}
		}

		if currentNumber < closestLocation {
			closestLocation = currentNumber
		}
	}

	return closestLocation
}

func parseMap(scanner *bufio.Scanner, currentMap [][]int) {
	scanner.Scan()
	content := scanner.Text()

	if content == "" {
		maps = append(maps, currentMap)
		return
	}

	numbers := parseNumbers(content)
	currentMap = append(currentMap, numbers)

	parseMap(scanner, currentMap)
}

func parseSeeds(scanner *bufio.Scanner) {
	scanner.Scan()
	seedText := scanner.Text()

	seedNumbers, _ := strings.CutPrefix(seedText, "seeds: ")
	seeds = parseNumbers(seedNumbers)

	scanner.Scan()
}

func parseNumbers(text string) []int {
	numberStrings := strings.Split(text, " ")
	numbers := []int{}

	for _, seed := range numberStrings {
		seedNum, err := strconv.Atoi(seed)

		if err != nil {
			panic("Failed to convert seed to number")
		}

		numbers = append(numbers, seedNum)
	}

	return numbers
}
