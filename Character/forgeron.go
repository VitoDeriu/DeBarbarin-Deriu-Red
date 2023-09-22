package character

var BlacksmithEquipments []Equipement

func InitBlacksmith() {
	BlacksmithEquipments = append(BlacksmithEquipments, ChapeauDelAventurier)
	BlacksmithEquipments = append(BlacksmithEquipments, TuniqueDelAventurier)
	BlacksmithEquipments = append(BlacksmithEquipments, BotteDelAventurier)
}

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

func (char *Character) RemoveRecipeRessources(recipe map[Ressource]int) {
	for key, val := range recipe {
		for i := 0; i < val; i++ {
			char.RemoveRessource(key)
		}
	}
}

func (char *Character) CraftChapeauDelAventurier() {
	if char.CheckGold(5) && char.CheckRecipe(ChapeauDelAventurier.Recipe) && char.FullInventory() {
		char.Gold -= 5
		char.RemoveRecipeRessources(ChapeauDelAventurier.Recipe)
		char.AddEquipement(ChapeauDelAventurier)
	}
}
func (char *Character) CraftTuniqueDelAventurier() {
	if char.CheckGold(5) && char.CheckRecipe(TuniqueDelAventurier.Recipe) && char.FullInventory() {
		char.Gold -= 5
		char.RemoveRecipeRessources(TuniqueDelAventurier.Recipe)
		char.AddEquipement(TuniqueDelAventurier)
	}
}

func (char *Character) CraftBotteDelAventurier() {
	if char.CheckGold(5) && char.CheckRecipe(BotteDelAventurier.Recipe) && char.FullInventory() {
		char.Gold -= 5
		char.RemoveRecipeRessources(BotteDelAventurier.Recipe)
		char.AddEquipement(BotteDelAventurier)
	}
}
