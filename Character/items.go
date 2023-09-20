package character // initialisation de tous les items

/*type Item struct {
	name     string
	typeItem string
	price    int
}*/

var PotionDeSoin Potion
var PotionDePoison Potion
var SpellBookFireBall SpellBook
var fireBall Skill

var allPotion []Potion
var allEquipement []Equipement
var allSpellBook []SpellBook

func CreateItems() {

	//création des potions
	PotionDeSoin.name = "Potion de soin"
	PotionDeSoin.gainHp = 20
	PotionDeSoin.price = 0

	PotionDePoison.name = "Potion de poison"
	PotionDePoison.lossHp = 10
	PotionDePoison.effectOnTime = 3
	PotionDePoison.price = 30

	allPotion = append(allPotion, PotionDeSoin)
	allPotion = append(allPotion, PotionDePoison)

	//Créations des SpellBook
	SpellBookFireBall.name = "Livre de sort : Boule de feu"
	SpellBookFireBall.skill = fireBall
	SpellBookFireBall.price = 100

	allSpellBook = append(allSpellBook, SpellBookFireBall)

}

func CreateSkills() {
	elfPunch.name = "Coup de poing elfe"
	elfPunch.attack = 10
	elfPunch.defense = 0
	elfPunch.statBuffed = "perception"
	elfPunch.buff = 5
	elfPunch.mpCost = 0

	humanPunch.name = "Coup de poing humain"
	humanPunch.attack = 10
	humanPunch.defense = 0
	humanPunch.statBuffed = "strength"
	humanPunch.buff = 5
	humanPunch.mpCost = 0

	dwarfPunch.name = "Coup de poing nain"
	dwarfPunch.attack = 10
	dwarfPunch.defense = 0
	dwarfPunch.statBuffed = "defense"
	dwarfPunch.buff = 5
	dwarfPunch.mpCost = 0

	fireBall.name = "Boule de feu"
	fireBall.attack = 20
	fireBall.mpCost = 10

	skills = append(skills, humanPunch)
	skills = append(skills, elfPunch)
	skills = append(skills, dwarfPunch)
	skills = append(skills, fireBall)

}
