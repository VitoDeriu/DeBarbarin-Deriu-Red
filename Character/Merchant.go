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

func (char *Character) BuyItem(s Item) {

	if Merchant.inventory[s.name] > 0 {
		char.addInventory(s)
		println("achat de 1 ", s.name) // il faut ajouter la perte de piece d'or
	} else {
		println("je n'ai plus de ", s.name, " frérot !")
	}
}
