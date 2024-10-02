package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func change(a *User, b *User) {
	a.Name, b.Name = b.Name, a.Name
}

func sorting(list []User) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i].Name > list[j].Name {
				change(&list[i], &list[j])
			}
		}
	}
}

func main() {
	list := []User{
		{"Tom", 10},
		{"Jerry", 15},
		{"Mike", 20},
		{"Alice", 25},
	}

	sorting(list)
	for _, user := range list {
		fmt.Println(user)
	}
}
