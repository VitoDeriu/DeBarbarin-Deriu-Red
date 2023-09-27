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

func (char *Character) SpellBook(s SpellBook) {
	for _, sk := range char.Skills {
		if sk == s.Skill {
			return
		}
	}
	if len(char.Skills) >= 5 {
		char.Skills = char.Skills[1:]
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

// fonction pour équiper un item
func (char *Character) Equiper(e Equipement) {
	var toRemove Equipement
	var IndRemove int
	var isEmpty = true
	//boucle for range pour verifier si on a pas déjà un item équipé
	for i, val := range char.Equipement {
		if val.Slot == e.Slot {
			isEmpty = false
			toRemove = val
			IndRemove = i
		}
	}
	//si y'a rien d'équipé
	if isEmpty == true {
		char.HpMax += e.HpMax
		char.MpMax += e.MpMax
		char.Defense += e.Defense
		char.Agility += e.Agility
		char.Attack += e.Attack
		char.Equipement = append(char.Equipement, e)
		char.RemoveEquipement(e)
		return
	}
	//si y'a qqch d'équipé
	if isEmpty == false {
		//on retire les stats de l'objet a enlever
		char.HpMax -= toRemove.HpMax
		char.MpMax -= toRemove.MpMax
		char.Defense -= toRemove.Defense
		char.Agility -= toRemove.Agility
		char.Attack -= toRemove.Attack
		//on enlève l'objet a equiper de l'inventaire et on y ajoute celui qu'on retire
		char.RemoveEquipement(e)
		char.AddEquipement(toRemove)
		//on ajoute les stats du nouvel objet
		char.HpMax += e.HpMax
		char.MpMax += e.MpMax
		char.Defense += e.Defense
		char.Agility += e.Agility
		char.Attack += e.Attack
		//on ajoute le nouvel objet a la slice d'équipement
		char.Equipement = append(char.Equipement, e)
		//on retire l'ancien objet de la slice char.equipement
		char.Equipement = append(char.Equipement[:IndRemove], char.Equipement[IndRemove+1:]...)
		return
	}
}
