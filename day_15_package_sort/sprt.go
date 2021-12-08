package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	Age  int
}

type UserSlice []User

// interface --> line 16-26
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
	users := []User{
		{"herii", 12},
		{"purnama", 21},
		{"sari", 22},
		{"retno", 35},
		{"angga", 9},
	}

	fmt.Println("sebelum sort--> ", users)
	sort.Sort(UserSlice(users))
	fmt.Println("setelah sort--> ", users)

}
