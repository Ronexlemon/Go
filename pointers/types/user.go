package types

type User struct {
	UserEmail    string
	Username string
	Age      int
}



func (u User) Email() string {
	return u.UserEmail
}

func (u *User) UpdateEmail(email string) {
	u.UserEmail = email

}
