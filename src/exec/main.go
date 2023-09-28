package main

import menu "ProjetRed/Menu"

func main() {

	//Don't forget to enable the Presentation cinematic in "principal.go": line 310!!!
	menu.Run()

	//Only to test the game (it skips the loading screen)
	//menu.CreateDisplayVariables()
	//menu.PrincipalMenu()
}
