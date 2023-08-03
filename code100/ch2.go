package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
)

type Puzzle2 struct {
	JsonStr []string `json:"string"`
}

func main() {

	content, err := os.ReadFile("puzzle2.json") // replace with your JSON file's path
	if err != nil {
		log.Fatal(err)
	}

	puzzle := Puzzle2{}
	err = json.Unmarshal(content, &puzzle.JsonStr)
	if err != nil {
		log.Fatal(err)
	}

	anagrams := make(map[string][]string)
	for _, word := range puzzle.JsonStr {
		sortedWords := sortWords(word)
		anagrams[sortedWords] = append(anagrams[sortedWords], word)
	}

	var result []string
	for _, words := range anagrams {
		if len(words) > 1 {
			result = append(result, words...)
		}
	}
	sort.Strings(result)
	log.Println("Challenge 2:", result)
}

func sortWords(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
