package main

import "fmt"

func main() {
	var n, m, row_num int
	fmt.Scan(&n, &m)
	var arr = make([][]string, n)
	var point = [4][2]int{}

	for i := range arr {
		arr[i] = make([]string, m)
	}

	// fmt.Println("initial 2D array", arr)
	// fmt.Println("point value are", point)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[i][j])
			if arr[i][j] == "*" {
				point[row_num][0] = i
				point[row_num][1] = j
				row_num++
			}
		}
	}

	if point[0][0] == point[1][0] {
		point[3][0] = point[2][0]

		if point[2][1] == point[1][1] {
			point[3][1] = point[0][1]
		} else {
			point[3][1] = point[1][1]
		}

	} else {

		if point[2][0] == point[1][0] {
			point[3][0] = point[0][0]

			if point[0][1] == point[1][1] {
				point[3][1] = point[2][1]
			} else {
				point[3][1] = point[1][1]
			}

		} else {
			point[3][0] = point[1][0]
			point[3][1] = point[0][1]
		}
	}

	// fmt.Println("Final 2D array", arr)
	// fmt.Println("point array", point)

	for _, element := range point[3] {
		fmt.Printf("%d ", element+1)
	}

}
