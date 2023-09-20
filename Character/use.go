package character

import (
	"fmt"
	"time"
)

// fonction d'utilisation de la potion, a modifier en focntion générale pour tous les items suivant leur typeItem

func (char *Character) TakePotionSoin() {

	if char.Inventory["Potion de Poison"] != 0 {
		if char.Hp == char.HpMax {
			fmt.Println("Je n'ai pas soif !")
			return
		} else if char.HpMax-char.Hp < 20 {
			char.Hp = char.HpMax
		} else {
			char.Hp += PotionDeSoin.GainHp
		}
		char.Inventory["Potion de vie"] -= 1
		if char.Inventory["Potion de vie"] == 0 {
			delete(char.Inventory, "Potion de soin")
		}
		fmt.Println("+20 hp")
		fmt.Println(char.Hp, "/", char.HpMax)
		return
	} else {
		fmt.Println(char.Name, " ➵  Je n'ai pas de potion !")
	}
}

//fonction pour la popo de poison

func (char *Character) poisonPot() {

	if char.Inventory["Potion de Poison"] != 0 { //pas besoin d'aller check dans tout l'inventaire avec un for range, si y'a pas l'item ca renvoi au else
		println("\"C'EST DEGUEU !! Argh !\" Vous perdez 30hp...")
		for i := 0; i < PotionDePoison.EffectOnTime; i++ { // boucle qui prend en parametre le effectOnTime de item.potion pour appliquer l'effet plusieur fois en fonction du temps
			char.Hp -= PotionDePoison.LossHp
			char.Dead()
			time.Sleep(1 * time.Second)
		}
		println("Il vous reste ", char.Hp, " HP")
		println("Pas très malin tout ca...")
		char.RemovePotion(PotionDePoison) // appel fonction remove pour consommer la potion
	} else {
		println("je n'ai pas de Potion de Poison")
	}
}

func (char *Character) SpellBook(s SpellBook) {

	for _, sk := range char.Skills {
		if sk == s.Skill {
			return
		}
	}
	char.Skills = append(char.Skills, s.Skill)
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

func (char *Character) FullInventory() bool {
	if len(char.Inventory) >= 10 {
		println("inventaire plein.. deso")
		return false
	}
	return true
}

// Add et Remove In	ventory pour les Potions
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
