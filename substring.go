package main

import "fmt"

func main() {

	var a, b string
	fmt.Scan(&a, &b)

	var c, d, pointer, start_index int
	var result bool
	result = false
	c = len(a)
	d = len(b)

	fmt.Println("len a and b are ", c, d)

	if d > c {
		fmt.Println("substring is false")
		return
	}

	for i := 0; i < c; i++ {
		if a[i] == b[pointer] {
			pointer++
			start_index = i

			for pointer < d {
				if a[start_index+1] == b[pointer] {
					pointer++
					start_index++
					fmt.Println("pointer value is: ", pointer)
				} else {
					break
				}
			}

			if pointer == d {
				fmt.Println("Substring is true")
				result = true
				break
			}
		} else {
			pointer = 0
			start_index = 0
		}
	}

	if result == false {
		fmt.Println("substring is false")
	}
}
