package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]+")

func main() {
	file, err := os.Open("./day4/input.txt")
	// file, err := os.Open("./day4/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		total += getCardValue(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	log.Println(total)
}

func getCardValue(line string) int {
	_, numbers, _ := strings.Cut(line, ":")
	winningNums, myNums, _ := strings.Cut(numbers, "|")
	winningNumList := numberRegex.FindAllString(winningNums, -1)
	myNumList := numberRegex.FindAllString(myNums, -1)
	cardScore := 0

	for _, winNum := range winningNumList {
		for _, myNum := range myNumList {
			if winNum == myNum {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore = cardScore * 2
				}

				break
			}
		}
	}

	return cardScore
}
