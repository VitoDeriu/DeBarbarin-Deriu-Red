package character

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
	Defense: 5,
	Agility: 7,
	Loot:    map[string]int{"Manuel: coup d'épée": 1},
}

var OrdinaryGuard = Enemy{
	Name:    "Guarde ordinaire",
	Race:    "Guarde",
	Level:   1,
	Skills:  []Skill{SwordSlash, SwordGuard},
	Hp:      50,
	HpMax:   50,
	Attack:  7,
	Defense: 7,
	Agility: 7,
	Loot:    map[string]int{"Manuel: guarde d'épée": 1},
}

var Enemies = []Enemy{
	OrdinarySlime,
	OrdinaryGoblin,
	OrdinaryGuard,
}
