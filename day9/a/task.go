package main

import (
	"aoc23/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day9/input.txt")
	// file, err := os.Open("./day9/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		total += predictValue(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(total)
}

func predictValue(line string) int {
	rawValues := strings.Split(line, " ")
	values := []int{}

	for _, val := range rawValues {
		values = append(values, utils.ConvertStringToNumber(val))
	}

	return getPrediction([][]int{values})
}

func getPrediction(allValues [][]int) int {
	prevRow := allValues[len(allValues)-1]
	differences := []int{}
	allZeros := true

	for i := 1; i < len(prevRow); i++ {
		diff := prevRow[i] - prevRow[i-1]
		differences = append(differences, diff)

		if diff != 0 {
			allZeros = false
		}
	}

	if !allZeros {
		allValues = append(allValues, differences)
		return getPrediction(allValues)
	}

	prediction := 0

	for reverseI := len(allValues) - 1; reverseI >= 0; reverseI-- {
		currentRow := allValues[reverseI]
		lastValue := currentRow[len(currentRow)-1]
		prediction = prediction + lastValue
	}

	return prediction
}
