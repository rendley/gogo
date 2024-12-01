package main

import (
	"fmt"
	"slices"
)

func main() {
	x := min(9.99, 3.14, 5.27)
	fmt.Println(x)
	// 3.14

	numbers := []int{0, 42, -10, 8}
	fmt.Println(slices.Max(numbers))

	s := max("one", "two", "three")
	fmt.Println(s)
	// "one"
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	clear(m)
	fmt.Printf("%#v\n", m)

	s1 := []string{"one", "two", "three"}
	clear(s1)
	fmt.Printf("%#v\n", s1)
}
