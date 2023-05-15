package main

import "fmt"

func main() {

	var a, b string
	fmt.Scan(&a, &b)

	var c, d, pointer, start_index int
	c = len(a)
	d = len(b)

	fmt.Println("len a and b are ", c, d)

	for i := 0; i < c; i++ {
		if a[i] == b[pointer] {
			pointer++
			start_index = i
			break
		} else {
			continue
		}
	}

	fmt.Println("pointer value: ", pointer)

	for pointer < d {
		if a[start_index+1] == b[pointer] {
			pointer++
			start_index++
		} else {
			break
		}
	}

	if pointer == d {
		fmt.Println("Substring is true")
	} else {
		fmt.Println("substring is false")
	}

}
