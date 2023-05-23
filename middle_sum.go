package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var input_arr = make([][]int, n)
	for i := range input_arr {
		input_arr[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&input_arr[i][j])
		}
	}

	base := ((n + 1) / 2) - 1
	var up, down, left, right, u_l, d_r, u_r, d_l int

	for i := 1; i <= base; i++ {
		up += input_arr[base-i][base]
		down += input_arr[base+i][base]
		left += input_arr[base][base-i]
		right += input_arr[base][base+i]
		u_l += input_arr[base-i][base-i]
		d_r += input_arr[base+i][base+i]
		u_r += input_arr[base-i][base+i]
		d_l += input_arr[base+i][base-i]
		fmt.Printf("up %d down %d left %d right %d\n", up, down, left, right)
	}
	sum := input_arr[base][base] + up + down + left + right + u_l + d_r + u_r + d_l
	fmt.Println("sum ", sum)

}
