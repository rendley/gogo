package main

import "fmt"

func main() {
	next := intSeq()
	fmt.Println(next())
	fmt.Println(next())
	gen := intSeq()
	fmt.Println(gen())
}

func intSeq() func() int {
	i := 0
	// внутренняя функция замыкания
	return func() int {
		i++
		return i
	}
}
