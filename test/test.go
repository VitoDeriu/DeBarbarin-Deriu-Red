package main

import (
	char "ProjetRed/Character"
)

func main() {
	char.InitialiseMerchant()
	println("Bijour beaucoup, voici quoi que j'ai a te vendre :")
	char.DisplayBoutique() //inventaire du marchant
	//char.Merchant.BuyItem()
}