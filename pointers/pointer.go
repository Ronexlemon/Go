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
func Email(u User)string{
	return u.email
}
func (u *User) updateEmail(email string){
	u.email = email
	

}
func getUser()(*User,error){
	return nil,fmt.Errorf("lemonr")
}
func Pointer() {
	user := User{email: "johndoe@gmail.com"}
	user.updateEmail("lemonron@gmail.com")
	fmt.Println(getUser())

	fmt.Println(user.Email())
}
