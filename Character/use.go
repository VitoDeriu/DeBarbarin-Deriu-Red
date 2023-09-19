package character

import "fmt"

func (char *Character) TakePotion() {
	var havePotion bool
	for index := range char.inventory {
		if index == "Potion de vie" {
			havePotion = true
		}
	}
	if havePotion {
		if char.hp == char.hpMax {
			fmt.Println("Je n'ai pas soif !")
			return
		} else if char.hpMax-char.hp < 20 {
			char.hp = char.hpMax
		} else {
			char.hp += 20
		}
		char.inventory["Potion de vie"] -= 1
		if char.inventory["Potion de vie"] == 0 {
			delete(char.inventory, "Potion de vie")
		}
		fmt.Println("+20 hp")
		fmt.Println(char.hp, "/", char.hpMax)
		return
	}
	fmt.Println(char.name, " ➵  Je n'ai pas de potion !")
}

func (char *Character) addInventory(s string) {
	char.inventory[s] += 1
}

func (char *Character) removeInventory(s string) {
	if char.inventory[s] < 0 {
		char.inventory[s] -= 1
		if char.inventory[s] == 0 {
			delete(char.inventory, s)
		}
	} else {
		println("Impossible de retirer un objet que tu ne possède pas")
	}
}
