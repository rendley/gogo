package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	m, err := time.ParseDuration(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s, "=", m.Minutes(), "min")
	fmt.Printf("%s = %v min\n", s, m.Minutes())
	fmt.Printf("%s = %g min\n", s, m.Minutes())
	fmt.Printf("%s = %.0f min\n", s, m.Minutes())
}

// спецификатор %g работает как " %e для огромных экспонент, %f в противном случае".
// А вот спецификатор %v это универсальный спецификатор, который для типа boolean аналогичен %t,
// для целочисленных типов - %d,
// для чисел с плавающей точкой - %g,
// для строк - %s
// %T — тип переменной.
