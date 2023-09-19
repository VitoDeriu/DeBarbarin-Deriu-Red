package menu

import (
	"fmt"
	"time"
)

var MenuLateralBar,
	LogoLittle,
	Sword,
	SoldierArt,
	PrincipalMenuOptions,
	CharCreationMenuText,
	ElfDescription,
	HumanDescription,
	DwarfDescription,
	CharMenuText [][]rune

var BottomBar,
	CharMenuTitleBar,
	PrincipalMenuTitleBar,
	CharCreationMenuTitleBar,
	CharCreationName,
	CharCreationNameError,
	SimpleMenuCursor []rune

func CreateDisplayVariables() {

	MenuLateralBar = append(MenuLateralBar, []rune(" /┃\\ "))
	MenuLateralBar = append(MenuLateralBar, []rune("//┃\\\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\\\┃//"))
	MenuLateralBar = append(MenuLateralBar, []rune("/\\┃/\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\/┃\\/"))
	MenuLateralBar = append(MenuLateralBar, []rune("//┃\\\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\\\┃//"))
	MenuLateralBar = append(MenuLateralBar, []rune("/\\┃/\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\/┃\\/"))
	MenuLateralBar = append(MenuLateralBar, []rune("//┃\\\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\\\┃//"))
	MenuLateralBar = append(MenuLateralBar, []rune("/\\┃/\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\/┃\\/"))
	MenuLateralBar = append(MenuLateralBar, []rune("//┃\\\\"))
	MenuLateralBar = append(MenuLateralBar, []rune("\\\\┃//"))
	MenuLateralBar = append(MenuLateralBar, []rune(" \\┃/ "))

	LogoLittle = append(LogoLittle, []rune("    /"))
	LogoLittle = append(LogoLittle, []rune("\\  /\\"))
	LogoLittle = append(LogoLittle, []rune(" \\/¯¯\\"))

	Sword = append(Sword, []rune("    /("))
	Sword = append(Sword, []rune("O\\\\\\{}============-"))
	Sword = append(Sword, []rune("    \\("))

	SoldierArt = append(SoldierArt, []rune("   .-."))
	SoldierArt = append(SoldierArt, []rune(" __|=|__"))
	SoldierArt = append(SoldierArt, []rune("(_/`-`\\_)"))
	SoldierArt = append(SoldierArt, []rune("//\\___/\\\\"))
	SoldierArt = append(SoldierArt, []rune("<>/   \\<>"))
	SoldierArt = append(SoldierArt, []rune(" \\|_._|/"))
	SoldierArt = append(SoldierArt, []rune("  <_I_>"))
	SoldierArt = append(SoldierArt, []rune("   |||"))
	SoldierArt = append(SoldierArt, []rune("  /_|_\\"))

	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Nouvelle partie"))
	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Écran de chargement"))
	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Bonus"))
	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Crédits"))
	PrincipalMenuOptions = append(PrincipalMenuOptions, []rune("Quitter"))

	CharCreationMenuText = append(CharCreationMenuText, []rune("Choix de la race :"))
	CharCreationMenuText = append(CharCreationMenuText, []rune("Humain"))
	CharCreationMenuText = append(CharCreationMenuText, []rune("Elfe"))
	CharCreationMenuText = append(CharCreationMenuText, []rune("Nain"))
	CharCreationMenuText = append(CharCreationMenuText, []rune("Description :"))

	ElfDescription = append(ElfDescription, []rune("Les elfes se spécialisent dans la magie"))
	ElfDescription = append(ElfDescription, []rune("et la perception : 80 HP et 120 Mana"))

	HumanDescription = append(HumanDescription, []rune("Les humains sont équilibrés dans leurs statistiques"))
	HumanDescription = append(HumanDescription, []rune("et leurs compétences : 100 HP et 100 Mana"))

	DwarfDescription = append(DwarfDescription, []rune("Les nains se spécialisent dans l'endurance"))
	DwarfDescription = append(DwarfDescription, []rune("et leur constitution: 120 HP et 80 Mana"))

	CharMenuText = append(CharMenuText, []rune("Nom :"))
	CharMenuText = append(CharMenuText, []rune("Race :"))
	CharMenuText = append(CharMenuText, []rune("Niveau :"))
	CharMenuText = append(CharMenuText, []rune("Expérience :"))
	CharMenuText = append(CharMenuText, []rune("HP :"))
	CharMenuText = append(CharMenuText, []rune("MP :"))
	CharMenuText = append(CharMenuText, []rune("Force :"))
	CharMenuText = append(CharMenuText, []rune("Endurance :"))
	CharMenuText = append(CharMenuText, []rune("Perception :"))
	CharMenuText = append(CharMenuText, []rune("Prestige :"))
	CharMenuText = append(CharMenuText, []rune("Pièces d'or :"))
	CharMenuText = append(CharMenuText, []rune("Compétences :"))
	CharMenuText = append(CharMenuText, []rune("Équipement :"))
	CharMenuText = append(CharMenuText, []rune("Points stat"))
	CharMenuText = append(CharMenuText, []rune("Inventaire"))
	CharMenuText = append(CharMenuText, []rune("Équipement"))
	CharMenuText = append(CharMenuText, []rune("Retour"))

	BottomBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	PrincipalMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharCreationMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharCreationName = []rune("Nom du personnage (max 20 caractères) :")

	CharCreationNameError = []rune("Erreur : veuillez rentrer un nom valide !")

	SimpleMenuCursor = append(SimpleMenuCursor, '➢')

}

