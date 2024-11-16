package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	var matrix [][]string

	for _, arg := range args {
		var row []string
		for i := 0; i < len(arg); i++ {
			row = append(row, string(arg[i]))
		}
		matrix = append(matrix, row)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for l := 0; l <len(matrix[j]) ; l++ {
				
			}
			for k := '0'; k <= '9'; k++ {
				if matrix[i][j] == "."  &&   matrix[i][j] ==   matrix[i][k]  {
					matrix[i][j] = string(i)
				}
			}
		}
	}

	fmt.Println(matrix)
}

