package character

func DisplayForge() {
	println("Salam mon reuf ! voici ce que je pourrais te fabriquer moyenant moula et ressource :")

}

func (char Character) CheckRecipe(recipe map[Ressource]int) bool {
	for key, val := range recipe {
		if char.Inventory[key.Name] < val {
			return false
		}
	}
	return true
}

func (char *Character) CraftChapeauDelAventurier() {
	if char.CheckGold(5) && char.CheckRecipe(ChapeauDelAventurier.Recipe) {
		char.Gold -= 5
		char.RemoveRecipeRessources(ChapeauDelAventurier.Recipe)
		char.AddEquipement(ChapeauDelAventurier)
	}

}

func (char *Character) RemoveRecipeRessources(recipe map[Ressource]int) {
	for key, val := range recipe {
		for i := 0; i < val; i++ {
			char.RemoveRessource(key)
		}
	}
}
