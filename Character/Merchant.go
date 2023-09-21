package character

import "fmt"

//Fichier pour le marchant, on y retrouve des fonctions d'achat. en relation avec le fichier inventory.

var Merchant Character

func InitialiseMerchant() {

	Merchant.Name = "Marchant"
	Merchant.Inventory = map[string]int{"Potion de soin": 3, "Potion de poison": 3, "Grimoire boule de feu": 1, "Fourrure de loup": 5, "Peau de troll": 5, "Cuir de sanglier": 5, "Plume de corbeau": 5}
}

func DisplayBoutique() {
	for item, nb := range Merchant.Inventory {
		fmt.Println(item, " : ", nb)
	}
}

func (char *Character) BuyPotion(s Potion) {

	if Merchant.Inventory[s.Name] > 0 {
		char.AddPotion(s)
		char.CheckGold(s.Price)
		char.Gold -= s.Price
		println("achat de 1 ", s.Name)
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}

func (char *Character) BuySpellBook(s SpellBook) {

	if Merchant.Inventory[s.Name] > 0 {
		char.AddSpellBook(s)
		char.CheckGold(s.Price)
		char.Gold -= s.Price
		println("achat de 1 ", s.Name)
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}

func (char *Character) BuyEquipement(s Equipement) {

	if Merchant.Inventory[s.Name] > 0 {
		char.AddEquipement(s)
		char.CheckGold(s.Price)
		char.Gold -= s.Price
		println("achat de 1 ", s.Name)
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}

func (char *Character) BuyRessource(s Ressource) {

	if Merchant.Inventory[s.Name] > 0 {
		char.AddRessource(s)
		char.CheckGold(s.Price)
		char.Gold -= s.Price
		println("achat de 1 ", s.Name)
	} else {
		println("je n'ai plus de ", s.Name, " frérot !")
	}
}
