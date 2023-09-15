package character

import "fmt"

type Character struct {
	name      string
	class     string
	level     int
	hpMax     int
	hp        int
	inventory map[string]int
}

func (char *Character) CreateCharacter(name string, class string, level int, hpMax int, hp int, inventory map[string]int) {
	char.name = name
	char.class = class
	char.level = level
	char.hpMax = hpMax
	char.hp = hp
	char.inventory = inventory
}

func (char Character) DisplayInfo() {
	fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	fmt.Println("▓\t", "nom : ", char.name, "   \t\t\t\t\t\t▓")
	fmt.Println("▓\t", "classe : ", char.class, "   \t\t\t\t\t\t▓")
	fmt.Println("▓\t", "niveau : ", char.level, "   \t\t\t\t\t\t▓")
	fmt.Println("▓\t", "HP : ", char.hp, "/", char.hpMax, "   \t\t\t\t▓")
	fmt.Println("▓\t Inventaire : \t\t\t\t\t\t\t▓")
	for index, nb := range char.inventory {
		if index == "Hache des ténèbres" {
			fmt.Printf("▓\t\t%d   %v\t\t\t\t▓\n", nb, index)
			continue
		}
		fmt.Printf("▓\t\t%d   %v\t\t\t\t\t▓\n", nb, index)
	}
	fmt.Println("₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
}
