package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := readString()
	//str := "Анна-Мария 2 Волхоновская"
	words := strings.Fields(str) // получаем массив ["Today", "I", learned"]
	var res string
	for _, char := range words { // при итерации idx = int, char = string, char[0] - rune
		firsRune, _ := utf8.DecodeRuneInString(char) // return rune and index []rune(word)[0]
		if unicode.IsLetter(firsRune) {              // check letter, don't digit
			res += strings.ToUpper(string(firsRune)) // concatenate first letter in resalt
		}
	}
	fmt.Print(res)
}

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}

// another example
//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//"strings"
//"unicode"
//)
//
//func main() {
//	phrase := readString()
//
//	words := strings.Fields(phrase)
//	var abbr []rune
//	for _, word := range words {
//		letter := []rune(word)[0]
//		if !unicode.IsLetter(letter) {
//			continue
//		}
//		abbr = append(abbr, unicode.ToUpper(letter))
//	}
//
//	fmt.Println(string(abbr))
//}
//
//// readString читает строку из `os.Stdin` и возвращает ее
//func readString() string {
//	rdr := bufio.NewReader(os.Stdin)
//	str, _ := rdr.ReadString('\n')
//	return str
//}
