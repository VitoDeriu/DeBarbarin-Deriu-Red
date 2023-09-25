package character

//Fichier pour le Marchand, on y retrouve des fonctions d'achat. en relation avec le fichier inventory.

var Merchant = Character{
	Name:      "Marchand",
	Inventory: map[string]int{"Potion de soin": 3, "Potion de poison": 3, "Grimoire boule de feu": 1, "Fourrure de loup": 5, "Peau de troll": 5, "Cuir de sanglier": 5, "Plume de corbeau": 5},
	Gold:      10000,
}

func (buyer *Character) BuyPotion(seller *Character, s Potion) {

	if seller.Inventory[s.Name] > 0 {
		if buyer.CheckGold(s.Price) || buyer == &Merchant {
			buyer.AddPotion(s)
			buyer.Gold -= s.Price
			seller.Gold += s.Price
			seller.RemovePotion(s)
			// println("achat de 1 ", s.Name)
		} else {
			return
		}
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}

func (buyer *Character) BuySpellBook(seller *Character, s SpellBook) {

	if seller.Inventory[s.Name] > 0 {
		if buyer.CheckGold(s.Price) || buyer == &Merchant {
			buyer.AddSpellBook(s)
			buyer.Gold -= s.Price
			seller.Gold += s.Price
			seller.RemoveSpellBook(s)
			// println("achat de 1 ", s.Name)
		} else {
			return
		}
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}

func (buyer *Character) BuyEquipement(seller *Character, s Equipement) {

	if seller.Inventory[s.Name] > 0 {
		if buyer.CheckGold(s.Price) || buyer == &Merchant {
			buyer.AddEquipement(s)
			buyer.Gold -= s.Price
			seller.Gold += s.Price
			seller.RemoveEquipement(s)
			// println("achat de 1 ", s.Name)
		} else {
			return
		}
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}

func (buyer *Character) BuyRessource(seller *Character, s Ressource) {

	if seller.Inventory[s.Name] > 0 {
		if buyer.CheckGold(s.Price) || buyer == &Merchant {
			buyer.AddRessource(s)
			buyer.Gold -= s.Price
			seller.Gold += s.Price
			seller.RemoveRessource(s)
			// println("achat de 1 ", s.Name)
		} else {
			return
		}
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}
