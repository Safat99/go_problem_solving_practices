package main

import "fmt"

func main() {

	var n, j, k, l, sum int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&j, &k, &l)
		if (j + k + l) >= 2 {
			sum++
		}
	}

	fmt.Println(sum)
}
