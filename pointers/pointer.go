package pointers

import (
	"example/pointers/hello"
	"example/pointers/types"
	"fmt"
)

func Pointer() {
	user := types.User{UserEmail: "johndoe@gmail.com"}
	user.UpdateEmail("james milnaer")
	fmt.Println(user.Email())

	fmt.Println("player turn")
	num := 4

	player := types.Player{Health: 100}

	for i := 0; i < num; i++ {
		player.TakeDamageFromExplosion()
	}

	fmt.Println(player.NewPlayerHealth())

}
func Pointer2() {
	hello.Hello()
}

func PointerSchool() {
	school := types.School{Name: "JKUAT", NumberOfStudents: 30}

	fmt.Printf("School details before update name: %v and students: %v:", school.Name, school.NumberOfStudents)
	school.UpDateSchool("UON", 100)
	fmt.Printf("School details After update name: %v and students: %v:", school.Name, school.NumberOfStudents)

}
