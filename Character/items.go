package character // initialisation de tous les items

/*type Item struct {
	name     string
	typeItem string
	price    int
}*/

var PotionDeSoin Potion
var PotionDePoison Potion
var SpellBookFireBall SpellBook
var FireBall Skill

var AllPotion []Potion
var AllEquipement []Equipement
var AllSpellBook []SpellBook

func CreateItems() {

	//création des potions
	PotionDeSoin.Name = "Potion de soin"
	PotionDeSoin.GainHp = 20
	PotionDeSoin.Price = 0

	PotionDePoison.Name = "Potion de poison"
	PotionDePoison.LossHp = 10
	PotionDePoison.EffectOnTime = 3
	PotionDePoison.Price = 30

	AllPotion = append(AllPotion, PotionDeSoin)
	AllPotion = append(AllPotion, PotionDePoison)

	//Créations des SpellBook
	SpellBookFireBall.Name = "Livre de sort : Boule de feu"
	SpellBookFireBall.Skill = FireBall
	SpellBookFireBall.Price = 100

	AllSpellBook = append(AllSpellBook, SpellBookFireBall)

}

func CreateSkills() {
	ElfPunch.Name = "Coup de poing elfe"
	ElfPunch.Attack = 10
	ElfPunch.Defense = 0
	ElfPunch.StatBuffed = "Agilité"
	ElfPunch.Buff = 5
	ElfPunch.MpCost = 0

	HumanPunch.Name = "Coup de poing humain"
	HumanPunch.Attack = 10
	HumanPunch.Defense = 0
	HumanPunch.StatBuffed = "Attaque"
	HumanPunch.Buff = 5
	HumanPunch.MpCost = 0

	DwarfPunch.Name = "Coup de poing nain"
	DwarfPunch.Attack = 10
	DwarfPunch.Defense = 0
	DwarfPunch.StatBuffed = "Défense"
	DwarfPunch.Buff = 5
	DwarfPunch.MpCost = 0

	FireBall.Name = "Boule de feu"
	FireBall.Attack = 20
	FireBall.MpCost = 10

	Skills = append(Skills, HumanPunch)
	Skills = append(Skills, ElfPunch)
	Skills = append(Skills, DwarfPunch)
	Skills = append(Skills, FireBall)

}
