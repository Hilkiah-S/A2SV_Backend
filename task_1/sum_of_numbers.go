package main

import (
	"fmt"
)

func sum_nums(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum

}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(sum_nums(arr))

}
