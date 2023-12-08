package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var numberRegex = regexp.MustCompile("[0-9]+")
var instructions []rune
var nodes = map[string]map[rune]string{}

func main() {
	file, err := os.Open("./day8/input.txt")
	// file, err := os.Open("./day8/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions = []rune(scanner.Text())
	scanner.Scan()

	for scanner.Scan() {
		text := scanner.Text()
		parseNode(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(getAmountOfSteps())
}

func getAmountOfSteps() int {
	steps := 0
	currentNode := "AAA"

	for i := 0; i <= len(instructions); i++ {
		if i == len(instructions) {
			i = 0
		}

		nextDirection := instructions[i]
		currentNode = nodes[currentNode][nextDirection]
		steps++

		if currentNode == "ZZZ" {
			break
		}
	}

	return steps
}

var nodeNameRegexp = regexp.MustCompile("[A-Z]{3}")

func parseNode(line string) {
	nodeNames := nodeNameRegexp.FindAllString(line, -1)

	nodes[nodeNames[0]] = map[rune]string{
		'L': nodeNames[1],
		'R': nodeNames[2],
	}
}
