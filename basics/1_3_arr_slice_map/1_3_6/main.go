package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите положительное целое число: ")
	fmt.Scanf("%d", &number)
	if number < 0 {
		fmt.Println("Error. Введено отризательное число")
		return
	}
	result := countDigits(number)
	fmt.Printf("Количество цифр в числе %d:\n", number)
	for digit, count := range result {
		fmt.Printf("%d: %d\n", digit, count)
	}
}

func countDigits(number int) map[int]int {
	counter := make(map[int]int)
	for number > 0 {
		digit := number % 10 // получаем последнюю цифру беря остаток
		number /= 10         // убираем последнюю цифру сокращённая форма записи для операции деления с присваиванием.
		counter[digit]++     // увеличиваем счетчик для этой цифры
	}
	return counter
}
