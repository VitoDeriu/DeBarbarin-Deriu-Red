package character

import "fmt"

func (char *Character) TakePotion() { // fonction d'utilisation de la potion, a modifier en focntion générale pour tous les items suivant leur typeItem
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

// fonctions pour l'ajout et le retrait d'item dans l'inventaire. Peut servir pour le marchant

func (char *Character) addInventory(s Item) { // pour ajouter un item a l'inventaire quand on achete ou qu'on loot
	for key := range char.inventory {
		if key == s.name {
			char.inventory[key]++ //si l'item est déjà présent dans l'inventaire on ajoute 1
			return
		}
	}
	char.inventory[s.name] = 1 // si l'item n'est pas dans l'inventaire on le crée
	fmt.Println("ajout de 1 ", s.name, ".")
}

func (char *Character) removeInventory(s Item) { //pour supprimer un item de l'inventaire
	for key := range char.inventory {
		if key == s.name {
			char.inventory[key]--
			if char.inventory[key] <= 0 {
				delete(char.inventory, key) // si non on delete l'item de l'inventaire
			}
			fmt.Println("retrait de 1", s.name, ".")
			return
		}
	}
	fmt.Println("Impossible de retirer un objet que tu ne possède pas") //
}
