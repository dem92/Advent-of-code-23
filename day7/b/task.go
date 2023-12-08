package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type handInfo struct {
	raw      string
	strength int
	cards    [5]int
	bid      int
}

var numberRegex = regexp.MustCompile("[0-9]+")

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

func main() {
	file, err := os.Open("./day7/input.txt")
	// file, err := os.Open("./day7/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		parseHand(text)
	}

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
		} else if currentHandCards[i] > bestHandCards[i] && currentHandCards[i] != 11 {
			return true
		} else if bestHandCards[i] == 11 && currentHandCards[i] != 11 {
			return true
		}

		return false
	}

	return false
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

	strength := getHandStrength(cardCountList)
	hands = append(hands, handInfo{
		raw:      handS,
		strength: strength,
		bid:      bid,
		cards:    cardList,
	})
}
func getHandStrength(cardCountList [13]int) int {
	baseStrength, jokerCount := getHandBaseStrength(cardCountList)

	switch jokerCount {
	case 5: // JJJJJ
		return 6
	case 4: // JJJJX
		return 6
	case 3:
		if baseStrength == 1 { // JJJXX
			return 6
		}
		return 5 // JJJXY
	case 2:
		if baseStrength == 1 { // JJXXY
			return 5
		} else if baseStrength == 3 { // JJXXX
			return 6
		}
		return 3 // JJXYZ
	case 1:
		if baseStrength == 1 { // JXXYZ
			return 3
		} else if baseStrength == 2 { // JXXYY
			return 4
		} else if baseStrength == 3 { // JXXXY
			return 5
		} else if baseStrength == 5 { // JXXXX
			return 6
		}
		return 1 // JXYZÆ
	}

	return baseStrength // XYZÆØ
}

func getHandBaseStrength(cardCountList [13]int) (int, int) {
	jokerCount := 0
	pairsFound := 0
	threeOfKindFound := false
	fourOfKindFound := false

	for i, count := range cardCountList {
		if i == 9 { // Joker index
			jokerCount = count
			continue
		}
		switch count {
		case 5:
			return 6, jokerCount
		case 4:
			fourOfKindFound = true
		case 3:
			threeOfKindFound = true
		case 2:
			pairsFound++
		}
	}

	if fourOfKindFound {
		return 5, jokerCount
	}

	if threeOfKindFound {
		if pairsFound == 1 {
			return 4, jokerCount
		}

		return 3, jokerCount
	}

	return pairsFound, jokerCount
}
