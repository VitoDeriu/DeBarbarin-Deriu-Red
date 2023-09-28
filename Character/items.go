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

// Treasure item you seldom win when winning against an enemy (each race has its own, for consistency)
var WorldTreeLeaf = Ressource{
	Name: "Feuille d'Yggdrasil",
}
var MountainHeart = Ressource{
	Name: "Coeur de montagne",
}
var GraalFragment = Ressource{
	Name: "Fragment du Graal",
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
	Slot:   "arms",
	Attack: 5,
	Price:  15,
}
var ChapeauDelAventurier = Equipement{
	Name:   "Chapeau de l'aventurier",
	Slot:   "head",
	HpMax:  10,
	Price:  12,
	Recipe: map[Ressource]int{PlumeDeCorbeau: 1, CuirDeSanglier: 1},
}
var TuniqueDelAventurier = Equipement{
	Name:    "Tunique de l'aventurier",
	Slot:    "chest",
	Defense: 3,
	Price:   7,
	Recipe:  map[Ressource]int{FourrureDeLoup: 2, PeauDeTroll: 1},
}
var BotteDelAventurier = Equipement{
	Name:    "Bottes de l'aventurier",
	Slot:    "feet",
	Agility: 3,
	Price:   10,
	Recipe:  map[Ressource]int{FourrureDeLoup: 1, CuirDeSanglier: 1},
}
var ElfHat = Equipement{
	Name:  "Chapeau elfe",
	Slot:  "head",
	MpMax: 10,
	Price: 25,
}
var ElfSword = Equipement{
	Name:   "Épée elfe",
	Slot:   "arms",
	Attack: 10,
	MpMax:  20,
	Price:  35,
}
var ElfTunic = Equipement{
	Name:    "Tunique elfe",
	Slot:    "chest",
	Defense: 10,
	MpMax:   10,
	HpMax:   5,
	Price:   35,
}
var ElfBoots = Equipement{
	Name:    "Bottes elfes",
	Slot:    "feet",
	Defense: 5,
	Agility: 10,
	Price:   30,
}
var SteelSword = Equipement{
	Name:   "Épée en acier",
	Slot:   "arms",
	Attack: 25,
	Price:  350,
}

var AllEquipement = []Equipement{
	TrainingSword,
	ChapeauDelAventurier,
	TuniqueDelAventurier,
	BotteDelAventurier,
	ElfHat,
	ElfSword,
	ElfTunic,
	ElfBoots,
}

// Déclaration des Potions
var PotionDeSoin = Potion{
	Name:         "Potion de soin",
	StatBuffed:   "Hp",
	Buff:         20,
	EffectOnTime: 0,
	Price:        3,
}
var GrandePotionDeSoin = Potion{
	Name:         "Potion de soin +",
	StatBuffed:   "Hp",
	Buff:         40,
	EffectOnTime: 0,
	Price:        8,
}
var PotionDeMana = Potion{
	Name:         "Potion de mana",
	StatBuffed:   "Mp",
	Buff:         30,
	EffectOnTime: 0,
	Price:        5,
}
var GrandePotionDeMana = Potion{
	Name:         "Potion de mana +",
	StatBuffed:   "Mp",
	Buff:         50,
	EffectOnTime: 0,
	Price:        10,
}
var PotionDePoison = Potion{
	Name:         "Potion de poison",
	StatDebuffed: "Hp",
	Debuff:       10,
	EffectOnTime: 3,
	Price:        6,
}
var EnhanceInventory = Potion{
	Name:  "Inventaire ++",
	Price: 150,
}

var AllPotion = []Potion{
	PotionDePoison,
	PotionDeSoin,
	GrandePotionDeSoin,
	PotionDeMana,
	GrandePotionDeMana,
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
	Name:   "Boule de feu",
	Attack: 20,
	Kind:   "Attaque",
	MpCost: 8,
}
var ElfPunch = Skill{
	Name:    "Coup de poing elfe",
	Attack:  16,
	Defense: 0,
	Kind:    "Attaque",
	MpCost:  2,
}
var HumanPunch = Skill{
	Name:    "Coup de poing humain",
	Attack:  12,
	Defense: 0,
	Kind:    "Attaque",
	MpCost:  0,
}
var DwarfPunch = Skill{
	Name:    "Coup de poing nain",
	Attack:  10,
	Defense: 2,
	Kind:    "Attaque",
	MpCost:  0,
}
var SlimeShot = Skill{
	Name:   "Coup de Slime",
	Attack: 25,
	Kind:   "Attaque",
	MpCost: 0,
}
var OrcPunch = Skill{
	Name:   "Coup de poing orque",
	Attack: 30,
	Kind:   "Attaque",
	MpCost: 0,
}
var SwordSlash = Skill{
	Name:   "Coup d'épée",
	Attack: 15,
	Kind:   "Attaque",
	MpCost: 0,
}
var SwordGuard = Skill{
	Name:    "Blocage à l'épée",
	Defense: 20,
	Kind:    "Défense",
	MpCost:  0,
}

var Skills = []Skill{
	FireBall,
	ElfPunch,
	HumanPunch,
	DwarfPunch,
	SlimeShot,
	SwordSlash,
	SwordGuard,
	OrcPunch,
}
