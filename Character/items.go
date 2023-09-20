package character // initialisation de tous les items

/*type Item struct {
	name     string
	typeItem string
	price    int
}*/

var PotionDeSoin,
	PotionDePoison Item

var allitems []Item

func CreateItems() {
	PotionDeSoin.name = "Potion de soin"
	PotionDeSoin.typeItem = "Potion"
	PotionDeSoin.price = 0

	PotionDePoison.name = "Potion de poison"
	PotionDePoison.typeItem = "Potion"
	PotionDePoison.price = 30

	allitems = append(allitems, PotionDeSoin)
	allitems = append(allitems, PotionDePoison)

}
