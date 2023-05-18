// LLPS (largest lexicographic palidromic subsequence)
package main

import "fmt"

func main() {
	// fmt.Println("Hello")

	var s string
	var cl, output string

	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {

		if cl < s[i:i+1] {
			cl = s[i : i+1]
			output = cl
		} else if cl == s[i:i+1] {
			output += cl
		}

		// fmt.Println("i, cl and output are", i, cl, output)

	}

	// fmt.Println("===>>>>final output is ", output)
	fmt.Println(output)
}
