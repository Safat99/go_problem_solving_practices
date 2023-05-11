package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	var n, m, count int
	count = 0
	fmt.Scan(&n, &m)
	// fmt.Println("n and m are, ", n, m)
	for a := 0; a < 33; a++ {
		for b := 0; b < 33; b++ {
			if (a*a)+b == n && a+(b*b) == m {
				count++
			}
		}
	}
	fmt.Println(count)
}
