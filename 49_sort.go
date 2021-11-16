package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

/*
for the sort package to work, it needs 3 method as the part of its interface.
They are Len(), Less(), and Swap()
*/

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := UserSlice{
		{"Manuel", 40},
		{"Theodore", 20},
		{"Leleuly", 30},
	}

	sort.Sort(users)

	fmt.Println(users)
}