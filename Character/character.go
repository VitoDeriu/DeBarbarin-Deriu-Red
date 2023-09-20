package character

import (
	"fmt"
	"strconv"
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
	Elf.strength = 5
	Elf.endurance = 7
	Elf.perception = 12

	Human.name = "Humain"
	Human.hpMax = 100
	Human.mpMax = 100
	Human.strength = 7
	Human.endurance = 7
	Human.perception = 10

	Dwarf.name = "Nain"
	Dwarf.hpMax = 120
	Dwarf.mpMax = 80
	Dwarf.strength = 12
	Dwarf.endurance = 7
	Dwarf.perception = 5

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
	MainChar.strength = races[selectedRace-1].strength
	MainChar.endurance = races[selectedRace-1].endurance
	MainChar.perception = races[selectedRace-1].perception
	MainChar.skills = append(MainChar.skills, races[selectedRace-1].innateSkill)
	MainChar.inventory = map[string]int{"épée d'entrainement": 1, "potion de vie": 1}
	MainChar.gold = 10
	return MainChar
}

func Spaces(str string, totalLength int) string {
	var spaces string
	for i := 0; i < totalLength-len([]rune(str)); i++ {
		spaces += " "
	}
	return spaces
}

func (char Character) DisplayInfo() {

	nameSpaces := Spaces(char.name, 17)
	raceSpaces := Spaces(char.class.name, -1)
	levelSpaces := Spaces(strconv.Itoa(char.level), 1)
	hpSpaces := Spaces(strconv.Itoa(char.hp)+"/"+strconv.Itoa(char.hpMax), 3)

	fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	fmt.Println("  ▓                                                                            ▓")
	fmt.Println("  ▓      nom :      ", char.name, nameSpaces, "                                       ▓")
	fmt.Println("  ▓                                                                            ▓")
	fmt.Println("  ▓      race :     ", char.class.name, raceSpaces, "                                                    ▓")
	fmt.Println("  ▓                                                                            ▓")
	fmt.Println("  ▓      niveau :   ", char.level, levelSpaces, "                                                       ▓")
	fmt.Println("  ▓                                                                            ▓")
	fmt.Println("  ▓      HP :       ", char.hp, "/", char.hpMax, hpSpaces, "                                                 ▓")
	fmt.Println("  ▓                                                                            ▓")
	fmt.Println("  ▓      Inventaire :                                                          ▓")
	i := 0
	for index, nb := range char.inventory {
		i++
		itemSpaces1 := Spaces(strconv.Itoa(nb), 0)
		itemSpaces2 := Spaces(index, 41)
		fmt.Println("  ▓                   ", nb, itemSpaces1, "     ", index, itemSpaces2, "    ▓")
		if i == 4 {
			break
		}
	}
	for j := 4 - i; j > 0; j-- {
		fmt.Println("  ▓                                                                            ▓")
	}
	if i == 4 {
		fmt.Println("  ▓      Il n'y a pas assez de place pour afficher tout votre inventaire !     ▓")
	} else {
		fmt.Println("  ▓                                                                            ▓")
	}
	fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	time.Sleep(time.Second * 10)
}
