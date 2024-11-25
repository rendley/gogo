package main

import (
	"fmt"
	"strings"
)

func main() {
	var source string
	var times int
	fmt.Scan(&source, &times)

	var result string
	for i := 0; i < times; i++ {
		result += source
	}
	fmt.Println(result)

	fmt.Println("############")

	var result2 string
	for range times {
		result2 += source
	}
	fmt.Println(result2)

	fmt.Println("############")

	fmt.Println(strings.Repeat(source, times))

}
