package main

import "fmt"

func main() {
	var n, m int
	n = 3
	m = 3
	var matrix = make([][]string, n)
	fmt.Println("matrix value", matrix)

	for i := range matrix {
		matrix[i] = make([]string, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] = "safat"
		}
	}

	fmt.Println()
	fmt.Println("matrix value", matrix)

}
