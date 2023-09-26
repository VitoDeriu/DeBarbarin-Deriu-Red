package character

import (
	"fmt"
	"time"
)

// Fonction permettant d'utiliser n'importe quelle potion avec ses buffs et ses debuffs
func (char *Character) TakePotion(p Potion) {

	if char.Inventory[p.Name] != 0 {

		//Vérification de chaque stat qui peut être boostée
		switch p.StatBuffed {

		case "Hp":
			if char.Hp == char.HpMax {
				fmt.Println("Je n'ai pas soif !")
				return
			} else if char.HpMax-char.Hp < p.Buff {
				char.Hp = char.HpMax
			} else {
				if p.EffectOnTime > 0 {
					for i := 0; i < p.EffectOnTime; i++ {
						char.Hp += p.Buff
						time.Sleep(time.Second * 1)
					}
				} else {
					char.Hp += p.Buff
				}
			}
			fmt.Println("+", p.Buff, " ", p.StatBuffed)
			fmt.Println(char.Hp, "/", char.HpMax)

		case "Mp":
			if char.Mp == char.MpMax {
				fmt.Println("Je n'ai pas soif !")
				return
			} else if char.MpMax-char.Mp < p.Buff {
				char.Mp = char.MpMax
			} else {
				if p.EffectOnTime > 0 {
					for i := 0; i < p.EffectOnTime; i++ {
						char.Mp += p.Buff
						time.Sleep(time.Second * 1)
					}
				} else {
					char.Mp += p.Buff
				}
			}
			fmt.Println("+", p.Buff, " ", p.StatBuffed)
			fmt.Println(char.Mp, "/", char.MpMax)

		case "Attack":
			if p.EffectOnTime > 0 {
				for i := 0; i < p.EffectOnTime; i++ {
					char.Attack += p.Buff
					time.Sleep(time.Second * 1)
				}
			} else {
				char.Attack += p.Buff
			}
			fmt.Println("+", p.Buff, " ", p.StatBuffed)
			fmt.Println(char.Attack)

		case "Defense":
			if p.EffectOnTime > 0 {
				for i := 0; i < p.EffectOnTime; i++ {
					char.Defense += p.Buff
					time.Sleep(time.Second * 1)
				}
			} else {
				char.Defense += p.Buff
			}
			fmt.Println("+", p.Buff, " ", p.StatBuffed)
			fmt.Println(char.Defense)

		case "Agility":
			if p.EffectOnTime > 0 {
				for i := 0; i < p.EffectOnTime; i++ {
					char.Agility += p.Buff
					time.Sleep(time.Second * 1)
				}
			} else {
				char.Agility += p.Buff
			}
			fmt.Println("+", p.Buff, " ", p.StatBuffed)
			fmt.Println(char.Agility)
		}

		//Vérification de chaque stat qui peut être handicapée
		switch p.StatDebuffed {

		case "Hp":
			if char.Hp <= p.Debuff {
				char.Hp = 0
				char.Dead()
			} else {
				if p.EffectOnTime > 0 {
					for i := 0; i < p.EffectOnTime; i++ {
						char.Hp -= p.Debuff
						char.Dead()
						time.Sleep(time.Second * 1)
					}
				} else {
					char.Hp -= p.Debuff
				}
			}
			fmt.Println("\"C'EST DEGUEU !! Argh !\" Vous perdez ", p.StatDebuffed, " ", p.Debuff, "...")
			fmt.Println("-", p.Debuff, " ", p.StatDebuffed)
			fmt.Println("Il vous reste ", char.Hp, "/", char.HpMax)
			fmt.Println("Pas très malin tout ca...")

		case "Mp":
			if char.Mp <= p.Debuff {
				char.Mp = 0
			} else {
				if p.EffectOnTime > 0 {
					for i := 0; i < p.EffectOnTime; i++ {
						char.Mp -= p.Debuff
						time.Sleep(time.Second * 1)
					}
				} else {
					char.Mp -= p.Debuff
				}
				if char.Mp < 0 {
					char.Mp = 0
				}
			}
			fmt.Println("-", p.Debuff, " ", p.StatDebuffed)
			fmt.Println(char.Mp, "/", char.MpMax)

		case "Attack":
			if p.EffectOnTime > 0 {
				for i := 0; i < p.EffectOnTime; i++ {
					char.Attack -= p.Debuff
					time.Sleep(time.Second * 1)
				}
			} else {
				char.Attack -= p.Debuff
			}
			if char.Attack < 0 {
				char.Attack = 0
			}
			fmt.Println("-", p.Debuff, " ", p.StatDebuffed)
			fmt.Println(char.Attack)

		case "Defense":
			if p.EffectOnTime > 0 {
				for i := 0; i < p.EffectOnTime; i++ {
					char.Defense -= p.Debuff
					time.Sleep(time.Second * 1)
				}
			} else {
				char.Defense -= p.Debuff
			}
			if char.Defense < 0 {
				char.Defense = 0
			}
			fmt.Println("-", p.Debuff, " ", p.StatDebuffed)
			fmt.Println(char.Defense)

		case "Agility":
			if p.EffectOnTime > 0 {
				for i := 0; i < p.EffectOnTime; i++ {
					char.Agility -= p.Debuff
					time.Sleep(time.Second * 1)
				}
			} else {
				char.Agility -= p.Debuff
			}
			if char.Agility < 0 {
				char.Agility = 0
			}
			fmt.Println("-", p.Debuff, " ", p.StatDebuffed)
			fmt.Println(char.Agility)
		}
		char.RemovePotion(p)
	}
}

