package character

import "math/rand"

var OrdinarySlime = Enemy{
	Name:    "Slime ordinaire",
	Race:    "Slime",
	Level:   1,
	Skills:  []Skill{SlimeShot},
	Hp:      20,
	HpMax:   20,
	Attack:  2,
	Defense: 5,
	Agility: 10,
	Loot:    map[string]int{"Potion de soin": 1},
}

var OrdinaryGoblin = Enemy{
	Name:    "Goblin ordinaire",
	Race:    "Goblin",
	Level:   1,
	Skills:  []Skill{SwordSlash},
	Hp:      40,
	HpMax:   40,
	Attack:  5,
	Defense: 3,
	Agility: 8,
	Loot:    map[string]int{"Manuel: coup d'épée": 1},
}

var OrdinaryGuard = Enemy{
	Name:    "Guarde ordinaire",
	Race:    "Guarde",
	Level:   1,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      50,
	HpMax:   50,
	Attack:  5,
	Defense: 5,
	Agility: 7,
	Loot:    map[string]int{"Manuel: guarde d'épée": 1},
}

var Enemies = []Enemy{
	OrdinarySlime,
	OrdinaryGoblin,
	OrdinaryGuard,
}

func (enemy *Enemy) EnemyDeath(myChar *Character) (bool, bool) {
	if enemy.Hp <= 0 {
		myChar.Xp += enemy.Level * (enemy.HpMax / 4)
		if i := rand.Intn(10); i >= 5 && enemy.Loot != nil {
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
