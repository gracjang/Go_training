package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	partOne(input)

}

func partOne(input []byte) {
	lines := strings.Split(string(input), "\n")
	elfCalories := []int{}
	caloriesCounter := 0

	for _, line := range lines {
		if line == "" {
			elfCalories = append(elfCalories, caloriesCounter)
			caloriesCounter = 0
			continue
		}

		calories, _ := strconv.Atoi(line)
		caloriesCounter += calories
	}

	max := max(elfCalories)
	counter := partTwo(elfCalories)

	log.Println("Day1P1:", max)
	log.Println("Day1P2:", counter)
}

func partTwo(input []int) int {
	sort.Ints(input)
	topThree := input[len(input)-3:]
	var counter int
	for _, num := range topThree {
		counter += num
	}

	return counter
}

func max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	sort.Ints(numbers)
	return numbers[len(numbers)-1]
}
