package character

import (
	"time"
)

var Elf Race
var Human Race
var Dwarf Race
var races []Race

var elfPunch Skill
var humanPunch Skill
var dwarfPunch Skill
var skills []Skill

func CreateRaces() {
	Elf.name = "Elfe"
	Elf.hpMax = 80
	Elf.mpMax = 120
	Elf.attack = 5
	Elf.defense = 7
	Elf.agility = 12

	Human.name = "Humain"
	Human.hpMax = 100
	Human.mpMax = 100
	Human.attack = 7
	Human.defense = 7
	Human.agility = 10

	Dwarf.name = "Nain"
	Dwarf.hpMax = 120
	Dwarf.mpMax = 80
	Dwarf.attack = 12
	Dwarf.defense = 7
	Dwarf.agility = 5

	races = append(races, Human)
	races = append(races, Elf)
	races = append(races, Dwarf)
}

func CreateSkills() {
	elfPunch.name = "Coup de poing elfe"
	elfPunch.attack = 10
	elfPunch.defense = 0
	elfPunch.statBuffed = "perception"
	elfPunch.buff = 5
	elfPunch.mpCost = 0

	humanPunch.name = "Coup de poing humain"
	humanPunch.attack = 10
	humanPunch.defense = 0
	humanPunch.statBuffed = "strength"
	humanPunch.buff = 5
	humanPunch.mpCost = 0

	dwarfPunch.name = "Coup de poing nain"
	dwarfPunch.attack = 10
	dwarfPunch.defense = 0
	dwarfPunch.statBuffed = "defense"
	dwarfPunch.buff = 5
	dwarfPunch.mpCost = 0

	skills = append(skills, humanPunch)
	skills = append(skills, elfPunch)
	skills = append(skills, dwarfPunch)
}

func CreateMainCharacter(name string, selectedRace int) Character {
	CreateRaces()
	CreateSkills()
	var MainChar Character
	MainChar.name = name
	MainChar.class = races[selectedRace-1]
	MainChar.level = 1
	MainChar.xp = 0
	MainChar.statPoints = 1
	MainChar.hp = races[selectedRace-1].hpMax / 2
	MainChar.hpMax = races[selectedRace-1].hpMax
	MainChar.mp = races[selectedRace-1].mpMax / 2
	MainChar.mpMax = races[selectedRace-1].mpMax
	MainChar.attack = races[selectedRace-1].attack
	MainChar.defense = races[selectedRace-1].defense
	MainChar.agility = races[selectedRace-1].agility
	MainChar.skills = append(MainChar.skills, races[selectedRace-1].innateSkill)
	MainChar.inventory = map[string]int{"épée d'entrainement": 1, "potion de vie": 1}
	MainChar.gold = 10
	return MainChar
}

// fonction de mort et respawn. a faire checker a chaque dégat prit
func (char *Character) Dead() {
	if char.hp <= 0 {
		println("Vous avez succombé a vos blessures sale noob")
		time.Sleep(1 * time.Second)
		println("Aller, go respawn !")
		char.hp = char.hpMax / 2
		println("PV : ", char.hp)
	}
	return
}
