package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var arr = make([][]string, n)
	var point = [4][2]int{}

	for i := range arr {
		arr[i] = make([]string, m)
	}

	fmt.Println("initial 2D array", arr)
	fmt.Println("point value are", point)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[i][j])
			if arr[i][j] == "*" {
				point[row] = append(point[row])
			}
		}
	}
	fmt.Println("Final 2D array", arr)

}
