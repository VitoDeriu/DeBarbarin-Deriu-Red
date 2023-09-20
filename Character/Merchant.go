package character

import "fmt"

var Merchant Character

func InitialiseMerchant() {

	Merchant.name = "Marchant"
	Merchant.inventory = map[string]int{"Potion de soin": 1, "Truc de fou": 2, "un jean-luc": 1, "Item de malade": 5}
}

func DisplayBoutique() {
	for item, nb := range Merchant.inventory {
		fmt.Println(item, " : ", nb)
	}
}

func (char *Character) BuyPotion(s Potion) {

	if Merchant.inventory[s.name] > 0 {
		char.addPotion(s)
		println("achat de 1 ", s.name) // il faut ajouter la perte de piece d'or
	} else {
		println("je n'ai plus de ", s.name, " frérot !")
	}
}

func (char *Character) BuySpellBook(s SpellBook) {

	if Merchant.inventory[s.name] > 0 {
		char.addSpellBook(s)
		println("achat de 1 ", s.name) // il faut ajouter la perte de piece d'or
	} else {
		println("je n'ai plus de ", s.name, " frérot !")
	}
}
