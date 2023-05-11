package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 != 0 {
		fmt.Println(-1)
		return
	}

	nums := make([]int, n)

	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}

	// fmt.Println("nums: ", nums)

	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			temp := nums[i]
			nums[i] = nums[i+1]
			nums[i+1] = temp
		} else {
			continue
		}
	}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}

}
