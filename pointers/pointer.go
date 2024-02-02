package pointers

import (
	"fmt"
	"example/pointers/types"
)


func takeDamageFromExplosion(player types.Player){
	fmt.Println("Player is taking damage from an explosion")
	explosionDamage :=10

	player.Health-= explosionDamage
}


func Pointer() {
	user := types.User{UserEmail: "johndoe@gmail.com"}
	user.UpdateEmail("james milnaer")
	
	

	fmt.Println(user.Email())
}
