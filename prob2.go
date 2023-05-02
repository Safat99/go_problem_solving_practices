package main

import (
	"fmt"
	"sort"
)

func main() {

	var nums = [4]int{}
	var count int
	fmt.Scan(&nums[0], &nums[1], &nums[2], &nums[3])
	sort.Ints(nums[:])

	for i := 0; i < 3; i++ {
		if nums[i] == nums[i+1] {
			count++
		}
	}

	fmt.Println(count)
}
