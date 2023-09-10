package main

import (
	"fmt"
	"github.com/gracjang/gotraining/aod"
	"github.com/gracjang/gotraining/code100"
	"github.com/gracjang/gotraining/sudoku"
)

func main() {
	fmt.Println("Sudoku solution: ")
	sudoku.SolveSudoku()

	code100.SolveCh1("code100/puzzle.json")
	code100.SolveCh2("code100/puzzle2.json")
	code100.SolveCh3("code100/puzzle3.json")

	aod.Solve_d1()
	aod.Solve_d2()
}
