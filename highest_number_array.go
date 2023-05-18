package main

import "fmt"

func main() {
	var n, mid, largest int
	fmt.Scan(&n)

	if n%2 == 0 {
		mid = n / 2
	} else {
		mid = n/2 + 1
	}

	// fmt.Println("n and mid are", n, mid)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := mid + 1; i < n; i++ {
		p := n - i

		if largest < arr[i] {
			largest = arr[i]
		}
		if largest < arr[p] {
			largest = arr[p]
		}
		fmt.Println("i and p are: ", i, p)
	}

	fmt.Println("largest number", largest)
}
