package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var numberWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var numWordRegex *regexp.Regexp
var reverseNumWordRegex *regexp.Regexp
var numberRegex = regexp.MustCompile("[0-9]")
var numberMap map[string]int

func main() {
	setup()
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

func setup() {
	numWordRegexString := getNumberRegex()
	reverseNumWordRegexString := reverseString(numWordRegexString)
	numWordRegex = regexp.MustCompile(numWordRegexString)
	reverseNumWordRegex = regexp.MustCompile(reverseNumWordRegexString)
	numberMap = getNumberMap()
}

func findCalibrationNumber(line string) int {
	firstNumber := getFirstNumber(line, numWordRegex)
	lastNumber := getFirstNumber(reverseString(line), reverseNumWordRegex)
	number, err := strconv.Atoi(firstNumber + lastNumber)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func getFirstNumber(text string, rx *regexp.Regexp) string {
	regexMatch := rx.FindString(text)
	regexMatchIndex := rx.FindStringIndex(text)

	if regexMatch != "" {
		text = replaceAtIndex(text, convertWordToIntRune(regexMatch), regexMatchIndex[0])
	}

	return numberRegex.FindString(text)
}

func replaceAtIndex(text string, replacement rune, index int) string {
	runes := []rune(text)
	runes[index] = replacement
	return string(runes)
}

func convertWordToIntRune(word string) rune {
	intString := fmt.Sprintf("%d", numberMap[word])
	runes := []rune(intString)
	return runes[0]
}

func reverseString(s string) string {
	runes := []rune(s)

	for index, reverseIndex := 0, len(runes)-1; index < reverseIndex; index, reverseIndex = index+1, reverseIndex-1 {
		runes[index], runes[reverseIndex] = runes[reverseIndex], runes[index]
	}
	return string(runes)
}

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

func getNumberMap() map[string]int {
	numMap := map[string]int{}

	for i, numberWord := range numberWords {
		numMap[numberWord] = i
		numMap[reverseString(numberWord)] = i
	}

	return numMap
}
