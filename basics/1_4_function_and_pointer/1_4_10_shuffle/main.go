package main

import (
	"fmt"
	"math/rand"
)

func shuffle(nums []int) {
	rnd.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
}

// генератор случайных чисел
var rnd *rand.Rand

// init - это специальная функция, которую Go вызывает до main()
// обычно используется для инициализации начального состояния приложения
func init() {
	// Создадим новый генератор случайных чисел
	// и будем использовать его вместо глобального:
	rnd = rand.New(rand.NewSource(42))
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	shuffle(nums)
	fmt.Println(nums)
}
