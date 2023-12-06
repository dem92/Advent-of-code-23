package main

import (
	"bufio"
	// "// fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// var numberRegex = regexp.MustCompile("[0-9]+")

func main() {
	file, err := os.Open("./day5/input.txt")
	// file, err := os.Open("./day5/example.txt")
	// file, err := os.Open("../example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	getClosestLocation(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	log.Println(closestLocation)
}

var closestLocation = 1000000000000000000

func getClosestLocation(scanner *bufio.Scanner) {
	parseSeeds(scanner)

	for scanner.Scan() {
		parseMap(scanner, [][]int{})
	}

	for _, seedRange := range seeds {
		checkBranch(seedRange, 0)
	}
}

func checkBranch(numberRange []int, mapIndex int) {
	// fmt.Println()
	// fmt.Printf("Branch #: %d\n", numberRange)
	// fmt.Printf("MapIndex #: %d\n", mapIndex)
	// baseNumber := numberRange[0]
	// ceilingNumber := baseNumber + numberRange[1]
	// fmt.Printf("Base num #: %d, Ceiling num #: %d\n", baseNumber, ceilingNumber)

	if mapIndex == len(maps) {
		lowestNumber := numberRange[0]
		// fmt.Printf("End of the line, lowest number is: %d\n", lowestNumber)

		if lowestNumber < closestLocation {
			closestLocation = lowestNumber
		}

		return
	}

	checkRow(numberRange, mapIndex, 0)
}

// Check one row at a time
func checkRow(numberRange []int, mapI, rowI int) {
	mapRows := maps[mapI]

	if rowI == len(mapRows) {
		// fmt.Printf("End of the line, checking branch: %d\n", numberRange[0])
		checkBranch(numberRange, mapI+1)
		return
	}

	row := mapRows[rowI]
	rangeBaseNumber := numberRange[0]
	rangeCeilingNumber := rangeBaseNumber + numberRange[1]

	rowSourceStart := row[1]
	rowSourceEnd := rowSourceStart + row[2]
	rowDestinationStart := row[0]
	// fmt.Printf("Row: %d\n", rowI)

	if rangeBaseNumber >= rowSourceStart && rangeBaseNumber < rowSourceEnd {
		// fmt.Printf("1 m%d\n", mapI)
		// fmt.Printf("diff: %d\n", rangeBaseNumber-rowSourceStart)
		diffRangeRow := rangeBaseNumber - rowSourceStart
		destinationStart := rowDestinationStart + diffRangeRow
		var destinationEnd int

		// if rangeCeilingNumber >= rowSourceEnd {
		if rangeCeilingNumber > rowSourceEnd {
			newRangeLength := rangeCeilingNumber - rowSourceEnd
			checkRow([]int{rowSourceEnd, newRangeLength}, mapI, rowI+1)
			destinationEnd = rowSourceEnd - rangeBaseNumber
		} else {
			destinationEnd = numberRange[1]
		}

		checkBranch([]int{destinationStart, destinationEnd}, mapI+1)
	} else if rangeCeilingNumber > rowSourceStart && rangeCeilingNumber <= rowSourceEnd {
		// fmt.Printf("2 m%d\n", mapI)
		// fmt.Printf("diff: %d\n", rangeCeilingNumber-rowSourceStart)
		matchingRangeLength := rangeCeilingNumber - rowSourceStart
		checkBranch([]int{rowDestinationStart, matchingRangeLength}, mapI+1)

		newRangeLength := numberRange[1] - matchingRangeLength
		checkRow([]int{rangeBaseNumber, newRangeLength}, mapI, rowI+1)
	} else if rangeCeilingNumber >= rowSourceEnd && rangeBaseNumber < rowSourceStart {
		// fmt.Printf("3 m%d\n", mapI)
		// fmt.Printf("diff down: %d\n", rowSourceStart-rangeBaseNumber)
		// fmt.Printf("diff up: %d\n", rangeCeilingNumber-rowSourceEnd)
		checkBranch([]int{rowDestinationStart, row[2]}, mapI+1)
		checkRow([]int{rowSourceEnd, rangeCeilingNumber - rowSourceEnd}, mapI, rowI+1)
		checkRow([]int{rangeBaseNumber, rowSourceStart - rangeBaseNumber}, mapI, rowI+1)
	} else {
		// fmt.Printf("4 m%d\n", mapI)
		checkRow(numberRange, mapI, rowI+1)
	}
}

var seeds = [][]int{}
var maps = [][][]int{}

func parseMap(scanner *bufio.Scanner, currentMap [][]int) {
	// First line is title in iteration 1, we don't care about it
	scanner.Scan()
	content := scanner.Text()

	if content == "" {
		maps = append(maps, currentMap)
		// // fmt.Printf("New map: %#v\n", currentMap)
		return
	}

	numbers := parseNumbers(content)
	currentMap = append(currentMap, numbers)

	parseMap(scanner, currentMap)
}

func parseSeeds(scanner *bufio.Scanner) {
	scanner.Scan()
	seedText := scanner.Text()

	seedNumberString, _ := strings.CutPrefix(seedText, "seeds: ")
	seedNumbers := parseNumbers(seedNumberString)

	for i := 0; i < len(seedNumbers); i += 2 {
		seeds = append(seeds, []int{seedNumbers[i], seedNumbers[i+1]})
	}

	// // fmt.Printf("Seeds: %#v\n", seeds)

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
