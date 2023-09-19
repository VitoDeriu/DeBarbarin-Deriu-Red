package main

import "fmt"

type Character struct {
	name           string
	class          string
	level          int
	maxHealthPoint int
	HealtPoint     int
	inventory      map[string]int
} // création de la classe (struture en golang) charactere

func (p *Character) InitCharacter(name string, class string) {
	p.name = name
	p.class = class
	p.level = 1
	p.maxHealthPoint = 100
	p.HealtPoint = 40
	p.inventory = map[string]int{"Potion de soin": 3, "Epée de folie": 1}
} // fonction d'initialisation du personnage avec du stuff, il faut rentrer le nom et la classe dans la fonction main

func (p *Character) DisplayInfo() {
	fmt.Println("Nom : ", p.name)
	fmt.Println("Classe : ", p.class)
	fmt.Println("Lvl : ", p.level)
	fmt.Println("MaxHP : ", p.maxHealthPoint)
	fmt.Println("HP : ", p.HealtPoint)
} // fonction de display des infos du perso

func (p *Character) accessInventory() {
	fmt.Println("Inventaire : ")
	for index, nb := range p.inventory {
		fmt.Println(nb, index)
	}
} // fonction a ajouter au menu pour acceder a l'inventaire, ca affiche une liste de l'inventaire avec la clé (le nom de l'item) et la valeur (le nombre d'item)

func (p *Character) TakePot() {
	var havePotion bool
	for index := range p.inventory {
		if index == "Potion de soin" {
			havePotion = true
		}
		if havePotion {
			if p.HealtPoint == p.maxHealthPoint {
				println("ga pa souaf")
				return
			}
			if (p.maxHealthPoint - p.HealtPoint) >= 50 {
				p.HealtPoint += 50
				println("glouglou (+ 50 hp)")
				p.inventory["Potion de soin"] -= 1
				println("pv : ", p.HealtPoint, "/", p.maxHealthPoint)
			}
			if (p.maxHealthPoint - p.HealtPoint) < 50 {
				println("glouglou (+ ", (p.maxHealthPoint - p.HealtPoint), " hp)")
				p.HealtPoint = p.maxHealthPoint
				p.inventory["Potion de soin"] -= 1
				println("pv : ", p.HealtPoint, "/", p.maxHealthPoint)
			}
		}
	}
	println("ga pa potion")
	havePotion = false
}

func main() {
	var p1 Character

	p1.InitCharacter("xXx-UwU-xXx", "K1KouL0L")
	p1.DisplayInfo()
	p1.accessInventory()
	p1.TakePot()
	p1.TakePot()
	p1.TakePot()
}
