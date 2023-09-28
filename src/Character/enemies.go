package character

import (
	"math/rand"
)

var Trainer1 = Enemy{
	Name:    "Hector",
	Race:    "Entraineur",
	Level:   1,
	Skills:  []Skill{HumanPunch},
	Hp:      50,
	HpMax:   50,
	Attack:  6,
	Defense: 6,
	Agility: 7,
	Loot:    map[string]int{SpellBookSwordSlash.Name: 1, TrainingSword.Name: 1},
}

var Trainer2 = Enemy{
	Name:    "Hector",
	Race:    "Entraineur",
	Level:   2,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      60,
	HpMax:   60,
	Attack:  17,
	Defense: 17,
	Agility: 9,
	Loot:    map[string]int{PotionDeSoin.Name: 1},
}

var Trainer3 = Enemy{
	Name:    "Hector",
	Race:    "Entraineur",
	Level:   3,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      70,
	HpMax:   70,
	Attack:  26,
	Defense: 26,
	Agility: 11,
	Loot:    map[string]int{GrandePotionDeSoin.Name: 1},
}

var Trainers = []Enemy{
	Trainer1,
	Trainer2,
	Trainer3,
}

var OrdinaryGuard1 = Enemy{
	Name:    "Guarde ordinaire",
	Race:    "Guarde",
	Level:   4,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      80,
	HpMax:   80,
	Attack:  35,
	Defense: 35,
	Agility: 13,
	Loot:    map[string]int{GrandePotionDeSoin.Name: 1},
}

var OrdinaryGuard2 = Enemy{
	Name:    "Guarde ordinaire",
	Race:    "Guarde",
	Level:   5,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      95,
	HpMax:   95,
	Attack:  44,
	Defense: 44,
	Agility: 15,
	Loot:    map[string]int{GrandePotionDeSoin.Name: 1},
}

var OrdinaryGuard3 = Enemy{
	Name:    "Guarde ordinaire",
	Race:    "Guarde",
	Level:   6,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      110,
	HpMax:   110,
	Attack:  53,
	Defense: 53,
	Agility: 17,
	Loot:    map[string]int{GrandePotionDeSoin.Name: 1},
}

var EnemiesTournament = []Enemy{
	OrdinaryGuard1,
	OrdinaryGuard2,
	OrdinaryGuard3,
}

var Adventurer = Enemy{
	Name:    "Aventurier",
	Race:    "Humain",
	Level:   7,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      125,
	HpMax:   125,
	Attack:  62,
	Defense: 62,
	Agility: 19,
	Gold:    10,
	Loot:    map[string]int{GrandePotionDeSoin.Name: 1, GrandePotionDeMana.Name: 1, BotteDelAventurier.Name: 1, TuniqueDelAventurier.Name: 1, ChapeauDelAventurier.Name: 1},
}

var Mercenary = Enemy{
	Name:    "Mercenaire",
	Race:    "Elfe",
	Level:   8,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      140,
	HpMax:   140,
	Attack:  71,
	Defense: 71,
	Agility: 21,
	Gold:    12,
	Loot:    map[string]int{GrandePotionDeSoin.Name: 1, ElfSword.Name: 1, ElfBoots.Name: 1, ElfTunic.Name: 1, ElfBoots.Name: 1},
}

var EnemiesArena = []Enemy{
	Adventurer,
	Mercenary,
}

var OrdinarySlime = Enemy{
	Name:    "Slime",
	Race:    "Slime",
	Level:   10,
	Skills:  []Skill{SlimeShot},
	Hp:      200,
	HpMax:   200,
	Attack:  85,
	Defense: 95,
	Agility: 30,
	Gold:    2,
	Loot:    map[string]int{BotteDelAventurier.Name: 1},
}

var OrdinaryGoblin = Enemy{
	Name:    "Goblin ordinaire",
	Race:    "Goblin",
	Level:   12,
	Skills:  []Skill{SwordSlash},
	Hp:      230,
	HpMax:   230,
	Attack:  100,
	Defense: 85,
	Agility: 21,
	Gold:    15,
	Loot:    map[string]int{SteelSword.Name: 1},
}

var OrdinaryOrc = Enemy{
	Name:    "Orque ordinaire",
	Race:    "Orque",
	Level:   14,
	Skills:  []Skill{OrcPunch},
	Hp:      260,
	HpMax:   260,
	Attack:  110,
	Defense: 130,
	Agility: 15,
	Gold:    15,
	Loot:    map[string]int{SteelSword.Name: 1},
}

var EnemiesMission = []Enemy{
	OrdinaryGoblin,
	OrdinarySlime,
	OrdinaryOrc,
}

func (enemy *Enemy) EnemyDeath(myChar *Character) (bool, bool) {
	if enemy.Hp <= 0 {
		myChar.Xp += enemy.Level * (enemy.HpMax / 4)
		if i := rand.Intn(80); i <= 40 && enemy.Loot != nil {

			//How to win the treasure item of the race!
			if i < 8 {
				switch myChar.Class {
				case Elf:
					myChar.AddRessource(WorldTreeLeaf)
				case Dwarf:
					myChar.AddRessource(MountainHeart)
				case Human:
					myChar.AddRessource(GraalFragment)
				default:
				}
			}
			var loot string
			for item := range enemy.Loot {
				loot = item
				break
			}
			if myChar.Inventory[loot] == 0 && myChar.FullInventory() {
				myChar.Inventory[loot] = 1
			} else {
				myChar.Inventory[loot] += 1
			}
			myChar.Gold += (i / 10) * enemy.Gold
		}
		return true, myChar.LevelUp()
	}
	return false, false
}
