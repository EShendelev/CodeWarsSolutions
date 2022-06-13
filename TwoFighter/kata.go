package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func main() {

	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry"))

}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	// определим первого атакующего
	queue := fighter1.Name == firstAttacker
	// FIGHT!
	for {
		if queue {
			fighter2.Health = fighter2.Health - fighter1.DamagePerAttack
			queue = false
			fmt.Printf("%s attacks %s; %s now has %d health\n", fighter1.Name, fighter2.Name, fighter2.Name, fighter2.Health)
			switch {
			case fighter1.Health <= 0:
				return fighter2.Name
			case fighter2.Health <= 0:
				return fighter1.Name
			}

		}
		if !queue {
			fighter1.Health = fighter1.Health - fighter2.DamagePerAttack
			queue = true
			fmt.Printf("%s attacks %s; %s now has %d health\n", fighter2.Name, fighter1.Name, fighter1.Name, fighter1.Health)
			switch {
			case fighter1.Health <= 0:
				return fighter2.Name
			case fighter2.Health <= 0:
				return fighter1.Name
			}

		}
	}
}

// func newFighter(name string, health, dmg int) *Fighter {
// 	fighter := Fighter{Name: name, Health: health, DamagePerAttack: dmg}
// 	return &fighter
// }
