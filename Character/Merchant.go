package character

import "math/rand"

//Fichier pour le Marchand, on y retrouve des fonctions d'achat. en relation avec le fichier inventory.

var Merchant = Character{
	Name:      "Marchand",
	Inventory: map[string]int{"Potion de soin +": 10, "Potion de soin": 15, "Potion de poison": 3, "Grimoire: boule de feu": 1, "Fourrure de loup": 5, "Peau de troll": 5, "Cuir de sanglier": 5, "Plume de corbeau": 5, "Épée d'entrainement": 1, "Manuel: coup d'épée": 1, "Potion de mana": 5, "Potion de mana +": 1},
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

func (merchant *Character) AddItems() {
	i := rand.Intn(100)
	j := rand.Intn(100)
	switch {
	case i < 25:
		merchant.Inventory[AllEquipement[i%len(AllEquipement)].Name] = 3
		merchant.Inventory[AllEquipement[j%len(AllEquipement)].Name] = 3
	case i > 25 && i < 50:
		merchant.Inventory[AllSpellBook[i%len(AllSpellBook)].Name] = 3
		merchant.Inventory[AllSpellBook[j%len(AllSpellBook)].Name] = 3
	case i > 50 && i < 75:
		merchant.Inventory[AllPotion[i%len(AllPotion)].Name] = 3
		merchant.Inventory[AllPotion[j%len(AllPotion)].Name] = 3
	case i > 75:
		merchant.Inventory[AllRessources[i%len(AllRessources)].Name] = 3
		merchant.Inventory[AllRessources[j%len(AllRessources)].Name] = 3
	}
}
