package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(getConcatenation(nums))
}

func getConcatenation(nums []int) []int {
	result := make([]int, len(nums)*2) // создаем срез размером 2 * n получаемого среза
	for i := 0; i < len(nums); i++ {
		result[i] = nums[i]
		result[len(nums)+i] = nums[i]
	}
	return result
}
