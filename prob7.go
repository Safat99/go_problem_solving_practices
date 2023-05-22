package main

import "fmt"

func main() {
	var n, m, row_num, count int

	fmt.Scan(&n, &m)
	var arr = make([][]rune, n)
	var point = [4][2]int{}

	for i := range arr {
		arr[i] = make([]rune, m)
	}

	var input_arr = make([]rune, n*m)
	for i := 0; i < len(input_arr); i++ {
		fmt.Scanf("%c", &input_arr[i])
	}

	fmt.Println("input_arr ", input_arr)
	fmt.Printf("input_arr[0] is %c\n", input_arr[0])

	fmt.Println("initial 2D array", arr)
	fmt.Println("point value are", point)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// fmt.Scan(&arr[i][j])
			arr[i][j] = input_arr[count]
			count++

			if arr[i][j] == '*' {
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

	fmt.Println("Final 2D array", arr)
	fmt.Println("point array", point)

	for _, element := range point[3] {
		fmt.Printf("%d ", element+1)
	}

}
