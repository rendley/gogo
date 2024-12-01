package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// predicate — это любая функция, которая принимает int и возвращает bool.
// В main такая функция передается как аргумент в функцию filter.
// функция принимает два агрумента срез iterable и predicate функцию
func filter(predicate func(int) bool, iterable []int) []int {
	var result []int
	for _, value := range iterable { // итерируемся по срезу
		if predicate(value) { // возвращается true ecли четное через придикат
			result = append(result, value) // добавляем в массив
		}
	}
	return result
}

func main() {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//src := readInput()
	isEven := func(x int) bool { return x%2 == 0 } //Предикат, который проверяет, является ли число четным
	evenNumber := filter(isEven, src)
	fmt.Println("Четные числа:", evenNumber)

	//// Еще один предикат, который проверяет, делится ли число на 3 без остатка
	//isDivisibleByThree := func(x int) bool { return x%3 == 0 }
	//
	//// Фильтруем исходный срез с использованием предиката isDivisibleByThree
	//numbersDivisibleByThree := filter(slice, isDivisibleByThree)
	//fmt.Println("Числа, делящиеся на 3:", numbersDivisibleByThree)

}

func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}

//// Изменение поведения бизнес логики (doSomething) из функции main (из клиента)
//// не залезая в бизнес логику
//func doSomething(callback func(int, int) int, s string) int {
//	result := callback(10, 2)
//	fmt.Println(s)
//	return result
//}
//
//// Изменение
//func main() {
//	sumCallback := func(n1, n2 int) int {
//		return n1 + n2
//	}
//	result := doSomething(sumCallback, "sum")
//	fmt.Println(result)
//
//	fmt.Println("##########################")
//
//	multipleCallback := func(n1, n2 int) int { return n1 * n2 }
//	multipleResult := doSomething(multipleCallback, "multiple")
//	fmt.Println(multipleResult)
//}
