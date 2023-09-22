package character

import (
	"fmt"
	"time"
)

// initialisation de tous les items et leur caractéristiques

/*type Item struct {
	name     string
	typeItem string
	price    int
}*/

// Déclaration des Equipements
var TrainingSword Equipement
var ChapeauDelAventurier Equipement
var TuniqueDelAventurier Equipement
var BotteDelAventurier Equipement

var AllEquipement []Equipement

// Déclaration des Potions
var PotionDeSoin Potion
var PotionDePoison Potion

var AllPotion []Potion

// Déclaration des Ressources
var FourrureDeLoup Ressource
var PeauDeTroll Ressource
var CuirDeSanglier Ressource
var PlumeDeCorbeau Ressource

var AllRessources []Ressource

// Déclarations des SpellBooks
var SpellBookFireBall SpellBook

var AllSpellBook []SpellBook

// Déclaration des Skills
var FireBall Skill
var ElfPunch Skill
var HumanPunch Skill
var DwarfPunch Skill

var Skills []Skill

func CreateItems() {

	//création des équipements
	TrainingSword.Name = "Épée d'entrainement"
	TrainingSword.Attack = 5

	ChapeauDelAventurier.Name = "Chapeau de l'aventurier"
	ChapeauDelAventurier.Slot = "Tete"
	// ChapeauDelAventurier.Recipe[PlumeDeCorbeau] = 1
	// ChapeauDelAventurier.Recipe[CuirDeSanglier] = 1
	ChapeauDelAventurier.Recipe = map[Ressource]int{PlumeDeCorbeau: 1, CuirDeSanglier: 1}
	fmt.Printf("%#v", ChapeauDelAventurier.Recipe)
	time.Sleep(time.Second * 5)
	ChapeauDelAventurier.HpMax = 10

	TuniqueDelAventurier.Name = "Tunique de l'aventurier"
	TuniqueDelAventurier.Slot = "Torse"
	TuniqueDelAventurier.Recipe = map[Ressource]int{FourrureDeLoup: 2, PeauDeTroll: 1}

	BotteDelAventurier.Name = "Bottes de l'aventurier"
	BotteDelAventurier.Slot = "Pieds"
	BotteDelAventurier.Recipe = map[Ressource]int{FourrureDeLoup: 1, CuirDeSanglier: 1}

	AllEquipement = append(AllEquipement, TrainingSword)
	AllEquipement = append(AllEquipement, ChapeauDelAventurier)
	AllEquipement = append(AllEquipement, TuniqueDelAventurier)
	AllEquipement = append(AllEquipement, BotteDelAventurier)

	//création des potions
	PotionDeSoin.Name = "Potion de soin"
	PotionDeSoin.StatBuffed = "Hp"
	PotionDeSoin.Buff = 20
	PotionDeSoin.EffectOnTime = 0
	PotionDeSoin.Price = 3

	PotionDePoison.Name = "Potion de poison"
	PotionDePoison.StatDebuffed = "Hp"
	PotionDePoison.Debuff = 10
	PotionDePoison.EffectOnTime = 3
	PotionDePoison.Price = 6

	AllPotion = append(AllPotion, PotionDeSoin)
	AllPotion = append(AllPotion, PotionDePoison)

	//Création des Ressources
	FourrureDeLoup.Name = "Fourrure de loup"
	FourrureDeLoup.Price = 4

	PeauDeTroll.Name = "Peau de troll"
	PeauDeTroll.Price = 7

	CuirDeSanglier.Name = "Cuir de sanglier"
	CuirDeSanglier.Price = 3

	PlumeDeCorbeau.Name = "Plume de corbeau"
	PlumeDeCorbeau.Price = 1

	AllRessources = append(AllRessources, FourrureDeLoup)
	AllRessources = append(AllRessources, PeauDeTroll)
	AllRessources = append(AllRessources, CuirDeSanglier)
	AllRessources = append(AllRessources, PlumeDeCorbeau)

	//Créations des SpellBook
	SpellBookFireBall.Name = "Grimoire boule de feu"
	SpellBookFireBall.Skill = FireBall
	SpellBookFireBall.Price = 25

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
