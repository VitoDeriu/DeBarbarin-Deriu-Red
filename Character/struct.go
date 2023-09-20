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
	attack     int
	defense    int
	agility    int
	skills     []Skill
	inventory  map[string]int
	gold       int
}

type Race struct {
	name        string
	hpMax       int
	mpMax       int
	attack      int
	defense     int
	agility     int
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

// struct des items

type Potion struct {
	name         string
	gainHp       int
	lossHp       int
	effectOnTime int //durée de l'effet en seconde
	price        int
}

type Equipement struct {
	name       string
	shield     int
	perception int
	force      int
	mpMax      int
	price      int
}

type SpellBook struct {
	name  string
	skill Skill
	price int
}
