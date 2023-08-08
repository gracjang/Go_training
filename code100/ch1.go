package main

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

func init() {
	content, err := os.ReadFile("puzzle.json")
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
