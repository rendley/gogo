package main

import (
	"fmt"
)

func main() {
	for idx, char := range "ого" {
		fmt.Println(idx, char, string(char))
		fmt.Println("")
	}

	str := "ого"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %d\n", i, str[i])
	}

}
