package character

// initialisation de tous les items et leur caractéristiques

/*type Item struct {
	name     string
	typeItem string
	price    int
}*/

// Déclaration des Ressources
var FourrureDeLoup = Ressource{
	Name:  "Fourrure de loup",
	Price: 4,
}
var PeauDeTroll = Ressource{
	Name:  "Peau de troll",
	Price: 7,
}
var CuirDeSanglier = Ressource{
	Name:  "Cuir de sanglier",
	Price: 3,
}
var PlumeDeCorbeau = Ressource{
	Name:  "Plume de corbeau",
	Price: 1,
}

var AllRessources = []Ressource{
	FourrureDeLoup,
	PeauDeTroll,
	CuirDeSanglier,
	PlumeDeCorbeau,
}

// Déclaration des Equipements
var TrainingSword = Equipement{
	Name:   "Épée d'entrainement",
	Slot:   "Mains",
	Attack: 5,
}
var ChapeauDelAventurier = Equipement{
	Name:   "Chapeau de l'aventurier",
	Slot:   "Tete",
	Recipe: map[Ressource]int{PlumeDeCorbeau: 1, CuirDeSanglier: 1},
	HpMax:  10,
}
var TuniqueDelAventurier = Equipement{
	Name:   "Tunique de l'aventurier",
	Slot:   "Torse",
	Recipe: map[Ressource]int{FourrureDeLoup: 2, PeauDeTroll: 1},
}
var BotteDelAventurier = Equipement{
	Name:   "Bottes de l'aventurier",
	Slot:   "Pieds",
	Recipe: map[Ressource]int{FourrureDeLoup: 1, CuirDeSanglier: 1},
}

var AllEquipement = []Equipement{
	TrainingSword,
	ChapeauDelAventurier,
	TuniqueDelAventurier,
	BotteDelAventurier,
}

// Déclaration des Potions
var PotionDeSoin = Potion{
	Name:         "Potion de soin",
	StatBuffed:   "Hp",
	Buff:         20,
	EffectOnTime: 0,
	Price:        3,
}
var PotionDePoison = Potion{
	Name:         "Potion de poison",
	StatDebuffed: "Hp",
	Debuff:       10,
	EffectOnTime: 3,
	Price:        6,
}

var AllPotion = []Potion{
	PotionDePoison,
	PotionDeSoin,
}

// Déclarations des SpellBooks
var SpellBookFireBall = SpellBook{
	Name:  "Grimoire: boule de feu",
	Skill: FireBall,
	Price: 25,
}
var SpellBookSwordSlash = SpellBook{
	Name:  "Manuel: coup d'épée",
	Skill: SwordSlash,
	Price: 20,
}
var SpellBookSwordGuard = SpellBook{
	Name:  "Manuel: guarde d'épée",
	Skill: SwordGuard,
	Price: 15,
}

var AllSpellBook = []SpellBook{
	SpellBookFireBall,
	SpellBookSwordSlash,
	SpellBookSwordGuard,
}

// Déclaration des Skills
var FireBall = Skill{
	Name:       "Boule de feu",
	Attack:     20,
	StatBuffed: "Attaque",
	MpCost:     10,
}
var ElfPunch = Skill{
	Name:       "Coup de poing elfe",
	Attack:     10,
	Defense:    0,
	StatBuffed: "Agilité",
	Buff:       5,
	MpCost:     0,
}
var HumanPunch = Skill{
	Name:       "Coup de poing humain",
	Attack:     10,
	Defense:    0,
	StatBuffed: "Attaque",
	Buff:       5,
	MpCost:     0,
}
var DwarfPunch = Skill{
	Name:       "Coup de poing nain",
	Attack:     10,
	Defense:    0,
	StatBuffed: "Défense",
	Buff:       5,
	MpCost:     0,
}
var SlimeShot = Skill{
	Name:       "Coup de Slime",
	Attack:     2,
	StatBuffed: "Attaque",
	MpCost:     0,
}
var SwordSlash = Skill{
	Name:       "Coup d'épée",
	Attack:     15,
	StatBuffed: "Attaque",
	MpCost:     0,
}
var SwordGuard = Skill{
	Name:       "Blocage à l'épée",
	Defense:    20,
	StatBuffed: "Défense",
	MpCost:     0,
}

var Skills = []Skill{
	FireBall,
	ElfPunch,
	HumanPunch,
	DwarfPunch,
	SlimeShot,
	SwordSlash,
	SwordGuard,
}
