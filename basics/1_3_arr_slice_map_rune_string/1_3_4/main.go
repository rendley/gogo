package main

import "fmt"

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width) // Eyjafjallajokull 6

	res := text
	if len(text) > width {
		res = res[:width] + "..."
	}
	fmt.Println(res)
}

//func main() {
//	text := "hel"
//	width := 6
//	result := truncateText(text, width)
//	fmt.Println(result)
//}
//
//func truncateText(text string, width int) string {
//	if len(text) <= width {
//		return text
//	}
//	return text[:width] + "..."
//}
