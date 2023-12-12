package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

type coordinate struct {
	x         int
	y         int
	direction rune
}

var startRegex = regexp.MustCompile("S")
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
	x, y := startCoordinate.x, startCoordinate.y
	var direction rune
	area := 0

	if pipes[loop[y][x+1]]['R'].direction != 0 {
		x++
		direction = 'R'
		area = y + 1
	} else if pipes[loop[y][x-1]]['L'].direction != 0 {
		x--
		direction = 'L'
		area = -(y + 1)
	} else if pipes[loop[y+1][x]]['D'].direction != 0 {
		y++
		direction = 'D'
	} else if pipes[loop[y-1][x]]['U'].direction != 0 {
		y--
		direction = 'U'
	}

	return move(0, area, x, y, direction)
}

func move(totalSteps int, area int, x int, y int, direction rune) int {
	totalSteps++
	currentShape := loop[y][x]

	if currentShape == 'S' {
		if area < 0 {
			area *= -1
		}

		// Honestly not sure why the +1 is needed here
		return area - (totalSteps / 2) + 1
	}

	currentPipe := pipes[currentShape][direction]

	if currentPipe.x != 0 {
		area += currentPipe.x * (y + 1)
	}

	return move(totalSteps, area, x+currentPipe.x, y+currentPipe.y, currentPipe.direction)
}

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
