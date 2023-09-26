package character

import (
	"fmt"
)

// Fichier pour les fonctions de gestion d'inventaire, Add et Remove ainsi que les différents check pour les achats et vente.

// fonctions génériques pour l'ajout et le retrait d'item dans l'inventaire. ne pas décommenter, elles sont fausses en plus

//func (char *Character) addInventory(s Item) { // pour ajouter un item a l'inventaire quand on achete ou qu'on loot
//	for key := range char.inventory {
//		if key == s.name {
//			char.inventory[key]++ //si l'item est déjà présent dans l'inventaire on ajoute 1
//			return
//		}
//	}
//	char.inventory[s.name] = 1 // si l'item n'est pas dans l'inventaire on le crée
//	fmt.Println("ajout de 1 ", s.name, ".")
//}

//func (char *Character) removeInventory(s Item) { //pour supprimer un item de l'inventaire
//	for key := range char.inventory {
//		if key == s.name {
//			char.inventory[key]--
//			if char.inventory[key] <= 0 {
//				delete(char.inventory, key) // si non on delete l'item de l'inventaire
//			}
//			fmt.Println("retrait de 1", s.name, ".")
//			return
//		}
//	}
//	fmt.Println("Impossible de retirer un objet que tu ne possède pas")
//}

// pour checker si on a assez d'argent dans l'inventaire, a utiliser a chaque fonction Buy...
func (char *Character) CheckGold(price int) bool {
	if price > char.Gold {
		println("t'as pas assez de moneyyy")
		return false
	}
	return true
}

// check si l'inventaire est plein a chaque fonction Add
var InvMax int = 10

func (char *Character) FullInventory() bool {
	if len(char.Inventory) >= InvMax {
		println("inventaire plein.. deso")
		return false
	}
	return true
}

// Augmentation de la taille de l'inventaire
func (char *Character) UpgradeInventory() {
	char.CheckGold(30)
	if InvMax < 40 {
		InvMax += 10
		char.Gold -= 30
	} else {
		println("L'inventaire est déjà maxé")
	}
}

// Add et Remove Inventory pour les Potions
func (char *Character) AddPotion(s Potion) {
	if char.FullInventory() {
		for key := range char.Inventory {
			if key == s.Name {
				char.Inventory[key]++
				return
			}
		}
		char.Inventory[s.Name] = 1
		fmt.Println("ajout de 1 ", s.Name, ".")
	}
}

func (char *Character) RemovePotion(s Potion) {

	for key := range char.Inventory {
		if key == s.Name {
			char.Inventory[key]--
			if char.Inventory[key] <= 0 {
				delete(char.Inventory, key)
			}
			fmt.Println("retrait de 1", s.Name, ".")
			return
		}
	}
	fmt.Println("Impossible de retirer un objet que tu ne possède pas")
}

// Add et Remove pour les SpellBooks
func (char *Character) AddSpellBook(s SpellBook) {
	if char.FullInventory() {
		for key := range char.Inventory {
			if key == s.Name {
				char.Inventory[key]++
				return
			}
		}
		char.Inventory[s.Name] = 1
		fmt.Println("ajout de 1 ", s.Name, ".")
	}
}

func (char *Character) RemoveSpellBook(s SpellBook) {
	for key := range char.Inventory {
		if key == s.Name {
			char.Inventory[key]--
			if char.Inventory[key] <= 0 {
				delete(char.Inventory, key)
			}
			fmt.Println("retrait de 1", s.Name, ".")
			return
		}
	}
	fmt.Println("Impossible de retirer un objet que tu ne possède pas")
}

// Add et Remove pour les Equipements
func (char *Character) AddEquipement(s Equipement) {
	if char.FullInventory() {
		for key := range char.Inventory {
			if key == s.Name {
				char.Inventory[key]++
				return
			}
		}
		char.Inventory[s.Name] = 1
		fmt.Println("ajout de 1 ", s.Name, ".")
	}
}

func (char *Character) RemoveEquipement(s Equipement) {
	for key := range char.Inventory {
		if key == s.Name {
			char.Inventory[key]--
			if char.Inventory[key] <= 0 {
				delete(char.Inventory, key)
			}
			fmt.Println("retrait de 1", s.Name, ".")
			return
		}
	}
	fmt.Println("Impossible de retirer un objet que tu ne possède pas")
}

// Add et Remove pour les Ressources
func (char *Character) AddRessource(s Ressource) {
	if char.FullInventory() {
		for key := range char.Inventory {
			if key == s.Name {
				char.Inventory[key]++
				return
			}
		}
		char.Inventory[s.Name] = 1
		fmt.Println("ajout de 1 ", s.Name, ".")
	}
}

func (char *Character) RemoveRessource(s Ressource) {
	for key := range char.Inventory {
		if key == s.Name {
			char.Inventory[key]--
			if char.Inventory[key] <= 0 {
				delete(char.Inventory, key)
			}
			fmt.Println("retrait de 1", s.Name, ".")
			return
		}
	}
	fmt.Println("Impossible de retirer un objet que tu ne possède pas")
}

func FindPotion(s string) Potion {
	for _, potion := range AllPotion {
		if s == potion.Name {
			return potion
		}
	}
	var nothing Potion
	return nothing
}
