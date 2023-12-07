package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile("[0-9]+")

func main() {
	file, err := os.Open("./day7/input.txt")
	// file, err := os.Open("./day7/example.txt")
	// file, err := os.Open("../example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		parseHand(text)
	}
	fmt.Printf("%v\n", hands)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	sort.Slice(hands, func(i, j int) bool {
		bestHand := hands[i]
		currentHand := hands[j]

		if currentHand.strength > bestHand.strength {
			return true
		} else if currentHand.strength == bestHand.strength {
			return isCurrentHandBetter(currentHand, bestHand)
		}

		return false
	})

	fmt.Printf("%v\n", hands)
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	log.Println(total)
}

func isCurrentHandBetter(currentHand, bestHand handInfo) bool {
	currentHandCards, bestHandCards := currentHand.cards, bestHand.cards

	for i := 0; i < 5; i++ {
		if currentHandCards[i] == bestHandCards[i] {
			continue
		} else if currentHandCards[i] > bestHandCards[i] {
			return true
		}

		return false
	}

	return false
}

type handInfo struct {
	raw      string
	strength int
	cards    [5]int
	bid      int
}

var hands = []handInfo{}
var cardMap = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func parseHand(line string) {
	handS, bidS, _ := strings.Cut(line, " ")
	bid, err := strconv.Atoi(bidS)

	if err != nil {
		log.Fatal(err)
	}

	cardList := [5]int{}
	cardCountList := [13]int{}

	for i, c := range handS {
		value := cardMap[c]
		cardCountList[value-2]++
		cardList[i] = value
	}

	fmt.Printf("%v\n", cardCountList)

	strength := getHandStrength(cardCountList)
	hands = append(hands, handInfo{
		raw:      handS,
		strength: strength,
		bid:      bid,
		cards:    cardList,
	})
}

func getHandStrength(cardCountList [13]int) int {
	pairFound := false
	threeOfKindFound := false

	for _, count := range cardCountList {
		switch count {
		case 5:
			return 6
		case 4:
			return 5
		case 3:
			if pairFound {
				return 4
			}
			threeOfKindFound = true
		case 2:
			if threeOfKindFound {
				return 4
			} else if pairFound {
				return 2
			}
			pairFound = true
		}
	}

	if threeOfKindFound {
		return 3
	} else if pairFound {
		return 1
	}

	return 0
}
