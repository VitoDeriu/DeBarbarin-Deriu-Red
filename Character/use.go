package character

import (
	"fmt"
	"time"
)

func (char *Character) TakePotionSoin() { // fonction d'utilisation de la potion, a modifier en focntion générale pour tous les items suivant leur typeItem

	if char.inventory["Potion de Poison"] != 0 {
		if char.hp == char.hpMax {
			fmt.Println("Je n'ai pas soif !")
			return
		} else if char.hpMax-char.hp < 20 {
			char.hp = char.hpMax
		} else {
			char.hp += PotionDeSoin.gainHp
		}
		char.inventory["Potion de vie"] -= 1
		if char.inventory["Potion de vie"] == 0 {
			delete(char.inventory, "Potion de soin")
		}
		fmt.Println("+20 hp")
		fmt.Println(char.hp, "/", char.hpMax)
		return
	} else {
		fmt.Println(char.name, " ➵  Je n'ai pas de potion !")
	}
}

//fonction pour la popo de poison

func (char *Character) poisonPot() {

	if char.inventory["Potion de Poison"] != 0 { //pas besoin d'aller check dans tout l'inventaire avec un for range, si y'a pas l'item ca renvoi au else
		println("\"C'EST DEGUEU !! Argh !\" Vous perdez 30hp...")
		for i := 0; i < PotionDePoison.effectOnTime; i++ { // boucle qui prend en parametre le effectOnTime de item.potion pour appliquer l'effet plusieur fois en fonction du temps
			char.hp -= PotionDePoison.lossHp
			char.Dead()
			time.Sleep(1 * time.Second)
		}
		println("Il vous reste ", char.hp, " HP")
		println("Pas très malin tout ca...")
		char.removePotion(PotionDePoison) // appel fonction remove pour consommer la potion
	} else {
		println("je n'ai pas de Potion de Poison")
	}

}

// fonctions pour l'ajout et le retrait d'item dans l'inventaire. Peut servir pour le marchant

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

// Add et Remove Inventory pour les Potions
func (char *Character) addPotion(s Potion) {
	for key := range char.inventory {
		if key == s.name {
			char.inventory[key]++
			return
		}
	}
	char.inventory[s.name] = 1
	fmt.Println("ajout de 1 ", s.name, ".")
}

func (char *Character) removePotion(s Potion) {
	for key := range char.inventory {
		if key == s.name {
			char.inventory[key]--
			if char.inventory[key] <= 0 {
				delete(char.inventory, key)
			}
			fmt.Println("retrait de 1", s.name, ".")
			return
		}
	}
	fmt.Println("Impossible de retirer un objet que tu ne possède pas")
}

// Add et Remove pour les SpellBooks
func (char *Character) addSpellBook(s SpellBook) {
	for key := range char.inventory {
		if key == s.name {
			char.inventory[key]++
			return
		}
	}
	char.inventory[s.name] = 1
	fmt.Println("ajout de 1 ", s.name, ".")
}

func (char *Character) removeSpellBook(s SpellBook) {
	for key := range char.inventory {
		if key == s.name {
			char.inventory[key]--
			if char.inventory[key] <= 0 {
				delete(char.inventory, key)
			}
			fmt.Println("retrait de 1", s.name, ".")
			return
		}
	}
	fmt.Println("Impossible de retirer un objet que tu ne possède pas")
}
