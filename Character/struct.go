package character

//Struct de la création de perso

type Character struct {
	name       string
	class      Race
	level      int
	xp         int
	statPoints int
	hp         int
	hpMax      int
	mp         int
	mpMax      int
	strength   int
	endurance  int
	perception int
	skills     []Skill
	inventory  map[string]int
	gold       int
}

type Race struct {
	name        string
	hpMax       int
	mpMax       int
	strength    int
	endurance   int
	perception  int
	innateSkill Skill
}

type Skill struct {
	name       string
	attack     int
	defense    int
	statBuffed string
	buff       int
	mpCost     int
}

// Struct du Marchant

type Item struct {
	name     string
	typeItem string
	price    int
}
