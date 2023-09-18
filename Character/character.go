package character

import (
	"fmt"
	"strconv"
	"time"
)

type Character struct {
	name      string
	class     Race
	level     int
	hp        int
	hpMax     int
	inventory map[string]int
}

type Race struct {
	name  string
	hpMax int
}

var Elf Race
var Human Race
var Dwarf Race
var races []Race

func CreateRaces() {
	Elf.name = "Elfe"
	Elf.hpMax = 80

	Human.name = "Humain"
	Human.hpMax = 100

	Dwarf.name = "Nain"
	Dwarf.hpMax = 120

	races = append(races, Human)
	races = append(races, Elf)
	races = append(races, Dwarf)
}

func CreateMainCharacter(name string, selectedRace int) Character {
	CreateRaces()
	var MainChar Character
	MainChar.name = name
	MainChar.class = races[selectedRace-1]
	MainChar.level = 1
	MainChar.hp = races[selectedRace-1].hpMax / 2
	MainChar.hpMax = races[selectedRace-1].hpMax
	MainChar.inventory = map[string]int{"épée d'entrainement": 1, "potion de vie": 1}
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
