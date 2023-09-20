package character // initialisation de tous les items

/*type Item struct {
	name     string
	typeItem string
	price    int
}*/

var PotionDeSoin,
	PotionDePoison,
	SpellBookFireBall Item

var allitems []Item

func CreateItems() {
	PotionDeSoin.name = "Potion de soin"
	PotionDeSoin.typeItem.potion.gainHp = 20
	PotionDeSoin.price = 0

	PotionDePoison.name = "Potion de poison"
	PotionDePoison.typeItem.potion.lossHp = 10
	PotionDePoison.typeItem.potion.effectOnTime = 3
	PotionDePoison.price = 30

	SpellBookFireBall.name = "Livre de sort : Boule de feu"
	SpellBookFireBall.typeItem.spellBook.skill = 
	SpellBookFireBall.price = 100

	allitems = append(allitems, PotionDeSoin)
	allitems = append(allitems, PotionDePoison)

}
