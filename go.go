package main

import (
	"fmt"
	"os"
)

func IntMatrix() [][]int {
	args := os.Args[1:]
	var matrix [][]int
	for _, arg := range args {
		var row []int
		for i := 0; i < len(arg); i++ {
			if arg[i] == '.' {
				row = append(row, 0)
			} else {
				num := int(arg[i] - '0')
				row = append(row, num)
			}
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func correctNbr(matrix [][]int, row, column, nbr int) bool {
	// row
	for i := 0; i < 9; i++ {
		if matrix[row][i] == nbr {
			return false
		}
	}

	// column
	for i := 0; i < 9; i++ {
		if matrix[i][column] == nbr {
			return false
		}
	}

	// square
	x := (column / 3) * 3
	y := (row / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matrix[y+i][x+j] == nbr {
				return false
			}
		}
	}

	return true
}

func solution(matrix [][]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if matrix[row][column] == 0 {
				for nbr := 1; nbr <= 9; nbr++ {
					if correctNbr(matrix, row, column, nbr) {
						matrix[row][column] = nbr
						if solution(matrix) {
							return true
						}
						matrix[row][column] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	matrix := IntMatrix()
	if solution(matrix) {
		for i := 0; i < len(matrix); i++ {
			
			fmt.Println(matrix[i])
		}
	} else {
		fmt.Println("ERROR")
	}
}
