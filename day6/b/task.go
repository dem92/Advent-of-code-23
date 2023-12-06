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
	time := getNumber(scanner)
	distance := getNumber(scanner)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	result := getWaysToWin(time, distance)
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

func getNumber(scanner *bufio.Scanner) int {
	scanner.Scan()
	text := scanner.Text()
	numberString := numberRegex.FindAllString(text, -1)
	number := ""

	for _, n := range numberString {
		number += n
	}

	intNum, err := strconv.Atoi(number)

	if err != nil {
		panic(err)
	}

	return intNum
}
