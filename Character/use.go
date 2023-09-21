package character

import (
	"fmt"
	"time"
)

// fonction d'utilisation de la potion, a modifier en fonction générale pour tous les items suivant leur typeItem
// du coup ne pas modifier ca sert a rien

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

//fonction pour l'utilisation de la popo de poison

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

// fonction pour l'apprentissage des sort via les spellbooks
func (char *Character) SpellBook(s SpellBook) {

	for _, sk := range char.Skills {
		if sk == s.Skill {
			return
		}
	}
	char.Skills = append(char.Skills, s.Skill)
	char.RemoveSpellBook(s)
}
