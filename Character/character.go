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
	Elf.Name = "Elfe"
	Elf.HpMax = 80
	Elf.MpMax = 120
	Elf.Attack = 5
	Elf.Defense = 7
	Elf.Agility = 12

	Human.Name = "Humain"
	Human.HpMax = 100
	Human.MpMax = 100
	Human.Attack = 7
	Human.Defense = 7
	Human.Agility = 10

	Dwarf.Name = "Nain"
	Dwarf.HpMax = 120
	Dwarf.MpMax = 80
	Dwarf.Attack = 12
	Dwarf.Defense = 7
	Dwarf.Agility = 5

	races = append(races, Human)
	races = append(races, Elf)
	races = append(races, Dwarf)
}

func CreateMainCharacter(name string, selectedRace int) Character {
	CreateRaces()
	CreateSkills()
	var MainChar Character
	MainChar.Name = name
	MainChar.Class = races[selectedRace-1] // la bendo la bendo
	MainChar.Level = 1
	MainChar.Xp = 0
	MainChar.StatPoints = 1
	MainChar.Hp = races[selectedRace-1].HpMax / 2
	MainChar.HpMax = races[selectedRace-1].HpMax
	MainChar.Mp = races[selectedRace-1].MpMax / 2
	MainChar.MpMax = races[selectedRace-1].MpMax
	MainChar.Attack = races[selectedRace-1].Attack
	MainChar.Defense = races[selectedRace-1].Defense
	MainChar.Agility = races[selectedRace-1].Agility
	MainChar.Skills = append(MainChar.Skills, races[selectedRace-1].InnateSkill)
	MainChar.Inventory = map[string]int{"épée d'entrainement": 1, "potion de vie": 1}
	MainChar.Gold = 10
	return MainChar
}

// fonction de mort et respawn. a faire checker a chaque dégat prit
func (char *Character) Dead() {
	if char.Hp <= 0 {
		println("Vous avez succombé a vos blessures sale noob")
		time.Sleep(1 * time.Second)
		println("Aller, go respawn !")
		char.Hp = char.HpMax / 2
		println("PV : ", char.Hp)
	}
	return
}
