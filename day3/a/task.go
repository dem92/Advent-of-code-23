package main

import (
	"bufio"
	"fmt"
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
		symbolMap[index] = getSymbolIndexes(text)
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
		lineNumbers := numberMap[lineIndex]

		// Combine indexes with symbols on current line, and lines directly above and below
		symbolIndexes := []int{}
		symbolIndexes = append(symbolIndexes, symbolMap[lineIndex-1]...)
		symbolIndexes = append(symbolIndexes, symbolMap[lineIndex]...)
		symbolIndexes = append(symbolIndexes, symbolMap[lineIndex+1]...)

		for _, num := range lineNumbers {
			isPartNumber := false
			indexRange := num.indexRange
			leftEdge := indexRange[0] - 1
			rightEdge := indexRange[1]

			for _, symIndex := range symbolIndexes {
				// Check if any symbol index overlaps with a number's surrounding area
				if leftEdge <= symIndex && rightEdge >= symIndex {
					isPartNumber = true
					break
				}
			}

			if isPartNumber {
				total += num.number
				fmt.Println(num.number)
			}
		}
	}

	return total
}

func getSymbolIndexes(text string) []int {
	symbolRegex := regexp.MustCompile("[^0-9.]")
	symbolIndexes := []int{}
	regexMatchIndex := symbolRegex.FindAllStringIndex(text, -1)

	for _, indexRange := range regexMatchIndex {
		if len(indexRange) > 0 {
			symbolIndexes = append(symbolIndexes, indexRange[0])
		}
	}

	return symbolIndexes
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
