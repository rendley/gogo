// https://leetcode.com/problems/reverse-string/description/
package main

import (
	"fmt"
)

func main() {
	str := "hello"
	strRevRune := reverseStringRune(str)
	fmt.Println(strRevRune)
	strRevByte := reverseStringByte(str)
	fmt.Println(strRevByte)
}

func reverseStringRune(s string) string {
	rns := []rune(s)
	left := 0
	right := len(s) - 1
	for left < right {
		rns[left], rns[right] = rns[right], rns[left]
		left++
		right--
	}
	return string(rns)
}

func reverseStringByte(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
