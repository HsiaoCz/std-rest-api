package main

import "fmt"

// weapon type
// axe
// sword
// wooden stick
// knife

type WeaponType int

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "SWORD"
	case Axe:
		return "AXE"
	case WoodenStick:
		return "WOODENSTICK"
	case Knife:
		return "KNIFE"
	default:
		panic("need the weapon type")
	}
}

const (
	Axe WeaponType = iota
	Sword
	WoodenStick
	Knife
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case Knife:
		return 40
	case WoodenStick:
		return 1
	default:
		panic("no weapon")
	}
}

func main() {
	fmt.Println("damage of weapon", getDamage(Axe))
	fmt.Printf("weapon :(%d)  and damage : (%d)\n", Axe, getDamage(Axe))
}
