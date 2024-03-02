package main

import "fmt"

// why we need the pointer

// if player is not a pointer we are adjusting the copy the player
// not the actual player
type Player struct {
	HP int
}

// Player's function
// get damage and recall HP
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Printf("player is taking damage. New HP -> %d\n", p.HP)
}

func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Printf("player is taking damage. New HP -> %d\n", player.HP)
}

func main() {
	player := &Player{
		HP: 100,
	}
	player.TakeDamage(10)
	// TakeDamage(player, 10)
	fmt.Println(player)
}
