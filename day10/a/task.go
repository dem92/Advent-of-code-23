package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type coordinate struct {
	x         int
	y         int
	direction rune
}

var loop = [][]rune{}
var startCoordinate coordinate

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
var pipes = map[rune]map[rune]coordinate{
	'|': {
		'U': {
			x:         0,
			y:         -1,
			direction: 'U',
		},
		'D': {
			x:         0,
			y:         1,
			direction: 'D',
		},
	},
	'-': {
		'L': {
			x:         -1,
			y:         0,
			direction: 'L',
		},
		'R': {
			x:         1,
			y:         0,
			direction: 'R',
		},
	},
	'L': {
		'D': {
			x:         1,
			y:         0,
			direction: 'R',
		},
		'L': {
			x:         0,
			y:         -1,
			direction: 'U',
		},
	},
	'J': {
		'D': {
			x:         -1,
			y:         0,
			direction: 'L',
		},
		'R': {
			x:         0,
			y:         -1,
			direction: 'U',
		},
	},
	'7': {
		'U': {
			x:         -1,
			y:         0,
			direction: 'L',
		},
		'R': {
			x:         0,
			y:         1,
			direction: 'D',
		},
	},
	'F': {
		'U': {
			x:         1,
			y:         0,
			direction: 'R',
		},
		'L': {
			x:         0,
			y:         1,
			direction: 'D',
		},
	},
}

func main() {
	// file, err := os.Open("./day10/example.txt")
	// file, err := os.Open("../example.txt")
	file, err := os.Open("./day10/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanText(file)
	total := getSteps()

	file.Close()
	log.Println(total)
}

func getSteps() int {
	return move(0, startCoordinate.x+1, startCoordinate.y, 'R')
}

func move(totalSteps int, x int, y int, direction rune) int {
	totalSteps++
	currentShape := loop[y][x]
	readableRune := string(string(currentShape))
	fmt.Printf("Rune: %v, dir: %v\n", readableRune, string(direction))

	if currentShape == 'S' {
		return totalSteps / 2
	}

	currentPipe := pipes[currentShape][direction]
	return move(totalSteps, x+currentPipe.x, y+currentPipe.y, currentPipe.direction)
}

var startRegex = regexp.MustCompile("S")

func scanText(file *os.File) {
	scanner := bufio.NewScanner(file)
	line := 0

	for scanner.Scan() {
		text := scanner.Text()
		loop = append(loop, []rune(text))

		if startXIndex := startRegex.FindStringIndex(text); startXIndex != nil {
			startCoordinate = coordinate{x: startXIndex[0], y: line}
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
