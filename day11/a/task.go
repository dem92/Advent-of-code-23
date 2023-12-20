package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
)

var galaxyRegex = regexp.MustCompile("#")
var universe = [][]rune{}

// y, x
var galaxies = [][]int{}

func main() {
	// file, err := os.Open("./day11/example.txt")
	// file, err := os.Open("../example.txt")
	file, err := os.Open("./day11/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanText(file)
	findGalaxies()
	total := getPathSum()

	file.Close()
	log.Println(total)
}

func getPathSum() int {
	total := 0

	for i := 0; i < len(galaxies); i++ {
		galaxy := galaxies[i]
		y, x := galaxy[0], galaxy[1]

		for j := i + 1; j < len(galaxies); j++ {
			otherGalaxy := galaxies[j]
			yDiff := otherGalaxy[0] - y
			xDiff := int(math.Abs(float64(otherGalaxy[1] - x)))
			diff := yDiff + xDiff
			total += diff
		}
	}

	return total
}

func findGalaxies() {
	for i := 0; i < len(universe); i++ {
		for j, r := range universe[i] {
			if r == '#' {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}
}

func scanText(file *os.File) {
	scanner := bufio.NewScanner(file)
	lineIndex := 0

	for scanner.Scan() {
		text := scanner.Text()
		universe = append(universe, []rune(text))

		if galaxy := galaxyRegex.FindString(text); galaxy == "" {
			universe = append(universe, []rune(text))
		}

		lineIndex++
	}
	expandUniverse()

	for i := 0; i < len(universe); i++ {
		fmt.Println(string(universe[i]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func expandUniverse() {
	for i := 0; i < len(universe[0]); i++ {
		galaxyFound := false

		for j := 0; j < len(universe); j++ {
			if universe[j][i] == '#' {
				galaxyFound = true
				break
			}
		}

		if !galaxyFound {
			for j := 0; j < len(universe); j++ {
				currentLine := universe[j]

				if len(universe[0]) == i {
					currentLine = append(currentLine, '.')
				} else {
					currentLine = append(currentLine[:i+1], currentLine[i:]...)
					currentLine[i] = '.'
				}

				universe[j] = currentLine
			}

			i++
		}
	}
}
