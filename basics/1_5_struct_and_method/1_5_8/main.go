package main

import (
	"fmt"
)

type validator func(s string) bool

type password struct {
	value string
	validator
}

func minlen(length int) validator {
	return func(s string) bool {
		return len(s) >= length
	}
}

func (p *password) isValid() bool {
	return p.validator(p.value)
}

func main() {
	validator := minlen(5)
	p := password{value: "hello", validator: validator}
	fmt.Println(p.isValid())
}

//package main
//
//import "fmt"
//
//type validator func(s string) bool
//
//func minlen(length string) validator {
//	return func(s string) bool {
//		return len(s) >= len(length)
//	}
//}
//func main() {
//
//	s := validator(minlen("leet"))
//	fmt.Println(s)
//}