func LoadingScreen() {
	ClearTerminal()
	// fmt.Println("                                                                               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("          PROJET  RED                                    /ito Deriu            ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("                                                     \\  /\\                     ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("                                                      \\/¯¯\\ntoine de Barbarin  ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("                             -|             |-                                 ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("         -|                  [-_-_-_-_-_-_-_-]                  |-             ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("         [-_-_-_-_-]          |             |          [-_-_-_-_-]             ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("          | o   o |           [  0   0   0  ]           | o   o |              ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |     |    -|       |           |       |-    |     |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |     |_-___-___-___-|         |-___-___-___-_|     |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |  o  ]              [    0    ]              [  o  |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |     ]   o   o   o  [ _______ ]  o   o   o   [     | ----__________")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("_____----- |     ]              [ ||||||| ]              [     |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("           |     ]              [ ||||||| ]              [     |               ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("       _-_-|_____]--------------[_|||||||_]--------------[_____|-_-_           ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println("      ( (__________------------_____________-------------_________) )          ")
	time.Sleep(time.Millisecond * 70)
	fmt.Println()
	fmt.Println("Chargement :")
	loadingBar()
	time.Sleep(time.Second * 1)
}

func loadingBar() {
	var loadingBar string
	var loadingPercentage float64
	for i := 0; i < 74; i++ {
		time.Sleep(time.Millisecond * 40)
		loadingBar += "₪"
		loadingPercentage += 1.36
		fmt.Print("\r", int(loadingPercentage), "% ", loadingBar)
	}
}

// func displayMenuBars() {
// 	line := 0
// 	for index := range MenuLateralBar
// 	DisplayText(0, line, string(MenuLateralBar[]))
// }

