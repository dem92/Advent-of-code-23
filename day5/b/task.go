package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day5/input.txt")
	// file, err := os.Open("./day5/example.txt")

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

// Default value is higher than any location number in the input
var closestLocation = 1000000000000000000
var seeds = [][]int{}
var maps = [][][]int{}

func getClosestLocation(scanner *bufio.Scanner) {
	parseSeeds(scanner)

	for scanner.Scan() {
		parseMap(scanner, [][]int{})
	}

	for _, seedRange := range seeds {
		checkBranch(seedRange, 0)
	}
}

// Check for the range in each row of the map with the given index
func checkBranch(numberRange []int, mapIndex int) {
	if mapIndex == len(maps) {
		// We've reached past the last map,
		// which means that the numbers in the range are location numbers.
		// Compare the lowest location number to the closest location
		// we've found so far.
		lowestNumber := numberRange[0]

		if lowestNumber < closestLocation {
			closestLocation = lowestNumber
		}

		return
	}

	checkRow(numberRange, mapIndex, 0)
}

// Check one row in a map at a time
func checkRow(numberRange []int, mapI, rowI int) {
	mapRows := maps[mapI]

	if rowI == len(mapRows) {
		// We've reached past the last row, with no match for the range.
		// Pass the range to the next map.
		checkBranch(numberRange, mapI+1)
		return
	}

	row := mapRows[rowI]

	rangeBaseNumber := numberRange[0]
	rangeCeilingNumber := rangeBaseNumber + numberRange[1]

	rowSourceStart := row[1]
	rowSourceEnd := rowSourceStart + row[2]
	rowDestinationStart := row[0]

	// <> = range area, || = source area, . = 0 or more numbers
	// Case covered |.<.>.| and |.<.|.>
	if rangeBaseNumber >= rowSourceStart && rangeBaseNumber < rowSourceEnd {
		diffRangeRow := rangeBaseNumber - rowSourceStart
		destinationStart := rowDestinationStart + diffRangeRow
		var destinationEnd int

		if rangeCeilingNumber > rowSourceEnd {
			newRangeLength := rangeCeilingNumber - rowSourceEnd
			// Pass the sub range that doesn't match on to the next row
			checkRow([]int{rowSourceEnd, newRangeLength}, mapI, rowI+1)
			destinationEnd = rowSourceEnd - rangeBaseNumber
		} else {
			destinationEnd = numberRange[1]
		}

		// Pass the sub range that matches on to the next map
		checkBranch([]int{destinationStart, destinationEnd}, mapI+1)
		// Case covered <.|.>.|
	} else if rangeCeilingNumber > rowSourceStart && rangeCeilingNumber <= rowSourceEnd {
		matchingRangeLength := rangeCeilingNumber - rowSourceStart
		// Pass the sub range that matches on to the next map
		checkBranch([]int{rowDestinationStart, matchingRangeLength}, mapI+1)

		newRangeLength := numberRange[1] - matchingRangeLength
		// Pass the sub range that doesn't match on to the next row
		checkRow([]int{rangeBaseNumber, newRangeLength}, mapI, rowI+1)
		// Case covered <.|.|.>
	} else if rangeCeilingNumber >= rowSourceEnd && rangeBaseNumber < rowSourceStart {
		// Pass the sub range that matches on to the next map
		checkBranch([]int{rowDestinationStart, row[2]}, mapI+1)
		// Pass the first sub range that doesn't match on to the next row
		checkRow([]int{rowSourceEnd, rangeCeilingNumber - rowSourceEnd}, mapI, rowI+1)
		// Pass the second sub range that doesn't match on to the next row
		checkRow([]int{rangeBaseNumber, rowSourceStart - rangeBaseNumber}, mapI, rowI+1)
		// Nothing matches
	} else {
		// Pass the sub range that doesn't match on to the next row
		checkRow(numberRange, mapI, rowI+1)
	}
}

// Parse a map from the input, and store it in the maps array
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

	seedNumberString, _ := strings.CutPrefix(seedText, "seeds: ")
	seedNumbers := parseNumbers(seedNumberString)

	for i := 0; i < len(seedNumbers); i += 2 {
		seeds = append(seeds, []int{seedNumbers[i], seedNumbers[i+1]})
	}

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
