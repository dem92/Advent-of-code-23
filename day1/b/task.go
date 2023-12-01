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

type matches struct {
	firstMatch numberMatch
	lastMatch  numberMatch
}

type numberMatch struct {
	value string
	index int
}

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

	file.Close()
	log.Println(total)
}

func findCalibrationNumber(line string) int {
	wordMatches := findNumberWords(line)
	numberMatches := findNumbers(line)

	if wordMatches == nil && numberMatches == nil {
		return 0
	}

	var firstMatch numberMatch
	var lastMatch numberMatch

	if wordMatches == nil {
		firstMatch = numberMatches.firstMatch
		lastMatch = numberMatches.lastMatch
	} else if numberMatches == nil {
		firstMatch = wordMatches.firstMatch
		lastMatch = wordMatches.lastMatch
	} else {
		firstMatch = wordMatches.firstMatch
		lastMatch = wordMatches.lastMatch

		if numberMatches.firstMatch.index < firstMatch.index {
			firstMatch = numberMatches.firstMatch
		}
		if numberMatches.lastMatch.index > lastMatch.index {
			lastMatch = numberMatches.lastMatch
		}
	}

	number, err := strconv.Atoi(firstMatch.value + lastMatch.value)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func findNumberWords(line string) *matches {
	regexString := getNumberRegex()
	regex := regexp.MustCompile(regexString)
	regexMatch := regex.FindString(line)
	regexMatchIndex := regex.FindStringIndex(line)

	if regexMatch == "" {
		return nil
	}

	firstMatch := numberMatch{
		index: regexMatchIndex[0],
		value: convertNumberWordToInt(regexMatch),
	}

	regex = regexp.MustCompile(reverseString(regexString))
	reversedLine := reverseString(line)
	regexMatch = regex.FindString(reversedLine)
	regexMatchIndex = regex.FindStringIndex(reversedLine)

	lastMatch := numberMatch{
		index: len(line) - regexMatchIndex[1],
		value: convertNumberWordToInt(reverseString(regexMatch)),
	}

	return &matches{
		firstMatch: firstMatch,
		lastMatch:  lastMatch,
	}
}

func findNumbers(line string) *matches {
	numberMatches := []numberMatch{}

	for i, c := range line {
		if unicode.IsDigit(c) {
			numberMatches = append(numberMatches, numberMatch{value: string(c), index: i})
		}
	}

	if len(numberMatches) == 0 {
		return nil
	}

	return &matches{
		firstMatch: numberMatches[0],
		lastMatch:  numberMatches[len(numberMatches)-1],
	}
}

var numberWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getNumberRegex() string {
	regex := ""

	for i, word := range numberWords {
		if i != 0 {
			regex += "|"
		}

		regex += word
	}

	return regex
}

func convertNumberWordToInt(word string) string {
	for i, numberWord := range numberWords {
		if word == numberWord {
			return fmt.Sprintf("%d", i)
		}
	}

	panic("failed to convert number to word")
}

func reverseString(s string) string {
	runes := []rune(s)

	for index, reverseIndex := 0, len(runes)-1; index < reverseIndex; index, reverseIndex = index+1, reverseIndex-1 {
		runes[index], runes[reverseIndex] = runes[reverseIndex], runes[index]
	}
	return string(runes)
}
