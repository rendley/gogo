// https://stepik.org/lesson/526869/step/11?thread=solutions&unit=519588

package main

import (
	"fmt"
	"os"
)

func normalize(value ...*float64) {

	// Вычисление суммы всех значений
	sum := 0.0
	for _, v := range value {
		sum += *v
	}
	// Нормализация значений
	for i, _ := range value {
		*value[i] /= sum
	}
	// or this variant
	//for _, v := range vals {
	//	*v = *v / sum
	//}

}

func main() {
	a, b, c, d := 1.0, 2.0, 3.0, 4.0
	normalize(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	fmt.Println("PASS")
	os.Exit(0)

}
