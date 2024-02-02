package types

import (
	"fmt"
)

type Player struct {
	Health int
}

func(player *Player) TakeDamageFromExplosion() {
	fmt.Println("Player is taking damage from an explosion")
	explosionDamage := 10

	player.Health -= explosionDamage
}

func (p Player) NewPlayerHealth() int {
	return p.Health
}