// fonction d'utilisation de la potion, a modifier en fonction générale pour tous les items suivant leur typeItem
// du coup ne pas modifier ca sert a rien

// func (char *Character) TakePotionSoin() {
//
//		if char.Inventory["Potion de soin"] != 0 {
//			if char.Hp == char.HpMax {
//				fmt.Println("Je n'ai pas soif !")
//				return
//			} else if char.HpMax-char.Hp < 20 {
//				char.Hp = char.HpMax
//			} else {
//				char.Hp += PotionDeSoin.GainHp
//			}
//			char.Inventory["Potion de soin"] -= 1
//			if char.Inventory["Potion de soin"] == 0 {
//				delete(char.Inventory, "Potion de soin")
//			}
//			fmt.Println("+20 hp")
//			fmt.Println(char.Hp, "/", char.HpMax)
//			return
//		} else {
//			fmt.Println(char.Name, " ➵  Je n'ai pas de potion !")
//		}
//	}
//
// //fonction pour l'utilisation de la popo de poison
//
// func (char *Character) PoisonPot() {
//
//		if char.Inventory["Potion de Poison"] != 0 { //pas besoin d'aller check dans tout l'inventaire avec un for range, si y'a pas l'item ca renvoi au else
//			println("\"C'EST DEGUEU !! Argh !\" Vous perdez 30hp...")
//			for i := 0; i < PotionDePoison.EffectOnTime; i++ { // boucle qui prend en parametre le effectOnTime de item.potion pour appliquer l'effet plusieur fois en fonction du temps
//				char.Hp -= PotionDePoison.LossHp
//				char.Dead()
//				time.Sleep(1 * time.Second)
//			}
//			println("Il vous reste ", char.Hp, " HP")
//			println("Pas très malin tout ca...")
//			char.RemovePotion(PotionDePoison) // appel fonction remove pour consommer la potion
//		} else {
//			println("je n'ai pas de Potion de Poison")
//		}
//	}

// Fonction pour l'apprentissage des sort via les spellbooks
func (char *Character) SpellBook(s SpellBook) {
	for _, sk := range char.Skills {
		if sk == s.Skill {
			return
		}
	}
	char.Skills = append(char.Skills, s.Skill)
	char.RemoveSpellBook(s)
}

func (myChar *Character) UseSkill(skill Skill, enemy *Enemy, buffDefense int) (bool, int, int) {
	if myChar.Mp < skill.MpCost {
		return false, 0, 0
	}
	if skill.Kind == "Attaque" {
		var damage int
		myChar.Mp -= skill.MpCost
		damage = enemy.Attack + skill.Attack
		if myChar.Defense+buffDefense >= damage {
			damage /= 2
		} else {
			damage -= (myChar.Defense + buffDefense) / 10
		}
		enemy.Hp -= damage
		return true, skill.Defense, damage
	} else {
		myChar.Mp -= skill.MpCost
		return true, skill.Defense, 0
	}
}

func (enemy *Enemy) UseSkill(skill Skill, myChar *Character, buffDefense int) (int, int) {
	if enemy.Mp < skill.MpCost {
		return 0, 0
	}
	if skill.Kind == "Attaque" {
		var damage int
		enemy.Mp -= skill.MpCost
		damage = enemy.Attack + skill.Attack
		if myChar.Defense+buffDefense >= damage {
			damage /= 2
		} else {
			damage -= (myChar.Defense + buffDefense) / 10
		}
		myChar.Hp -= damage
		return skill.Defense, damage
	} else {
		enemy.Mp -= skill.MpCost
		return skill.Defense, 0
	}
}
