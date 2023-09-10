package code100

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
)

type Node struct {
	ID    uuid.UUID `json:"id"`
	Value int       `json:"value"`
	Next  string    `json:"next"`
}

type Puzzle3 struct {
	Nodes []Node    `json:"linkedList"`
	Top   uuid.UUID `json:"top"`
}

func SolveCh3(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	puzzle := Puzzle3{}
	err = json.Unmarshal(content, &puzzle)
	if err != nil {
		log.Fatal(err)
	}
	puzzle.Traverse()
}

func (p *Puzzle3) Traverse() {
	nodesMap := make(map[uuid.UUID]*Node)
	for i := range p.Nodes {
		nodesMap[p.Nodes[i].ID] = &p.Nodes[i]
	}

	node := nodesMap[p.Top]

	var values []int
	for node != nil {
		values = append(values, node.Value)
		nextID, err := uuid.Parse(node.Next)
		if err != nil {
			break
		}

		node = nodesMap[nextID]
	}

	log.Println("Challange3:", values)
}
