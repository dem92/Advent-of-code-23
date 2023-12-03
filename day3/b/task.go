package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type scannedNumber struct {
	indexRange []int
	number     int
}

func main() {
	// file, err := os.Open("./day3/example.txt")
	file, err := os.Open("./day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	total := scanText(file)
	file.Close()
	log.Println(total)
}

func scanText(file *os.File) int {
	symbolMap := map[int][]int{0: {}}
	numberMap := map[int][]scannedNumber{0: {}}
	scanner := bufio.NewScanner(file)
	index := 1

	for scanner.Scan() {
		text := scanner.Text()
		symbolMap[index] = getGearIndexes(text)
		numberMap[index] = getNumberIndexes(text)

		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	symbolMap[index] = []int{}
	numberMap[index] = []scannedNumber{}
	total := getPartNumberTotal(symbolMap, numberMap)

	return total
}

func getPartNumberTotal(symbolMap map[int][]int, numberMap map[int][]scannedNumber) int {
	total := 0

	// Iterate over every line from input
	for lineIndex := 1; lineIndex < len(numberMap)-1; lineIndex++ {
		// Combine numbers on current line, and lines directly above and below
		lineNumbers := []scannedNumber{}
		lineNumbers = append(lineNumbers, numberMap[lineIndex-1]...)
		lineNumbers = append(lineNumbers, numberMap[lineIndex]...)
		lineNumbers = append(lineNumbers, numberMap[lineIndex+1]...)

		gearIndexes := symbolMap[lineIndex]

		for _, i := range gearIndexes {
			isProperGear := false
			partNumbers := []int{}

			for _, num := range lineNumbers {
				indexRange := num.indexRange
				numLeftEdge := indexRange[0] - 1
				numRightEdge := indexRange[1]

				// Check if any symbol index overlaps with a number's surrounding area
				if numLeftEdge <= i && numRightEdge >= i {
					partNumbers = append(partNumbers, num.number)
				}
			}

			if len(partNumbers) == 2 {
				isProperGear = true
			}

			if isProperGear {
				total += partNumbers[0] * partNumbers[1]
			}
		}
	}

	return total
}

func getGearIndexes(text string) []int {
	gearRegex := regexp.MustCompile("[*]")
	gearIndexes := []int{}
	regexMatchIndex := gearRegex.FindAllStringIndex(text, -1)

	for _, indexRange := range regexMatchIndex {
		if len(indexRange) > 0 {
			gearIndexes = append(gearIndexes, indexRange[0])
		}
	}

	return gearIndexes
}

func getNumberIndexes(text string) []scannedNumber {
	numberRegex := regexp.MustCompile("[0-9]+")
	numberIndexes := []scannedNumber{}
	regexMatch := numberRegex.FindAllString(text, -1)
	regexMatchIndex := numberRegex.FindAllStringIndex(text, -1)

	for i, match := range regexMatch {
		number, err := strconv.Atoi(match)

		if err != nil {
			panic("Failed to convert word to int")
		}

		numberIndexes = append(numberIndexes, scannedNumber{
			number:     number,
			indexRange: regexMatchIndex[i],
		})
	}

	return numberIndexes
}
