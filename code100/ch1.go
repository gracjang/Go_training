package code100

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Puzzle struct {
	JsonStr string `json:"string"`
}

func SolveCh1(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	puzzle := Puzzle{
		JsonStr: string(content),
	}

	regex := regexp.MustCompile(`\d+`)
	nums := regex.FindAllString(puzzle.JsonStr, -1)
	count := 0
	for _, num := range nums {

		numInt, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			continue
		}

		count += numInt
	}

	log.Println(`Challenge 1:`, count)
}
