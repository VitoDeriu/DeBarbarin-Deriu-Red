package character

import (
	"fmt"
	"time"
)

var Elf = Race{
	Name:        "Elfe",
	HpMax:       80,
	MpMax:       120,
	Attack:      5,
	Defense:     7,
	Agility:     12,
	InnateSkill: ElfPunch,
}
var Human = Race{
	Name:        "Humain",
	HpMax:       100,
	MpMax:       100,
	Attack:      7,
	Defense:     7,
	Agility:     10,
	InnateSkill: HumanPunch,
}
var Dwarf = Race{
	Name:        "Nain",
	HpMax:       120,
	MpMax:       80,
	Attack:      12,
	Defense:     7,
	Agility:     5,
	InnateSkill: DwarfPunch,
}
var Races = []Race{
	Human,
	Elf,
	Dwarf,
}

func CreateMainCharacter(name string, selectedRace int) Character {
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
	MainChar.Inventory = map[string]int{TrainingSword.Name: 1, PotionDeSoin.Name: 1, SpellBookFireBall.Name: 1, PotionDePoison.Name: 2}
	MainChar.Gold = 100
	return MainChar
}

// fonction de mort et respawn. a faire checker a chaque dégat pris
func (char *Character) Dead() bool {
	if char.Hp <= 0 {
		fmt.Println("Vous avez succombé a vos blessures sale noob")
		time.Sleep(1 * time.Second)
		fmt.Println("Aller, go respawn !")
		char.Hp = char.HpMax / 2
		fmt.Println("HP : ", char.Hp)
		return true
	}
	return false
}

func (char *Character) LevelUp() bool {
	var hasLevelledUp bool
	for char.Xp >= 80*char.Level {
		hasLevelledUp = true
		char.Xp -= 80 * char.Level
		char.Level++
		char.HpMax += 5
		char.Attack += 5
		char.Defense += 5
		char.Agility += 2
		char.Hp = char.HpMax
		char.MpMax += 5
		char.Mp = char.MpMax
	}
	return hasLevelledUp
}

func (myChar *Character) UseSpecialItem() {
	myChar.Hp = myChar.HpMax
	myChar.Mp = myChar.MpMax
	myChar.Attack += 3
	myChar.Defense += 3
	myChar.Agility += 1
	myChar.Xp += 10
}
