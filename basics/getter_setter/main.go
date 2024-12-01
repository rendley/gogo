package main

import (
	"errors"
	"fmt"
	"learn_go/calendar"
	"log"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

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
}
