package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	digits := []rune{}

	for _, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		}
	}

	calibrationRunes := []rune{digits[0], digits[len(digits)-1]}
	number, err := strconv.Atoi(string(calibrationRunes))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(number)

	return number
}
