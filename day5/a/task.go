package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]+")

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

	// Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
	for _, seedNumber := range seeds {
		currentNumber := seedNumber
		fmt.Printf("Seed #: %d\n", currentNumber)

		for i, m := range maps {
			fmt.Printf("Map: %d\n", i)
			for _, row := range m {
				fmt.Printf("Row: %#v\n", row)
				sourceStart := row[1]
				sourceEnd := sourceStart + row[2]

				if currentNumber >= sourceStart && currentNumber < sourceEnd {
					currentNumber = row[0] + currentNumber - sourceStart
					fmt.Printf("Update: %d\n", currentNumber)
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

type mapContent struct {
	rangeEnd    int
	destination int
}

var seeds = []int{}

// var maps = []map[int]mapContent{}

// var maps = []map[int]*int{}
var maps = [][][]int{}

func parseMap(scanner *bufio.Scanner, currentMap [][]int) {
	// First line is title in iteration 1, we don't care about it
	scanner.Scan()
	content := scanner.Text()

	if content == "" {
		maps = append(maps, currentMap)
		// fmt.Printf("New map: %#v\n", currentMap)
		return
	}

	numbers := parseNumbers(content)
	currentMap = append(currentMap, numbers)
	// currentMap[num]

	// for i := 0; i < numbers[2]; i++ {
	// 	destination := numbers[0] + i
	// 	currentMap[numbers[1]+i] = &destination
	// }

	parseMap(scanner, currentMap)
}

func parseSeeds(scanner *bufio.Scanner) {
	scanner.Scan()
	seedText := scanner.Text()

	seedNumbers, _ := strings.CutPrefix(seedText, "seeds: ")
	seeds = parseNumbers(seedNumbers)
	fmt.Printf("Seeds: %#v\n", seeds)

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
