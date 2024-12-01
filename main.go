package main

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.1f litters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.0f millilitres equals %0.3f gallons\n", water, water.ToGallons())
	milk := Gallons(2)
	fmt.Printf("%0.1f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.1f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())
	ml := Milliliters(500)
	fmt.Printf("%0.0f milliliters equals %0.3f liters\n", ml, ml.ToLiters())
	l := Liters(3)
	fmt.Printf("%0.1f liters equal %0.1f milliliters", l, l.ToMilliliters())

}
