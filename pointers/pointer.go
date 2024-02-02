package pointers

import (
	"fmt"
)

type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func Pointer() {
	user := User{email: "johndoe@gmail.com"}

	fmt.Println(user.Email())
}
