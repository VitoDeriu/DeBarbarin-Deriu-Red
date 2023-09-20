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
	PotionDeSoin.Name = "Potion de soin"
	PotionDeSoin.GainHp = 20
	PotionDeSoin.Price = 0

	PotionDePoison.Name = "Potion de poison"
	PotionDePoison.LossHp = 10
	PotionDePoison.EffectOnTime = 3
	PotionDePoison.Price = 30

	allPotion = append(allPotion, PotionDeSoin)
	allPotion = append(allPotion, PotionDePoison)

	//Créations des SpellBook
	SpellBookFireBall.Name = "Livre de sort : Boule de feu"
	SpellBookFireBall.Skill = fireBall
	SpellBookFireBall.Price = 100

	allSpellBook = append(allSpellBook, SpellBookFireBall)

}

func CreateSkills() {
	elfPunch.Name = "Coup de poing elfe"
	elfPunch.Attack = 10
	elfPunch.Defense = 0
	elfPunch.StatBuffed = "perception"
	elfPunch.Buff = 5
	elfPunch.MpCost = 0

	humanPunch.Name = "Coup de poing humain"
	humanPunch.Attack = 10
	humanPunch.Defense = 0
	humanPunch.StatBuffed = "strength"
	humanPunch.Buff = 5
	humanPunch.MpCost = 0

	dwarfPunch.Name = "Coup de poing nain"
	dwarfPunch.Attack = 10
	dwarfPunch.Defense = 0
	dwarfPunch.StatBuffed = "defense"
	dwarfPunch.Buff = 5
	dwarfPunch.MpCost = 0

	fireBall.Name = "Boule de feu"
	fireBall.Attack = 20
	fireBall.MpCost = 10

	skills = append(skills, humanPunch)
	skills = append(skills, elfPunch)
	skills = append(skills, dwarfPunch)
	skills = append(skills, fireBall)

}
