package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile("[0-9]+")

func main() {
	file, err := os.Open("./day6/input.txt")
	// file, err := os.Open("./day6/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	times := getNumbers(scanner)
	distances := getNumbers(scanner)
	scanner.Scan()
	result := 1

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, time := range times {
		result = result * getWaysToWin(time, distances[i])

	}

	file.Close()
	log.Println(result)
}

func getWaysToWin(time int, distance int) int {
	ways := 0

	for spentTime := 1; spentTime < time; spentTime++ {
		remainingTime := time - spentTime

		if spentTime*remainingTime > distance {
			ways++
		}
	}

	return ways
}

func getNumbers(scanner *bufio.Scanner) []int {
	scanner.Scan()
	text := scanner.Text()
	numberString := numberRegex.FindAllString(text, -1)
	numberList := []int{}

	for _, number := range numberString {
		intNum, err := strconv.Atoi(number)

		if err != nil {
			panic(err)
		}

		numberList = append(numberList, intNum)
	}

	return numberList
}
