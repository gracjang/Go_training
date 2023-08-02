package main

import "fmt"

type SudokuGrid [9][9]int

func main() {
	grid := SudokuGrid{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	if sudokuSolver(&grid) {
		printSudoku(&grid)
	} else {
		fmt.Println("Sudoku not solved.")
	}
}

func findUnassigned(grid *SudokuGrid) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func sudokuSolver(grid *SudokuGrid) bool {
	i, j := findUnassigned(grid)
	if i == -1 && j == -1 {
		return true
	}

	nums := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, num := range nums {
		if rowCheck(grid, i, num) && colCheck(grid, j, num) && boxCheck(grid, i, j, num) {
			grid[i][j] = num
			if sudokuSolver(grid) {
				return true
			}
			grid[i][j] = 0
		}
	}
	return false
}

func boxCheck(grid *SudokuGrid, row int, col int, value int) bool {
	r := (row/3)*3 + 1
	c := (col/3)*3 + 1
	ranges := [3]int{-1, 0, 1}

	for _, i := range ranges {
		for _, j := range ranges {
			if grid[r+i][c+j] == value {
				return false
			}
		}
	}
	return true
}

func colCheck(grid *SudokuGrid, col int, value int) bool {
	for i := 0; i < 9; i++ {
		if value == grid[i][col] {
			return false
		}
	}
	return true
}

func rowCheck(grid *SudokuGrid, row int, value int) bool {
	for j := 0; j < 9; j++ {
		if value == grid[row][j] {
			return false
		}
	}
	return true
}

func printSudoku(grid *SudokuGrid) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