func principalMenuDisplay(pointingAt int) {

	switch pointingAt {
	case 1:
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//          /(                                                            \\\\┃//")
		fmt.Println("/\\┃/\\      O\\\\\\{}============-      Nouvelle partie                          /\\┃/\\")
		fmt.Println("\\/┃\\/          \\(                                                            \\/┃\\/")
		fmt.Println("//┃\\\\                               Écran de chargement                      //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Bonus                                    /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                        \\/┃\\/")
		fmt.Println("//┃\\\\                               Crédits                                  //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Quitter                                  /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                    /   \\/┃\\/")
		fmt.Println("//┃\\\\                                                                \\  /\\   //┃\\\\")
		fmt.Println("\\\\┃//                                                                 \\/¯¯\\  \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	case 2:
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Nouvelle partie                          /\\┃/\\")
		fmt.Println("\\/┃\\/          /(                                                            \\/┃\\/")
		fmt.Println("//┃\\\\      O\\\\\\{}============-      Écran de chargement                      //┃\\\\")
		fmt.Println("\\\\┃//          \\(                                                            \\\\┃//")
		fmt.Println("/\\┃/\\                               Bonus                                    /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                        \\/┃\\/")
		fmt.Println("//┃\\\\                               Crédits                                  //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Quitter                                  /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                    /   \\/┃\\/")
		fmt.Println("//┃\\\\                                                                \\  /\\   //┃\\\\")
		fmt.Println("\\\\┃//                                                                 \\/¯¯\\  \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	case 3:
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Nouvelle partie                          /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                        \\/┃\\/")
		fmt.Println("//┃\\\\                               Écran de chargement                      //┃\\\\")
		fmt.Println("\\\\┃//          /(                                                            \\\\┃//")
		fmt.Println("/\\┃/\\      O\\\\\\{}============-      Bonus                                    /\\┃/\\")
		fmt.Println("\\/┃\\/          \\(                                                            \\/┃\\/")
		fmt.Println("//┃\\\\                               Crédits                                  //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Quitter                                  /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                    /   \\/┃\\/")
		fmt.Println("//┃\\\\                                                                \\  /\\   //┃\\\\")
		fmt.Println("\\\\┃//                                                                 \\/¯¯\\  \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	case 4:
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Nouvelle partie                          /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                        \\/┃\\/")
		fmt.Println("//┃\\\\                               Écran de chargement                      //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Bonus                                    /\\┃/\\")
		fmt.Println("\\/┃\\/          /(                                                            \\/┃\\/")
		fmt.Println("//┃\\\\      O\\\\\\{}============-      Crédits                                  //┃\\\\")
		fmt.Println("\\\\┃//          \\(                                                            \\\\┃//")
		fmt.Println("/\\┃/\\                               Quitter                                  /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                    /   \\/┃\\/")
		fmt.Println("//┃\\\\                                                                \\  /\\   //┃\\\\")
		fmt.Println("\\\\┃//                                                                 \\/¯¯\\  \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	case 5:
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Nouvelle partie                          /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                        \\/┃\\/")
		fmt.Println("//┃\\\\                               Écran de chargement                      //┃\\\\")
		fmt.Println("\\\\┃//                                                                        \\\\┃//")
		fmt.Println("/\\┃/\\                               Bonus                                    /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                        \\/┃\\/")
		fmt.Println("//┃\\\\                               Crédits                                  //┃\\\\")
		fmt.Println("\\\\┃//          /(                                                            \\\\┃//")
		fmt.Println("/\\┃/\\      O\\\\\\{}============-      Quitter                                  /\\┃/\\")
		fmt.Println("\\/┃\\/          \\(                                                        /   \\/┃\\/")
		fmt.Println("//┃\\\\                                                                \\  /\\   //┃\\\\")
		fmt.Println("\\\\┃//                                                                 \\/¯¯\\  \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	}
}

func displayCredits() {
	ClearTerminal()
	fmt.Println("   ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Crédits ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	fmt.Println(" /┃\\                                                                       /┃\\  ")
	fmt.Println("//┃\\\\                      /                              PROJET RED      //┃\\\\")
	fmt.Println("\\\\┃//                     / ITO                                           \\\\┃//")
	fmt.Println("/\\┃/\\             \\      /\\                           SEPTEMBRE 2023      /\\┃/\\")
	fmt.Println("\\/┃\\/              \\    /  \\                                              \\/┃\\/")
	fmt.Println("//┃\\\\               \\  /¯¯¯¯\\ NTOINE                                      //┃\\\\")
	fmt.Println("\\\\┃//                \\/      \\                                            \\\\┃//")
	fmt.Println("/\\┃/\\                                                                     /\\┃/\\")
	fmt.Println("\\/┃\\/                                                                     \\/┃\\/")
	fmt.Println("//┃\\\\  |¯¯\\  |¯¯\\  /¯¯\\  |¯¯\\  |    |  /¯¯\\ ¯¯¯¯|¯¯¯¯ ¯¯|¯¯ /¯¯\\  |\\   |  //┃\\\\")
	fmt.Println("\\\\┃//  |__/  |__/ /    \\ |   \\ |    | /         |       |  /    \\ | \\  |  \\\\┃//")
	fmt.Println("/\\┃/\\  |     |\\   \\    / |   / |    | \\         |       |  \\    / |  \\ |  /\\┃/\\")
	fmt.Println("\\/┃\\/ _|_    | \\   \\__/  |__/  |____|  \\__/     |     __|__ \\__/  |   \\|  \\/┃\\/")
	fmt.Println("//┃\\\\                                                                     //┃\\\\")
	fmt.Println("\\\\┃//                                                                     \\\\┃//")
	fmt.Println(" \\┃/                                                                  Ⓒ Ⓡ  \\┃/  ")
	fmt.Printf("   ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	time.Sleep(time.Second * 3)
}

func CharacterCreationDisplayName(err bool) {
	if !err {
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//   Nom du personnage (max 20 caractères) :             .-.              \\\\┃//")
		fmt.Println("/\\┃/\\                                                     __|=|__            /\\┃/\\")
		fmt.Println("\\/┃\\/                                                    (_/`-`\\_)           \\/┃\\/")
		fmt.Println("//┃\\\\                                                    //\\___/\\\\           //┃\\\\")
		fmt.Println("\\\\┃//                                                    <>/   \\<>           \\\\┃//")
		fmt.Println("/\\┃/\\                                                     \\|_._|/            /\\┃/\\")
		fmt.Println("\\/┃\\/                                                      <_I_>             \\/┃\\/")
		fmt.Println("//┃\\\\                                                       |||              //┃\\\\")
		fmt.Println("\\\\┃//                                                      /_|_\\             \\\\┃//")
		fmt.Println("/\\┃/\\                                                                        /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                   /    \\/┃\\/")
		fmt.Println("//┃\\\\                                                               \\  /\\    //┃\\\\")
		fmt.Println("\\\\┃//                                                                \\/¯¯\\   \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	} else {

		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\   Erreur : veuillez rentrer un nom valide !                            //┃\\\\")
		fmt.Println("\\\\┃//   Nom du personnage (max 20 caractères) :             .-.              \\\\┃//")
		fmt.Println("/\\┃/\\                                                     __|=|__            /\\┃/\\")
		fmt.Println("\\/┃\\/                                                    (_/`-`\\_)           \\/┃\\/")
		fmt.Println("//┃\\\\                                                    //\\___/\\\\           //┃\\\\")
		fmt.Println("\\\\┃//                                                    <>/   \\<>           \\\\┃//")
		fmt.Println("/\\┃/\\                                                     \\|_._|/            /\\┃/\\")
		fmt.Println("\\/┃\\/                                                      <_I_>             \\/┃\\/")
		fmt.Println("//┃\\\\                                                       |||              //┃\\\\")
		fmt.Println("\\\\┃//                                                      /_|_\\             \\\\┃//")
		fmt.Println("/\\┃/\\                                                                        /\\┃/\\")
		fmt.Println("\\/┃\\/                                                                   /    \\/┃\\/")
		fmt.Println("//┃\\\\                                                               \\  /\\    //┃\\\\")
		fmt.Println("\\\\┃//                                                                \\/¯¯\\   \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	}
}

func CharacterCreationDisplayRace(name, spaces string, pointingAt int) {
	switch pointingAt {
	case 1:
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//   Nom du personnage (max 20 caractères) :             .-.              \\\\┃//")
		fmt.Println("/\\┃/\\                                                     __|=|__            /\\┃/\\")
		fmt.Println("\\/┃\\/       ", name, spaces, "                         (_/`-`\\_)           \\/┃\\/")
		fmt.Println("//┃\\\\                                                    //\\___/\\\\           //┃\\\\")
		fmt.Println("\\\\┃//   Choix de la race :                               <>/   \\<>           \\\\┃//")
		fmt.Println("/\\┃/\\                                                     \\|_._|/            /\\┃/\\")
		fmt.Println("\\/┃\\/     ➢     Humain                                     <_I_>             \\/┃\\/")
		fmt.Println("//┃\\\\           Elfe                                        |||              //┃\\\\")
		fmt.Println("\\\\┃//           Nain                                       /_|_\\             \\\\┃//")
		fmt.Println("/\\┃/\\                                                                        /\\┃/\\")
		fmt.Println("\\/┃\\/   Détails :                                                       /    \\/┃\\/")
		fmt.Println("//┃\\\\          Les humains sont équilibrés dans leurs statistiques  \\  /\\    //┃\\\\")
		fmt.Println("\\\\┃//          et leurs compétences : 100 HP et 100 Mana             \\/¯¯\\   \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	case 2:
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//   Nom du personnage (max 20 caractères) :             .-.              \\\\┃//")
		fmt.Println("/\\┃/\\                                                     __|=|__            /\\┃/\\")
		fmt.Println("\\/┃\\/       ", name, spaces, "                         (_/`-`\\_)           \\/┃\\/")
		fmt.Println("//┃\\\\                                                    //\\___/\\\\           //┃\\\\")
		fmt.Println("\\\\┃//   Choix de la race :                               <>/   \\<>           \\\\┃//")
		fmt.Println("/\\┃/\\                                                     \\|_._|/            /\\┃/\\")
		fmt.Println("\\/┃\\/           Humain                                     <_I_>             \\/┃\\/")
		fmt.Println("//┃\\\\     ➢     Elfe                                        |||              //┃\\\\")
		fmt.Println("\\\\┃//           Nain                                       /_|_\\             \\\\┃//")
		fmt.Println("/\\┃/\\                                                                        /\\┃/\\")
		fmt.Println("\\/┃\\/   Détails :                                                       /    \\/┃\\/")
		fmt.Println("//┃\\\\          Les elfes se spécialisent dans la magie              \\  /\\    //┃\\\\")
		fmt.Println("\\\\┃//          et la perception : 80 HP et 120 Mana                  \\/¯¯\\   \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	case 3:
		ClearTerminal()
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
		fmt.Println(" /┃\\                                                                          /┃\\  ")
		fmt.Println("//┃\\\\                                                                        //┃\\\\")
		fmt.Println("\\\\┃//   Nom du personnage (max 20 caractères) :             .-.              \\\\┃//")
		fmt.Println("/\\┃/\\                                                     __|=|__            /\\┃/\\")
		fmt.Println("\\/┃\\/       ", name, spaces, "                         (_/`-`\\_)           \\/┃\\/")
		fmt.Println("//┃\\\\                                                    //\\___/\\\\           //┃\\\\")
		fmt.Println("\\\\┃//   Choix de la race :                               <>/   \\<>           \\\\┃//")
		fmt.Println("/\\┃/\\                                                     \\|_._|/            /\\┃/\\")
		fmt.Println("\\/┃\\/           Humain                                     <_I_>             \\/┃\\/")
		fmt.Println("//┃\\\\           Elfe                                        |||              //┃\\\\")
		fmt.Println("\\\\┃//     ➢     Nain                                       /_|_\\             \\\\┃//")
		fmt.Println("/\\┃/\\                                                                        /\\┃/\\")
		fmt.Println("\\/┃\\/   Détails :                                                       /    \\/┃\\/")
		fmt.Println("//┃\\\\          Les nains se spécialisent dans l'endurance           \\  /\\    //┃\\\\")
		fmt.Println("\\\\┃//          et leur constitution: 120 HP et 80 Mana               \\/¯¯\\   \\\\┃//")
		fmt.Println(" \\┃/                                                                          \\┃/  ")
		fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	}
}

func asciiArtTesting() {
	ClearTerminal()
	fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
	fmt.Println(" /┃\\                                                                          ◄┃\\")
	fmt.Println("//┃\\\\                                                                         ◀┃\\\\")
	fmt.Println("\\\\┃//                                                                         ◢┃//")
	fmt.Println("/\\┃/\\                                                                         ◥┃/\\")
	fmt.Println("\\/┃\\/                                                                         ◀┃\\/")
	fmt.Println("//┃\\\\                                                                         ◄┃\\\\")
	fmt.Println("\\\\┃//                                                                         ◀┃//")
	fmt.Println("/\\┃/\\                                                                         ◢┃/\\")
	fmt.Println("\\/┃\\/                                                                         ◥┃\\/")
	fmt.Println("//┃\\\\                                                                         ◀┃\\\\")
	fmt.Println("\\\\┃//                                                                         ◄┃//")
	fmt.Println("/\\┃/\\                                                                         ◀┃/\\")
	fmt.Println("\\/┃\\/                                                                         ◢┃\\/")
	fmt.Println("//┃\\\\                                                                         ◥┃\\\\")
	fmt.Println("\\\\┃//                                                                         ◀┃//")
	fmt.Println(" \\┃/                                                                          ◄┃/")
	fmt.Println("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")
}
