package pointers

import (
	"fmt"
	"example/pointers/types"
)





func Pointer() {
	user := types.User{UserEmail: "johndoe@gmail.com"}
	user.UpdateEmail("james milnaer")
	fmt.Println(user.Email())

	fmt.Println("player turn")
	num:=4

	player := types.Player{Health: 100}



	
	

	for i:=0;i<num;i++{
		 player.TakeDamageFromExplosion()
	}

	fmt.Println(player.NewPlayerHealth())
	
	
	

}
