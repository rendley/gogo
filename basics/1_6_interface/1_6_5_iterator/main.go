package main

import "fmt"

type element interface{}

type iterator struct {
	value []int
}

type iterator interface {
	next()
	val()
}

func (it iterator) next() {
	return
}

func (it iterator) val() {
	return
}

func iterate(it iterator) {
	for it.next() {
		curr := it.val()
		fmt.Println(curr)
	}
}
