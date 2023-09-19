package character

func InitialiseMerchant() {
	var Merchant Character

	Merchant.name = "Marchant"
	Merchant.inventory = map[string]int{"Potion de soin": 1, "Truc de fou": 2, "Item de malade": 5}
}

func () BuyPotion() {
	if Merchant.inventory["Potion de soin"] > 0 {

		p.inventory["Potion de soin"] += 1
		println("Vous avez acheté une potion de soin (Gratuit)")
		Merchant.inventory["Potion de soin"] -= 1

	} else {

		println("J'ai plus de potion frérot")
	}
}
