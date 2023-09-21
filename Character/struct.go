package character

//Struct de la création de perso

type Character struct {
	Name       string
	Class      Race
	Level      int
	Xp         int
	StatPoints int
	Hp         int
	HpMax      int
	Mp         int
	MpMax      int
	Attack     int
	Defense    int
	Agility    int
	Skills     []Skill
	Inventory  map[string]int
	Gold       int
	Equipement []Equipement
}

type Race struct {
	Name        string
	HpMax       int
	MpMax       int
	Attack      int
	Defense     int
	Agility     int
	InnateSkill Skill
}

type Skill struct {
	Name       string
	Attack     int
	Defense    int
	StatBuffed string
	Buff       int
	MpCost     int
}

// struct des items

type Potion struct {
	Name         string
	GainHp       int
	LossHp       int
	EffectOnTime int //durée de l'effet en seconde
	Price        int
}

type Equipement struct {
	Name    string
	Slot    string // "Tete" , "Torse", "Mains", "Jambe", "Pieds"
	Recipe  map[Ressource]int
	Defense int
	Agility int
	Attack  int
	Price   int
}

type SpellBook struct {
	Name  string
	Skill Skill
	Price int
}

type Ressource struct {
	Name  string
	Price int
}
