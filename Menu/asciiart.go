package menu

import (
	char "ProjetRed/Character"
	"fmt"
	"strconv"
	"time"

	rwidth "github.com/mattn/go-runewidth"
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
	CharMenuText,
	CharMenuMainGrid,
	CharInventoryGrid,
	CharInventoryText,
	MerchantInventoryText,
	SellMerchantInventoryText,
	BlacksmithMenuText,
	StrollGrid,
	StrollCursor,
	CombatGrid [][]rune

var BottomBar,
	CharMenuTitleBar,
	PrincipalMenuTitleBar,
	BonusTitleBar,
	CharCreationMenuTitleBar,
	CharInventoryTitleBar,
	StrollCastleTitleBar,
	StrollBarracksTitleBar,
	StrollCityTitleBar,
	StrollOutsideTitleBar,
	StrollMarketTitleBar,
	StrollArenaTitleBar,
	StrollMerchantTitleBar,
	BlacksmithMenuTitleBar,
	CombatTitleBar,
	CharCreationName,
	CharCreationNameError,
	CharCreationMenuCursor,
	CharMenuCursor,
	CharInventoryItemsCursor []rune

const (
	BONUS_MENU         = -1
	PRINCIPAL_MENU     = 0
	CHAR_CREATION_MENU = 1
	CHAR_MENU          = 2
	CHAR_INVENTORY     = 3
	STROLL_CASTLE      = 4
	STROLL_CITY        = 5
	STROLL_BARRACKS    = 6
	STROLL_OUTSIDE     = 7
	STROLL_MARKET      = 8
	STROLL_ARENA       = 9
	STROLL_MERCHANT    = 10
	STROLL_BLACKSMITH  = 11
	STROLL_MISSIONS    = 12
	BUY_MERCHANT       = 13
	SELL_MERCHANT      = 14
	COMBAT             = 15
	TOURNAMENT         = 16
	TRAINING           = 17
)

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

	ElfDescription = append(ElfDescription, []rune("Les elfes se spécialisent dans la magie            "))
	ElfDescription = append(ElfDescription, []rune("et la perception : 80 HP et 120 Mana     "))

	HumanDescription = append(HumanDescription, []rune("Les humains sont équilibrés dans leurs statistiques"))
	HumanDescription = append(HumanDescription, []rune("et leurs compétences : 100 HP et 100 Mana"))

	DwarfDescription = append(DwarfDescription, []rune("Les nains se spécialisent dans l'endurance         "))
	DwarfDescription = append(DwarfDescription, []rune("et leur constitution : 120 HP et 80 Mana "))

	CharMenuText = append(CharMenuText, []rune("Niveau :"))
	CharMenuText = append(CharMenuText, []rune("Attaque :"))
	CharMenuText = append(CharMenuText, []rune("Défense :"))
	CharMenuText = append(CharMenuText, []rune("Agilité :"))
	CharMenuText = append(CharMenuText, []rune("XP :"))
	CharMenuText = append(CharMenuText, []rune("HP :"))
	CharMenuText = append(CharMenuText, []rune("MP :"))
	CharMenuText = append(CharMenuText, []rune("Or :"))
	CharMenuText = append(CharMenuText, []rune("Tête"))
	CharMenuText = append(CharMenuText, []rune("Torse"))
	CharMenuText = append(CharMenuText, []rune("Mains"))
	CharMenuText = append(CharMenuText, []rune("Jambes"))
	CharMenuText = append(CharMenuText, []rune("Pieds"))
	CharMenuText = append(CharMenuText, []rune("Inventaire"))
	CharMenuText = append(CharMenuText, []rune("Équipement"))
	CharMenuText = append(CharMenuText, []rune("Retour"))
	CharMenuText = append(CharMenuText, []rune("Quitter"))

	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╭─────────────────────────────────╮ ╭───────── Équipement : ─────────╮"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("│ ├────────┬───────────────────────┤"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("├─────────────────┬───────────────┤ │"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("├─────────────────┴───────────────╯ ╰────────┴───────────────────────╯"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╰────────────────╮╭───── Compétence ─────┬ Att ┬ Def ┬─ Stat ──┐Buff ╮"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("│├──────────────────────┼─────┼─────┼─────────┼─────┤"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╰┴──────────────────────┴─────┴─────┴─────────┴─────╯"))

	CharInventoryGrid = append(CharInventoryGrid, []rune("─────────┴──────────────────────────────────────────────────────────────"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("──────────────────┬─────────────────────────────────────────────────────"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("│"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("│"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("│"))
	CharInventoryGrid = append(CharInventoryGrid, []rune("│"))

	CharInventoryText = append(CharInventoryText, []rune("╳  Esc"))
	CharInventoryText = append(CharInventoryText, []rune("Nom"))
	CharInventoryText = append(CharInventoryText, []rune("Type"))
	CharInventoryText = append(CharInventoryText, []rune("Quantité"))
	CharInventoryText = append(CharInventoryText, []rune("Utiliser"))
	CharInventoryText = append(CharInventoryText, []rune("Jeter"))
	CharInventoryText = append(CharInventoryText, []rune("Annuler"))
	CharInventoryText = append(CharInventoryText, []rune("Description :"))

	MerchantInventoryText = append(MerchantInventoryText, []rune("╳  Esc"))
	MerchantInventoryText = append(MerchantInventoryText, []rune("Nom"))
	MerchantInventoryText = append(MerchantInventoryText, []rune("Prix"))
	MerchantInventoryText = append(MerchantInventoryText, []rune("Quantité"))
	MerchantInventoryText = append(MerchantInventoryText, []rune("Acheter"))
	MerchantInventoryText = append(MerchantInventoryText, []rune("Annuler"))
	MerchantInventoryText = append(MerchantInventoryText, []rune("Description :"))

	SellMerchantInventoryText = append(SellMerchantInventoryText, []rune("╳  Esc"))
	SellMerchantInventoryText = append(SellMerchantInventoryText, []rune("Nom"))
	SellMerchantInventoryText = append(SellMerchantInventoryText, []rune("Prix"))
	SellMerchantInventoryText = append(SellMerchantInventoryText, []rune("Quantité"))
	SellMerchantInventoryText = append(SellMerchantInventoryText, []rune("Vendre"))
	SellMerchantInventoryText = append(SellMerchantInventoryText, []rune("Annuler"))
	SellMerchantInventoryText = append(SellMerchantInventoryText, []rune("Description :"))

	BlacksmithMenuText = append(BlacksmithMenuText, []rune(" ╳  Esc  │         Nom                         Ressources nécessaires"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("─────────┴─────────────────────────────────┬────────────────────────────"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("──────────────────┬────────────────────────┴────────────────────────────"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("Fabriquer  │  Description :"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("Annuler    │"))
	BlacksmithMenuText = append(BlacksmithMenuText, []rune("│"))

	StrollGrid = append(StrollGrid, []rune("─────────────────────────────────────────┴──────────────┴───────────────"))
	StrollGrid = append(StrollGrid, []rune("      ╭────────────╮                                 ╭────────────╮     "))
	StrollGrid = append(StrollGrid, []rune("      ╰─────┬──────╯                                 ╰─────┬──────╯     "))
	StrollGrid = append(StrollGrid, []rune("────────────┴──────────────────────────────────────────────┴────────────"))
	StrollGrid = append(StrollGrid, []rune("                             ╭───────────╮             ╭────────────────"))

	StrollCursor = append(StrollCursor, []rune(" O "))
	StrollCursor = append(StrollCursor, []rune("/|\\"))
	StrollCursor = append(StrollCursor, []rune("/‾\\"))

	CombatGrid = append(CombatGrid, []rune("────────────────────────────────────────────────────────────────────────"))
	CombatGrid = append(CombatGrid, []rune("                                                      O"))
	CombatGrid = append(CombatGrid, []rune("──────────────────╮                                  /|\\"))
	CombatGrid = append(CombatGrid, []rune(" Niveau"))
	CombatGrid = append(CombatGrid, []rune("│                                  /‾\\"))
	CombatGrid = append(CombatGrid, []rune(" HP :"))
	CombatGrid = append(CombatGrid, []rune("│"))
	CombatGrid = append(CombatGrid, []rune(" MP :"))
	CombatGrid = append(CombatGrid, []rune("│"))
	CombatGrid = append(CombatGrid, []rune(" Attaque :"))
	CombatGrid = append(CombatGrid, []rune("│     O"))
	CombatGrid = append(CombatGrid, []rune(" Défense :"))
	CombatGrid = append(CombatGrid, []rune("│    /|\\"))
	CombatGrid = append(CombatGrid, []rune(" Agilité :"))
	CombatGrid = append(CombatGrid, []rune("│    /‾\\"))
	CombatGrid = append(CombatGrid, []rune("──────────────────┼─────────────────────────────────────────────────────"))
	CombatGrid = append(CombatGrid, []rune("Attaquer   │"))
	CombatGrid = append(CombatGrid, []rune("Défendre   │"))
	CombatGrid = append(CombatGrid, []rune("Potion     │"))
	CombatGrid = append(CombatGrid, []rune("           │"))
	CombatGrid = append(CombatGrid, []rune("Fuir       │"))

	BottomBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	PrincipalMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu principal ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	BonusTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Bonus ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharCreationMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Création du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharInventoryTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Inventaire ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	StrollCastleTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Château ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	StrollBarracksTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Caserne ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	StrollCityTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Ville ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	StrollOutsideTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Dehors ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	StrollMarketTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Marché ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	StrollMerchantTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Marchand ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	StrollArenaTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Arène ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	BlacksmithMenuTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Forgeron ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CombatTitleBar = []rune("  ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Combat ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪")

	CharCreationName = []rune("Nom du personnage (max 20 caractères) :")

	CharCreationNameError = []rune("Erreur : veuillez rentrer un nom valide !")

	CharCreationMenuCursor = []rune("⎯{===-")

	CharMenuCursor = []rune("⎯{==-")

	CharInventoryItemsCursor = []rune("⎯{====-")
}

func LoadingScreen() {
	ClearTerminal()
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

func displayTopBar(menuNb int) {
	var CurrentTopBar []rune
	switch menuNb {

	case PRINCIPAL_MENU:
		CurrentTopBar = PrincipalMenuTitleBar

	case BONUS_MENU:
		CurrentTopBar = BonusTitleBar

	case CHAR_CREATION_MENU:
		CurrentTopBar = CharCreationMenuTitleBar

	case CHAR_MENU:
		CurrentTopBar = CharMenuTitleBar

	case CHAR_INVENTORY:
		CurrentTopBar = CharInventoryTitleBar

	case STROLL_CASTLE:
		CurrentTopBar = StrollCastleTitleBar

	case STROLL_BARRACKS:
		CurrentTopBar = StrollBarracksTitleBar

	case STROLL_CITY:
		CurrentTopBar = StrollCityTitleBar

	case STROLL_MARKET:
		CurrentTopBar = StrollMarketTitleBar

	case STROLL_OUTSIDE:
		CurrentTopBar = StrollOutsideTitleBar

	case STROLL_MERCHANT, BUY_MERCHANT, SELL_MERCHANT:
		CurrentTopBar = StrollMerchantTitleBar

	case STROLL_BLACKSMITH:
		CurrentTopBar = BlacksmithMenuTitleBar

	case STROLL_ARENA:
		CurrentTopBar = StrollArenaTitleBar

	case STROLL_MISSIONS:
		//CurrentTopBar = StrollMissionsTitleBar

	case COMBAT, TOURNAMENT, TRAINING:
		CurrentTopBar = CombatTitleBar
	}
	column := 0
	for _, char := range CurrentTopBar {
		DisplayRune(column, 0, char)
		column += rwidth.RuneWidth(char)
	}
}

func displayMenuBars() {
	for i := range MenuLateralBar {
		column := 0
		for _, char := range MenuLateralBar[i] {
			DisplayRune(column, i+1, char)
			column += rwidth.RuneWidth(char)
		}
	}
	for i := range MenuLateralBar {
		column := 77
		for _, char := range MenuLateralBar[i] {
			DisplayRune(column, i+1, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayBottomBar() {
	column := 0
	for _, char := range BottomBar {
		DisplayRune(column, 17, char)
		column += rwidth.RuneWidth(char)
	}
}

func displayLogo() {
	column := 69
	line := 13
	for i := range LogoLittle {
		column = 69
		for _, char := range LogoLittle[i] {
			DisplayRune(column, line+i, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displaySoldier() {
	column := 57
	line := 3
	for i := range SoldierArt {
		column = 57
		for _, char := range SoldierArt[i] {
			DisplayRune(column, line+i, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func DisplayBlankMenu(menuNb int) {
	displayTopBar(menuNb)
	displayMenuBars()
	displayBottomBar()
	if menuNb == CHAR_CREATION_MENU || menuNb == PRINCIPAL_MENU {
		displayLogo()
	}
	if menuNb == CHAR_CREATION_MENU {
		displaySoldier()
	}
}

func DisplayMenuOptions(menuNb int) {
	var menuOptions [][]rune
	switch menuNb {
	case PRINCIPAL_MENU:
		menuOptions = PrincipalMenuOptions
	}
	column := 36
	line := 4
	for i := range menuOptions {
		column = 36
		for _, char := range menuOptions[i] {
			DisplayRune(column, line+(2*i), char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func DisplayPrincipalCursor(option, previousOption int) {

	var cursor [][]rune
	var column, line int
	column = 11

	if previousOption != 0 {

		cursor = append(cursor, []rune("      "))
		cursor = append(cursor, []rune("                   "))
		cursor = append(cursor, []rune("      "))

		switch previousOption {

		case 1:
			line = 3
		case 2:
			line = 5
		case 3:
			line = 7
		case 4:
			line = 9
		case 5:
			line = 11
		}
		for i := range cursor {
			column = 11
			for _, char := range cursor[i] {
				DisplayRune(column, line+i, char)
				column += rwidth.RuneWidth(char)
			}
		}

		previousOption = 0

		DisplayPrincipalCursor(option, previousOption)

	} else {

		switch option {

		case 1:
			line = 3
		case 2:
			line = 5
		case 3:
			line = 7
		case 4:
			line = 9
		case 5:
			line = 11
		}

		for i := range Sword {
			column = 11
			for _, char := range Sword[i] {
				DisplayRune(column, line+i, char)
				column += rwidth.RuneWidth(char)
			}
		}
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

func CharCreationNameDisplay(err bool) {
	column := 8
	for _, char := range CharCreationName {
		DisplayRune(column, 3, char)
		column += rwidth.RuneWidth(char)
	}
	if err {
		column = 8
		for _, char := range CharCreationNameError {
			DisplayRune(column, 2, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func CharCreationMenuMainDisplay(name string) {

	DisplayText(13, 5, name)

	var column int
	var lines = []int{7, 9, 10, 11, 13}

	for i := range CharCreationMenuText {

		switch i {

		case 0, 4:
			column = 8

		case 1, 2, 3:
			column = 16

		}
		for _, char := range CharCreationMenuText[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func CharCreationMenuChangingDisplay(option, previousOption int) {

	column := 7
	var lines = []int{9, 10, 11}

	DisplayText(column, lines[option-1], string(CharCreationMenuCursor))

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "      ")
	}

	var description [][]rune

	switch option {

	case 1:
		description = HumanDescription

	case 2:
		description = ElfDescription

	case 3:
		description = DwarfDescription

	}
	line := 14
	column = 16
	for i := range description {
		column = 16
		for _, char := range description[i] {
			DisplayRune(column, line+i, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayCharMenuGrid() {

	var columns = []int{6, 40, 6, 6, 6, 23, 23}
	var lines = []int{1, 2, 3, 8, 9, 10, 16}
	var column int

	for i := range CharMenuMainGrid {
		column = columns[i]
		for _, char := range CharMenuMainGrid[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}

	columns = []int{6, 51, 75, 6, 6, 6, 6, 24, 24, 24, 24, 40, 40, 40, 40, 42, 42, 42, 42, 51, 51, 51, 51, 75, 75, 75, 75, 23, 23, 23, 23, 23, 24, 24, 24, 24, 24, 47, 47, 47, 47, 47, 53, 53, 53, 53, 53, 59, 59, 59, 59, 59, 69, 69, 69, 69, 69, 75, 75, 75, 75, 75}
	lines = []int{2, 3, 3, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 4, 5, 6, 7, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15, 11, 12, 13, 14, 15}

	for i, column := range columns {
		DisplayRune(column, lines[i], '│')
	}
}

func displayCharMenuText() {
	var column int
	var columns = []int{8, 8, 8, 8, 26, 26, 26, 26, 44, 44, 44, 44, 44, 12, 12, 12, 12}
	var lines = []int{4, 5, 6, 7, 4, 5, 6, 7, 3, 4, 5, 6, 7, 10, 12, 14, 15}

	for i := range CharMenuText {
		column = columns[i]
		for _, char := range CharMenuText[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func DisplayCharMenuCursor(option, previousOption int) {
	column := 6
	var lines = []int{10, 12, 14, 15}

	DisplayText(column, lines[option-1], string(CharMenuCursor))

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "     ")
	}
}

func displayCharMenuStats(myChar *char.Character) {
	var column int
	var columns = []int{9, 32, 20, 20, 20, 20, 31, 31, 31, 31, 53, 53, 53, 53, 53, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71, 26, 49, 55, 61, 71}
	var lines = []int{2, 2, 4, 5, 6, 7, 4, 5, 6, 7, 3, 4, 5, 6, 7, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 13, 13, 13, 13, 13, 14, 14, 14, 14, 14, 15, 15, 15, 15, 15}
	var stats []string
	stats = append(stats, myChar.Name)
	stats = append(stats, myChar.Class.Name)
	stats = append(stats, strconv.Itoa(myChar.Level))
	stats = append(stats, strconv.Itoa(myChar.Attack))
	stats = append(stats, strconv.Itoa(myChar.Defense))
	stats = append(stats, strconv.Itoa(myChar.Agility))
	stats = append(stats, strconv.Itoa(myChar.Xp))
	stats = append(stats, strconv.Itoa(myChar.Hp)+"/"+strconv.Itoa(myChar.HpMax))
	stats = append(stats, strconv.Itoa(myChar.Mp)+"/"+strconv.Itoa(myChar.MpMax))
	stats = append(stats, strconv.Itoa(myChar.Gold))
	stats = append(stats, "rien")
	stats = append(stats, "rien")
	stats = append(stats, "rien")
	stats = append(stats, "rien")
	stats = append(stats, "rien")
	for _, skill := range myChar.Skills {
		stats = append(stats, skill.Name)
		stats = append(stats, strconv.Itoa(skill.Attack))
		stats = append(stats, strconv.Itoa(skill.Defense))
		stats = append(stats, skill.StatBuffed)
		stats = append(stats, strconv.Itoa(skill.Buff))
	}
	for i, singleStat := range stats {
		column = columns[i]
		for _, char := range singleStat {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayCharMenu(myChar *char.Character) {
	DisplayBlankMenu(CHAR_MENU)
	displayCharMenuGrid()
	displayCharMenuText()
	displayCharMenuStats(myChar)
}

func displayCharInventoryGrid() {
	var columns = []int{5, 5, 14, 23, 23, 23}
	var lines = []int{2, 13, 1, 14, 15, 16}
	var column int

	for i := range CharInventoryGrid {
		column = columns[i]
		for _, char := range CharInventoryGrid[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayCharInventoryText(whichMenu int) {

	var columns []int
	var lines []int
	var column int
	var inventoryText [][]rune

	switch whichMenu {
	case CHAR_INVENTORY:
		inventoryText = CharInventoryText
		columns = []int{6, 24, 49, 64, 12, 12, 12, 26}
		lines = []int{1, 1, 1, 1, 14, 15, 16, 14}
	case BUY_MERCHANT:
		inventoryText = MerchantInventoryText
		columns = []int{6, 24, 49, 64, 12, 12, 26}
		lines = []int{1, 1, 1, 1, 14, 15, 14}
	case SELL_MERCHANT:
		inventoryText = SellMerchantInventoryText
		columns = []int{6, 24, 49, 64, 12, 12, 26}
		lines = []int{1, 1, 1, 1, 14, 15, 14}
	}

	for i := range inventoryText {
		column = columns[i]
		for _, char := range inventoryText[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayInventoryItems(myChar *char.Character, itemOptions *int, whichMenu int) []string {

	if len(myChar.Inventory) == 0 {
		DisplayText(17, 4, "Votre inventaire est vide !")
		time.Sleep(time.Second * 2)
		return nil
	}

	var inventory []string
	for item := range myChar.Inventory {
		inventory = append(inventory, item)
	}

	*itemOptions = len(inventory)

	line := 3
	for index, item := range inventory {
		if index == 10 {
			break
		}
		DisplayText(17, line+index, item)
		if whichMenu == CHAR_INVENTORY {
			DisplayText(47, line+index, retreiveItemType(item))
		} else {
			switch retreiveItemType(item) {
			case "Potion":
				for _, potion := range char.AllPotion {
					if item == potion.Name {
						DisplayText(49, line+index, strconv.Itoa(potion.Price))
					}
				}
			case "Livre de sort":
				for _, spellBook := range char.AllSpellBook {
					if item == spellBook.Name {
						DisplayText(49, line+index, strconv.Itoa(spellBook.Price))
					}
				}
			case "Equipement":
				for _, equipement := range char.AllEquipement {
					if item == equipement.Name {
						DisplayText(49, line+index, strconv.Itoa(equipement.Price))
					}
				}
			case "Ressource":
				for _, ressource := range char.AllRessources {
					if item == ressource.Name {
						DisplayText(49, line+index, strconv.Itoa(ressource.Price))
					}
				}
			}

		}
		DisplayText(67, line+index, strconv.Itoa(myChar.Inventory[item]))
	}
	return inventory
}

func displayInventory(myChar char.Character, itemOptions *int, whichMenu int) []string {
	DisplayBlankMenu(whichMenu)
	displayCharInventoryGrid()
	displayCharInventoryText(whichMenu)
	return displayInventoryItems(&myChar, itemOptions, whichMenu)
}

func displayCharInventoryItemsCursor(option, previousOption int) {
	column := 7
	var lines = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	DisplayText(column, lines[option-1], string(CharInventoryItemsCursor))

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "       ")
	}
}

func displayItemDescription(item string) {

	DisplayText(29, 15, "                                                ")
	DisplayText(29, 16, "                                                ")

	var description string

	switch retreiveItemType(item) {
	case "Équipement":
		for _, singleItem := range char.AllEquipement {
			if item == singleItem.Name {
				description = singleItem.Name + " est un équipement.  " // Add Description field in the Equipment struct!!!
			}
		}
	case "Potion":
		for _, singleItem := range char.AllPotion {
			if item == singleItem.Name {
				description = singleItem.Name + " est une potion.       " // Add Description field in the Potion struct!!!
			}
		}
	case "Livre de sort":
		for _, singleItem := range char.AllSpellBook {
			if item == singleItem.Name {
				description = singleItem.Name + " est un livre de sort.  " // Add Description field in the SpellBook struct!!!
			}
		}
	case "Ressource":
		for _, singleItem := range char.AllRessources {
			if item == singleItem.Name {
				description = singleItem.Name + " est une ressource.  " // Add Description field in the Ressource struct!!!
			}
		}
	default:
		description = "Item inconnu"
	}

	DisplayText(29, 15, description)
}

func displayCharInventoryActionCursor(option, previousOption, whichMenu int) {
	column := 6
	var lines []int
	if whichMenu == CHAR_INVENTORY {
		lines = []int{14, 15, 16}
	} else {
		lines = []int{14, 15}
	}
	if option != 0 {
		DisplayText(column, lines[option-1], string(CharMenuCursor))
	}

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "     ")
	}
}

// Rajouter tous les types d'item dans la fonction afin de prendre en compte toutes les possibilités !
func retreiveItemType(s string) string {

	for _, name := range char.AllEquipement {
		if s == name.Name {
			return "Équipement"
		}
	}
	for _, name := range char.AllPotion {
		if s == name.Name {
			return "Potion"
		}
	}
	for _, name := range char.AllSpellBook {
		if s == name.Name {
			return "Livre de sort"
		}
	}
	for _, name := range char.AllRessources {
		if s == name.Name {
			return "Ressource"
		}
	}
	return "Inconnu" // Dommage, cet item n'est par répertorié dans les slices présentes dans la fonction... :´(
}

func displayStrollGrid() {
	column := 5
	var lines = []int{2, 8, 10, 11, 15}
	for i := range StrollGrid {
		column = 5
		for _, char := range StrollGrid[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayStrollText(myChar *char.Character, nbMenu int) {

	var text = []string{myChar.Name, myChar.Class.Name, "│ HP :", strconv.Itoa(myChar.Hp) + "/" + strconv.Itoa(myChar.HpMax), "│  Or :", strconv.Itoa(myChar.Gold), "│ Space :  Menu   "}

	switch nbMenu {
	case STROLL_CASTLE:
		text = append(text, "│  Caserne   │")
		text = append(text, "│   Ville    │")
		text = append(text, "│  Château  │")
	case STROLL_BARRACKS:
		text = append(text, "│Entrainement│")
		text = append(text, "│  Tournoi   │")
		text = append(text, "│  Château  │")
	case STROLL_CITY:
		text = append(text, "│   Dehors   │")
		text = append(text, "│   Marché   │")
		text = append(text, "│  Château  │")
	case STROLL_OUTSIDE:
		text = append(text, "│   Arène    │")
		text = append(text, "│  Mission   │")
		text = append(text, "│   Ville   │")
	case STROLL_MARKET:
		text = append(text, "│  Marchand  │")
		text = append(text, "│  Forgeron  │")
		text = append(text, "│   Ville   │")
	case STROLL_ARENA:
		// Go to the Arena menu!
	case STROLL_MERCHANT:
		text = append(text, "│  Acheter   │")
		text = append(text, "│   Vendre   │")
		text = append(text, "│  Marché   │")
	case STROLL_BLACKSMITH:
		// Go to the Blacksmith menu!
	}

	var columns = []int{7, 31, 46, 53, 61, 69, 60, 11, 58, 34}
	var lines = []int{1, 1, 1, 1, 1, 1, 16, 9, 9, 16}
	var column int

	for i, str := range text {
		column = columns[i]
		for _, char := range str {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayStrollCursor(pointingAt, previousPointingAt int) {

	var columns = []int{16, 39, 63}
	var column int
	line := 12

	for i := range StrollCursor {
		column = columns[pointingAt-1]
		for _, char := range StrollCursor[i] {
			DisplayRune(column, line+i, char)
			column += rwidth.RuneWidth(char)
		}
	}

	if previousPointingAt != 0 {
		for i := 0; i < 3; i++ {
			DisplayText(columns[previousPointingAt-1], line+i, "   ")
		}
	}
}

func displayStrollMenu(myChar *char.Character, nbMenu int) {
	DisplayBlankMenu(nbMenu)
	displayStrollGrid()
	displayStrollText(myChar, nbMenu)
}

func displayBlacksmithBasicMenu() {
	var columns = []int{5, 5, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 5, 12, 12, 23}
	var column int
	line := 1

	for i := range BlacksmithMenuText {
		column = columns[i]
		for _, char := range BlacksmithMenuText[i] {
			DisplayRune(column, line+i, char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayBlacksmithEquipment() {
	column := 17
	line := 3

	for i, equipment := range char.BlacksmithEquipments {
		DisplayText(column, line+i, equipment.Name)
	}
}

func displayBlacksmithRecipe(pointingAt int) {
	column := 50
	line := 4

	for i := 0; i < 9; i++ {
		DisplayText(column, line+i, "                           ")
	}
	DisplayText(column, line, "10   Pièces d'or")
	line += 2
	i := 0
	for ressource, nb := range char.BlacksmithEquipments[pointingAt-1].Recipe {
		DisplayText(column, line+i, strconv.Itoa(nb))
		DisplayText(column+5, line+i, ressource.Name)
		i += 2
	}
}

func displayBlacksmithMenu() {
	DisplayBlankMenu(STROLL_BLACKSMITH)
	displayBlacksmithBasicMenu()
	displayBlacksmithEquipment()
}

func displayCombatGrid() {
	columns := []int{5, 5, 5, 5, 23, 5, 23, 5, 23, 5, 23, 5, 23, 5, 23, 5, 12, 12, 12, 12, 12}
	lines := []int{2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 12, 13, 14, 15, 16}
	var column int

	for i := range CombatGrid {
		column = columns[i]
		for _, char := range CombatGrid[i] {
			DisplayRune(column, lines[i], char)
			column += rwidth.RuneWidth(char)
		}
	}
}

func displayCombatStats(myChar *char.Character, otherChar *char.Enemy) {
	var StatsToDisplay []string
	StatsToDisplay = append(StatsToDisplay, otherChar.Name)
	StatsToDisplay = append(StatsToDisplay, otherChar.Race)
	StatsToDisplay = append(StatsToDisplay, "Niveau :")
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(otherChar.Level))
	StatsToDisplay = append(StatsToDisplay, "HP :")
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(otherChar.Hp)+"/"+strconv.Itoa(otherChar.HpMax))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Level))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Hp)+"/"+strconv.Itoa(myChar.HpMax))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Mp)+"/"+strconv.Itoa(myChar.MpMax))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Attack))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Defense))
	StatsToDisplay = append(StatsToDisplay, strconv.Itoa(myChar.Agility))

	columns := []int{6, 28, 45, 53, 62, 67, 18, 15, 15, 18, 18, 18}
	lines := []int{1, 1, 1, 1, 1, 1, 5, 6, 7, 8, 9, 10}
	var column int

	for i, str := range StatsToDisplay {
		column = columns[i]
		DisplayText(column, lines[i], str)
	}
}

func displayCombatSkills(myChar *char.Character) {
	line := 12
	for i, skill := range myChar.Skills {
		DisplayText(31, line+i, skill.Name)
		DisplayText(57, line+i, skill.StatBuffed)
		DisplayText(68, line+i, strconv.Itoa(skill.MpCost))
		DisplayText(73, line+i, "MP")
	}
}

func displayCombatPotions(myChar *char.Character) []string {
	line := 12
	var myPotions []string
	for item := range myChar.Inventory {
		if retreiveItemType(item) == "Potion" {
			myPotions = append(myPotions, item)
		}
	}
	for i, potion := range myPotions {
		currentPotion := char.FindPotion(potion)
		DisplayText(31, line+i, currentPotion.Name)
		if currentPotion.Buff != 0 {
			DisplayText(49, line+i, "+"+strconv.Itoa(currentPotion.Buff))
			DisplayText(53, line+i, currentPotion.StatBuffed)
		}
		if currentPotion.Debuff != 0 {
			DisplayText(62, line+i, "-"+strconv.Itoa(currentPotion.Debuff))
			DisplayText(66, line+i, currentPotion.StatDebuffed)
		}
		DisplayText(74, line+i, strconv.Itoa(myChar.Inventory[potion]))
	}
	return myPotions
}

func clearDisplayCombatOptions() {
	for i := 12; i < 17; i++ {
		DisplayText(24, i, "                                                     ")
	}
}

func displayCombatActionTypeCursor(option, previousOption int) {
	column := 6
	lines := []int{12, 13, 14, 16}

	if option != 0 {
		DisplayText(column, lines[option-1], string(CharMenuCursor))
	}

	if previousOption != 0 {
		DisplayText(column, lines[previousOption-1], "     ")
	}
}

func displayCombatSpecificCursor(option, previousOption int) {
	column := 25
	line := 12
	if option != 0 {
		DisplayText(column, line+option-1, string(CharMenuCursor))
	}

	if previousOption != 0 {
		DisplayText(column, line+previousOption-1, "     ")
	}
}

func DisplayCombatMenu(myChar *char.Character, enemy *char.Enemy, combatType int) {
	DisplayBlankMenu(combatType)
	displayCombatGrid()
	displayCombatStats(myChar, enemy)
}

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Menu du personnage ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\  ╭─────────────────────────────────╮ ╭───────── Équipement : ─────────╮  /┃\
//  //┃\\ │  MonPseudoEstTropLong   Humain  │ ├────────┬───────────────────────┤ //┃\\
//  \\┃// ┟─────────────────┬───────────────┧ ╽ Tête   │ Heaume du chasseur    ╽ \\┃//
//  /\┃/\ ┃ Niveau :    100 │ XP : 24.785   ┃ ┃ Torse  │ Armure de fantassin   ┃ /\┃/\
//  \/┃\/ ┃ Attaque :   150 │ HP : 350/400  ┃ ┃ Mains  │ Épée longue en mithril┃ \/┃\/
//  //┃\\ ╿ Défense :   170 │ MP : 430/600  ╿ ╿ Jambes │ Jambières elfiques    ╿ //┃\\
//  \\┃// │ Agilité :   275 │ Or : 12.570   │ │ Pieds  │ Chausses légères      │ \\┃//
//  /\┃/\ ├─────────────────┴───────────────╯ ╰────────┴───────────────────────╯ /\┃/\
//  \/┃\/ ╰────────────────╮╭───── Compétence ─────┬ Att ┬ Def ┬─ Stat ──┐Buff ╮ \/┃\/
//  //┃\\ ⎯{==- Inventaire │┟──────────────────────┼─────┼─────┼─────────┼─────┧ //┃\\
//  \\┃//      (Croissance)│┃ Coup de poing humain │ 240 │ 100 │ Agilité │ 105 ┃ \\┃//
//  /\┃/\       Équipement │┃ Coup de poing humain │ 240 │ 100 │ Attaque │ 105 ┃ /\┃/\
//  \/┃\/                  │╿ Coup de poing humain │ 240 │ 100 │ Défense │ 105 ╿ \/┃\/
//  //┃\\       Retour     ││ Coup de poing humain │ 240 │ 100 │ Attaque │ 105 │ //┃\\
//  \\┃//       Quitter    ││ Coup de poing humain │ 240 │ 100 │ Attaque │ 105 │ \\┃//
//   \┃/                   ╰┴──────────────────────┴─────┴─────┴─────────┴─────╯  \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Inventaire ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\  ╳  Esc  │         Nom                      Type           Quantité      /┃\
//  //┃\\─────────┴──────────────────────────────────────────────────────────────//┃\\
//  \\┃//  ⎯{====-   Heaume du chasseur            Équipement          2         \\┃//
//  /\┃/\            Armure de fantassin           Équipement          1         /\┃/\
//  \/┃\/            Épée longue en mithril        Équipement          1         \/┃\/
//  //┃\\            Jambières elfiques            Équipement          1         //┃\\
//  \\┃//            Chausses légères              Équipement          2         \\┃//
//  /\┃/\            Potion de soin                Potion              10        /\┃/\
//  \/┃\/            Potion de poison              Potion              24        \/┃\/
//  //┃\\            Boule de feu                  Livre de sort       1         //┃\\
//  \\┃//            Peau de mammouth              Ressource           35        \\┃//
//  /\┃/\            Morceau de mithril            Ressource           7         /\┃/\
//  \/┃\/──────────────────┬─────────────────────────────────────────────────────\/┃\/
//  //┃\\       Utiliser   │  Description :                                      //┃\\
//  \\┃// ⎯{==- Jeter      │     Le heaume du chasseur donne 35 pts de défense   \\┃//
//   \┃/        Annuler    │     et 15 pts d'attaque.                             \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Château ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\   MonPseudoEstTropLong    Humain         │ HP : 350/400 │  Or : 12.570   /┃\
//  //┃\\─────────────────────────────────────────┴──────────────┴───────────────//┃\\
//  \\┃//                                                                        \\┃//
//  /\┃/\                                                                        /\┃/\
//  \/┃\/                                                                        \/┃\/
//  //┃\\                                                                        //┃\\
//  \\┃//                                                                        \\┃//
//  /\┃/\      ╭────────────╮                                 ╭────────────╮     /\┃/\
//  \/┃\/      │   Ville    │                                 │Entrainement│     \/┃\/
//  //┃\\      ╰─────┬──────╯                                 ╰─────┬──────╯     //┃\\
//  \\┃//────────────┴──────────────────────────────────────────────┴────────────\\┃//
//  /\┃/\                                                           ₀            /\┃/\
//  \/┃\/                                                          ─╁─           \/┃\/
//  //┃\\                                                           Λ            //┃\\
//  \\┃//                             ╭───────────╮             ╭────────────────\\┃//
//   \┃/                              │  Château  │             │ Space :  Menu   \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Marchand ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\  ╳  Esc  │         Nom                      Prix           Quantité      /┃\
//  //┃\\─────────┴──────────────────────────────────────────────────────────────//┃\\
//  \\┃//  ⎯{====-   Heaume du chasseur              1.300             2         \\┃//
//  /\┃/\            Armure de fantassin             130               1         /\┃/\
//  \/┃\/            Épée longue en mithril          25.480            1         \/┃\/
//  //┃\\            Jambières elfiques              3.510             1         //┃\\
//  \\┃//            Chausses légères                60                2         \\┃//
//  /\┃/\            Potion de soin                  20                10        /\┃/\
//  \/┃\/            Potion de poison                15                24        \/┃\/
//  //┃\\            Grimoire boule de feu           315               1         //┃\\
//  \\┃//            Peau de mammouth                35                35        \\┃//
//  /\┃/\            Morceau de mithril              1.200             7         /\┃/\
//  \/┃\/──────────────────┬─────────────────────────────────────────────────────\/┃\/
//  //┃\\ ⎯{==- Acheter    │  Description :                                      //┃\\
//  \\┃//       Annuler    │     Le heaume du chasseur donne 35 pts de défense   \\┃//
//   \┃/                   │     et 15 pts d'attaque.                             \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Forgeron ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\  ╳  Esc  │         Nom                         Ressources nécessaires    /┃\
//  //┃\\─────────┴─────────────────────────────────┬────────────────────────────//┃\\
//  \\┃//  ⎯{====-   Épée longue en mithril         │                            \\┃//
//  /\┃/\            Armure de fantassin            │ 12.550 Pièces d'or         /\┃/\
//  \/┃\/            Heaume du chasseur             │                            \/┃\/
//  //┃\\            Jambières elfiques             │ 12     Morceau de mithril  //┃\\
//  \\┃//            Chausses légères               │                            \\┃//
//  /\┃/\            Chapeau de l'aventurier        │ 1      Écaille de Dragon   /\┃/\
//  \/┃\/            Bottes de l'aventurier         │                            \/┃\/
//  //┃\\            Épée d'entrainement            │ 5      Morceau d'argent    //┃\\
//  \\┃//            Tunique de l'aventurier        │                            \\┃//
//  /\┃/\            Tunique de mage                │ 1      Ruby                /\┃/\
//  \/┃\/──────────────────┬────────────────────────┴────────────────────────────\/┃\/
//  //┃\\ ⎯{==- Fabriquer  │  Description :                                      //┃\\
//  \\┃//       Annuler    │     Le heaume du chasseur donne 35 pts de défense   \\┃//
//   \┃/                   │     et 15 pts d'attaque.                             \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Combat ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\  NomEnnemiACombattre   NomRace          Niveau  100      HP   480/850    /┃\
//  //┃\\────────────────────────────────────────────────────────────────────────//┃\\
//  \\┃//                                                      O                 \\┃//
//  /\┃/\──────────────────╮                                  /|\                /\┃/\
//  \/┃\/ Niveau      100  │                                  /‾\                \/┃\/
//  //┃\\ HP :     350/400 │                                                     //┃\\
//  \\┃// MP :     430/600 │                                                     \\┃//
//  /\┃/\ Attaque :   150  │     O                                               /\┃/\
//  \/┃\/ Défense :   170  │    /|\                                              \/┃\/
//  //┃\\ Agilité :   275  │    /‾\                                              //┃\\
//  \\┃//──────────────────┼─────────────────────────────────────────────────────\\┃//
//  /\┃/\ ⎯{==- Attaquer   │ ⎯{==- Coup de poing humain      Agilité    105  MP  /\┃/\
//  \/┃\/       Défendre   │       Coup de poing humain      Attaque    105  MP  \/┃\/
//  //┃\\       Potion     │       Coup de poing humain      Défense    105  MP  //┃\\
//  \\┃//                  │       Coup de poing humain      Attaque    105  MP  \\┃//
//   \┃/        Fuir       │       Coup de poing humain      Attaque    105  MP   \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪

//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪ Combat ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
//   /┃\  NomEnnemiACombattre   NomRace          Niveau  100      HP   480/850    /┃\
//  //┃\\────────────────────────────────────────────────────────────────────────//┃\\
//  \\┃//                                                      O                 \\┃//
//  /\┃/\──────────────────╮                                  /|\                /\┃/\
//  \/┃\/ Niveau      100  │                                  /‾\                \/┃\/
//  //┃\\ HP :     350/400 │                                                     //┃\\
//  \\┃// MP :     430/600 │                                                     \\┃//
//  /\┃/\ Attaque :   150  │     O                                               /\┃/\
//  \/┃\/ Défense :   170  │    /|\                                              \/┃\/
//  //┃\\ Agilité :   275  │    /‾\                                              //┃\\
//  \\┃//──────────────────┼─────────────────────────────────────────────────────\\┃//
//  /\┃/\       Attaquer   │ ⎯{==- Potion de soin    +20 HP                   10 /\┃/\
//  \/┃\/       Défendre   │       Potion de poison               -20 HP      24 \/┃\/
//  //┃\\ ⎯{==- Potion     │       Potion de défense +50 Défense  -15 Agilité 5  //┃\\
//  \\┃//                  │       Potion de mana    +20 MP                   3  \\┃//
//   \┃/        Fuir       │       Potion Berserker  +50 Attaque  -30 HP      2   \┃/
//    ₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪₪
