package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]+")
var cardAmountMap = map[int]int{}

func main() {
	file, err := os.Open("./day4/input.txt")
	// file, err := os.Open("./day4/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	total := getAmountOfCards(file)

	file.Close()
	log.Println(total)
}

func getAmountOfCards(file *os.File) int {
	scanner := bufio.NewScanner(file)
	totalCards := 0
	cardNumber := 1

	for scanner.Scan() {
		text := scanner.Text()
		cardAmount := cardAmountMap[cardNumber]

		if cardAmount == 0 {
			cardAmount = 1
		}

		totalCards += cardAmount

		scratchCard(text, cardNumber, cardAmount)

		cardNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalCards
}

func scratchCard(text string, cardNumber int, amount int) {
	_, numbers, _ := strings.Cut(text, ":")
	winningNums, myNums, _ := strings.Cut(numbers, "|")
	winningNumList := numberRegex.FindAllString(winningNums, -1)
	myNumList := numberRegex.FindAllString(myNums, -1)
	cardsWon := 0

	for _, winNum := range winningNumList {
		for _, myNum := range myNumList {
			if winNum == myNum {
				cardsWon++
				addCard(cardNumber+cardsWon, amount)

				break
			}
		}
	}
}

func addCard(cardNumber int, amount int) {
	currentAmount := cardAmountMap[cardNumber]

	if currentAmount == 0 {
		currentAmount = 1
	}

	currentAmount += amount
	cardAmountMap[cardNumber] = currentAmount
}
