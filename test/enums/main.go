package main

import "fmt"

// weapon type
// axe
// sword
// wooden stick
// knife

func getDamage(weaponType string) int {
	switch weaponType {
	case "axe":
		return 100
	case "sword":
		return 90
	case "knife":
		return 40
	case "woodenStick":
		return 1
	default:
		panic("no weapon")
	}
}

func main() {
	fmt.Println("damage of weapon", getDamage("knife"))
}
