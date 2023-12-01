package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		total += findCalibrationNumber(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(total)

	file.Close()
}

func findCalibrationNumber(line string) int {
	wordMatches := findWords(line)
	numberMatches := findNumbers(line)

	if wordMatches == nil {
		wordMatches = numberMatches
	}

	firstMatch := wordMatches[0]

	if numberMatches[0].index < firstMatch.index {
		firstMatch = numberMatches[0]
	}

	lastMatch := wordMatches[len(wordMatches)-1]

	if numberMatches[len(numberMatches)-1].index > lastMatch.index {
		lastMatch = numberMatches[len(numberMatches)-1]
	}

	number, err := strconv.Atoi(firstMatch.value + lastMatch.value)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(line)
	fmt.Println(number)

	return number
}

type digitMatch struct {
	value string
	index int
}

func findWords(line string) []digitMatch {
	regexMatches := numberRegex.FindAllString(line, -1)

	if regexMatches == nil {
		return nil
	}

	regexMatchIndexes := numberRegex.FindAllStringIndex(line, -1)
	lastMatchIndex := len(regexMatches) - 1

	return []digitMatch{
		{index: regexMatchIndexes[0][0], value: convertNumberWordToInt(regexMatches[0])},
		{index: regexMatchIndexes[lastMatchIndex][1] - 1, value: convertNumberWordToInt(regexMatches[lastMatchIndex])},
	}
}

func findNumbers(line string) []digitMatch {
	matches := []digitMatch{}

	for i, c := range line {
		if unicode.IsDigit(c) {
			matches = append(matches, digitMatch{value: string(c), index: i})
		}
	}

	return matches
}

var numberWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var numberRegex = getNumberRegex()

func getNumberRegex() *regexp.Regexp {
	regex := ""

	for i, word := range numberWords {
		if i != 0 {
			regex += "|"
		}

		regex += word
	}

	return regexp.MustCompile(regex)
}

func convertNumberWordToInt(word string) string {
	for i, numberWord := range numberWords {
		if word == numberWord {
			return fmt.Sprintf("%d", i)
		}
	}

	return ""
}
