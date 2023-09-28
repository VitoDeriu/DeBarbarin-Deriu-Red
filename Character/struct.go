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

type Enemy struct {
	Name       string
	Race       string
	Level      int
	Hp         int
	HpMax      int
	Mp         int
	MpMax      int
	Attack     int
	Defense    int
	Agility    int
	Skills     []Skill
	Loot       map[string]int
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
	Name    string
	Attack  int
	Defense int
	Kind    string
	MpCost  int
}

// struct des items

type Potion struct {
	Name         string
	StatBuffed   string
	StatDebuffed string
	Buff         int
	Debuff       int
	EffectOnTime int //durée de l'effet en seconde
	Price        int
}

type Equipement struct {
	Name    string
	Slot    string // "head" , "chest", "arms", "legs", "feet"
	Recipe  map[Ressource]int
	HpMax   int
	MpMax   int
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

func (myChar *Character) UseSpecialItem() {
	myChar.Hp = myChar.HpMax
	myChar.Mp = myChar.MpMax
	myChar.Attack += 2
	myChar.Defense += 2
	myChar.Xp += 5
}

// fonction de mort et respawn. a faire checker a chaque dégat pris
func (char *Character) IsDead(enemy Enemy) (bool, int) {
	var xpLost int
	if char.Hp <= 0 {
		char.Hp = char.HpMax / 2
		if enemy.Level-char.Level < 4 {
			xpLost = enemy.Level * (enemy.HpMax / 4)
			char.Xp -= xpLost
		}
		return true, xpLost
	}
	return false, xpLost
}
