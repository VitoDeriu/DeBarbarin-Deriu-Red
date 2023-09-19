package character

var Merchant Character

func InitialiseMerchant() {

	Merchant.name = "Marchant"
	Merchant.inventory = map[string]int{"Potion de soin": 1, "Truc de fou": 2, "Item de malade": 5}
}

type Item struct {
	name     string
	typeItem string
	price    int
}

func (char *Character) BuyItem(s Item) {

	InitialiseMerchant()

	if Merchant.inventory[s.name] > 0 {

		char.inventory[s.name] += 1
		println("achat de 1 ", s.name) // il faut ajouter la perte de piece d'or

	} else {

		println("je n'ai plus de ", s.name, " frérot !")
	}
}
