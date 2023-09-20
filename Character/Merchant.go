package character

import "fmt"

var Merchant Character

func InitialiseMerchant() {

	Merchant.Name = "Marchant"
	Merchant.Inventory = map[string]int{"Potion de soin": 1, "Truc de fou": 2, "un jean-luc": 1, "Item de malade": 5}
}

func DisplayBoutique() {
	for item, nb := range Merchant.Inventory {
		fmt.Println(item, " : ", nb)
	}
}

func (char *Character) BuyPotion(s Potion) {

	if Merchant.Inventory[s.Name] > 0 {
		char.AddPotion(s)
		println("achat de 1 ", s.Name) // il faut ajouter la perte de piece d'or
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}

func (char *Character) BuySpellBook(s SpellBook) {

	if Merchant.Inventory[s.Name] > 0 {
		char.AddSpellBook(s)
		println("achat de 1 ", s.Name) // il faut ajouter la perte de piece d'or
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}
