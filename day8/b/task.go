package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var nodeNameRegexp = regexp.MustCompile("[A-Z]{3}")
var startNodeNameRegexp = regexp.MustCompile("A$")
var endNodeNameRegexp = regexp.MustCompile("Z$")
var numberRegex = regexp.MustCompile("[0-9]+")
var instructions []rune
var nodes = map[string]map[rune]string{}
var startNodes = []string{}

func main() {
	file, err := os.Open("./day8/input.txt")

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
	currentNodes := append([]string{}, startNodes...)
	stepAmounts := []int{}

	for i := 0; i <= len(instructions); i++ {
		if i == len(instructions) {
			i = 0
		}

		nextDirection := instructions[i]
		nextNodes := []string{}
		steps++

		for j := 0; j < len(currentNodes); j++ {
			node := currentNodes[j]
			nextNode := nodes[node][nextDirection]

			if endNodeNameRegexp.FindString(nextNode) != "" {
				stepAmounts = append(stepAmounts, steps)
				continue
			}

			nextNodes = append(nextNodes, nextNode)
		}

		if len(nextNodes) == 0 {
			break
		}

		currentNodes = nextNodes
	}

	return LCM(stepAmounts[0], stepAmounts[1], stepAmounts[2:]...)
}

func parseNode(line string) {
	nodeNames := nodeNameRegexp.FindAllString(line, -1)
	currentNode := nodeNames[0]

	nodes[currentNode] = map[rune]string{
		'L': nodeNames[1],
		'R': nodeNames[2],
	}

	if startNodeNameRegexp.FindString(currentNode) != "" {
		startNodes = append(startNodes, currentNode)
	}

}

// ----- Code below is from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
