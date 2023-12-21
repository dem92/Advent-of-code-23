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
var galaxies = map[string][]int{}
var voidSize = 1000000 - 1

func main() {
	// file, err := os.Open("./day11/example.txt")
	// file, err := os.Open("../example.txt")
	file, err := os.Open("./day11/input.txt")
	// file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanText(file)
	total := getPathSum()

	file.Close()
	log.Println(total)
}

func getPathSum() int {
	total := 0
	galaxyValues := [][]int{}

	for _, value := range galaxies {
		galaxyValues = append(galaxyValues, value)
	}

	for i := 0; i < len(galaxyValues); i++ {
		galaxy := galaxyValues[i]
		y, x := galaxy[0], galaxy[1]

		for j := i + 1; j < len(galaxies); j++ {
			otherGalaxy := galaxyValues[j]
			yDiff := int(math.Abs(float64(otherGalaxy[0] - y)))
			xDiff := int(math.Abs(float64(otherGalaxy[1] - x)))
			diff := yDiff + xDiff
			total += diff
		}
	}

	return total
}

func scanText(file *os.File) {
	scanner := bufio.NewScanner(file)
	y := 0
	galaxyIndex := 0
	timesIncreased := 0

	for scanner.Scan() {
		text := scanner.Text()
		galaxyIndexes := galaxyRegex.FindAllStringIndex(text, -1)

		if galaxyIndexes == nil {
			timesIncreased++
		} else {
			for _, indexes := range galaxyIndexes {
				x := indexes[0]
				galaxies[fmt.Sprintf("%d %d", y, x)] = []int{y + timesIncreased*voidSize, x}
				galaxyIndex++
			}
		}

		universe = append(universe, []rune(text))
		y++
	}

	for i := 0; i < len(universe); i++ {
		fmt.Println(string(universe[i]))
	}

	expandUniverse()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func expandUniverse() {
	timesIncreased := 0

	for x := 0; x < len(universe[0]); x++ {
		galaxyFound := false

		for y := 0; y < len(universe); y++ {
			currentRune := universe[y][x]

			if currentRune == '.' {
				continue
			}

			galaxyFound = true
			galaxies[fmt.Sprintf("%d %d", y, x)][1] += voidSize * timesIncreased
		}

		if !galaxyFound {
			timesIncreased++
		}
	}
}
