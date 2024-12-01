package main

import "fmt"

type User struct {
	name string
	age  int
	next *User
	prev *User
}

func main() {
	user1 := User{"A", 1, nil, nil}
	user2 := User{"B", 2, nil, nil}
	user3 := User{"C", 3, nil, nil}

	user1.next = &user2
	user2.next = &user3
	user2.prev = &user1
	user3.prev = &user2

	user := &user1
	for user != nil {
		fmt.Println(user.name, user.age)
		user = user.next

	}

	user = &user3
	for user != nil {
		fmt.Println(user.name, user.age)
		user = user.prev
	}
}
