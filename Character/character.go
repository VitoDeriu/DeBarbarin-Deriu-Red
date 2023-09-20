package character

import (
	"time"
)

var Elf Race
var Human Race
var Dwarf Race
var Races []Race

var ElfPunch Skill
var HumanPunch Skill
var DwarfPunch Skill
var Skills []Skill

func CreateRaces() {
	Elf.Name = "Elfe"
	Elf.HpMax = 80
	Elf.MpMax = 120
	Elf.Attack = 5
	Elf.Defense = 7
	Elf.Agility = 12
	Elf.InnateSkill = ElfPunch

	Human.Name = "Humain"
	Human.HpMax = 100
	Human.MpMax = 100
	Human.Attack = 7
	Human.Defense = 7
	Human.Agility = 10
	Human.InnateSkill = HumanPunch

	Dwarf.Name = "Nain"
	Dwarf.HpMax = 120
	Dwarf.MpMax = 80
	Dwarf.Attack = 12
	Dwarf.Defense = 7
	Dwarf.Agility = 5
	Dwarf.InnateSkill = DwarfPunch

	Races = append(Races, Human)
	Races = append(Races, Elf)
	Races = append(Races, Dwarf)
}

func CreateMainCharacter(name string, selectedRace int) Character {
	CreateSkills()
	CreateRaces()
	var MainChar Character
	MainChar.Name = name
	MainChar.Class = Races[selectedRace-1] // la bendo la bendo
	MainChar.Level = 1
	MainChar.Xp = 0
	MainChar.StatPoints = 1
	MainChar.Hp = Races[selectedRace-1].HpMax / 2
	MainChar.HpMax = Races[selectedRace-1].HpMax
	MainChar.Mp = Races[selectedRace-1].MpMax / 2
	MainChar.MpMax = Races[selectedRace-1].MpMax
	MainChar.Attack = Races[selectedRace-1].Attack
	MainChar.Defense = Races[selectedRace-1].Defense
	MainChar.Agility = Races[selectedRace-1].Agility
	MainChar.Skills = append(MainChar.Skills, Races[selectedRace-1].InnateSkill)
	MainChar.Inventory = map[string]int{"épée d'entrainement": 1, "potion de vie": 1}
	MainChar.Gold = 10
	return MainChar
}

// fonction de mort et respawn. a faire checker a chaque dégat pris
func (char *Character) Dead() {
	if char.Hp <= 0 {
		println("Vous avez succombé a vos blessures sale noob")
		time.Sleep(1 * time.Second)
		println("Aller, go respawn !")
		char.Hp = char.HpMax / 2
		println("PV : ", char.Hp)
	}
}
