package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("hello")
	// var s string = "hala_madrid"
	var s string
	var count int = 1

	fmt.Scan(&s)

	strrunes := []rune(s)
	sort.Slice(strrunes, func(i, j int) bool {
		return strrunes[i] < strrunes[j]
	})

	// fmt.Println("sorted string is: ", strrunes)

	for i := 0; i < len(strrunes)-2; i++ {
		if strrunes[i] != strrunes[i+1] {
			count++
		}
	}

	// fmt.Println("count value: ", count)

	if count%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGONER HIM!")
	}

}
