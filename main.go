package main

import (
	"fmt"
	"learn_go/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(15)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	fmt.Println(date.Day(), date.Month(), date.Year())
	event := calendar.Event{}
	event.SetYear(2020)
	event.SetMonth(10)
	event.SetDay(22)

	err = event.SetEvent("Hello this is New title!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Event(), event.Year(), event.Month(), date.Year())
	//// Coordinates
	//coordinates := geo.Coordinates{}
	//err = coordinates.SetLatitude(37.42)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = coordinates.SetLongitude(-122.08)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(coordinates)
	//// Landmark
	//location := geo.Landmark{}
	//err = location.SetName("The Googleplex")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = location.SetLatitude(37.24)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = location.SetLongitude(-122.08)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(location.Name())
	//fmt.Println(location.Longitude())
	//fmt.Println(location.Latitude())

}
