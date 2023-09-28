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
	CharMenuText,
	CharMenuMainGrid,
	CharInventoryGrid,
	CharInventoryText,
	CharInventoryTextBlacksmithFacility,
	MerchantInventoryText,
	SellMerchantInventoryText,
	BlacksmithMenuText,
	StrollGrid,
	StrollCursor,
	CombatGrid,
	ByeDisplay [][]rune

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
	CharMenuText = append(CharMenuText, []rune("Retour"))
	CharMenuText = append(CharMenuText, []rune("Quitter"))

	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╭─────────────────────────────────╮ ╭───────── Équipement : ─────────╮"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("│ ├────────┬───────────────────────┤"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("├─────────────────┬───────────────┤ │"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("├─────────────────┴───────────────╯ ╰────────┴───────────────────────╯"))
	CharMenuMainGrid = append(CharMenuMainGrid, []rune("╰────────────────╮╭───── Compétence ─────┬ Att ┬ Def ┬─ Type ──┬ MP ─╮"))
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

	CharInventoryTextBlacksmithFacility = append(CharInventoryTextBlacksmithFacility, []rune("╳  Esc"))
	CharInventoryTextBlacksmithFacility = append(CharInventoryTextBlacksmithFacility, []rune("Nom"))
	CharInventoryTextBlacksmithFacility = append(CharInventoryTextBlacksmithFacility, []rune("Type"))
	CharInventoryTextBlacksmithFacility = append(CharInventoryTextBlacksmithFacility, []rune("Quantité"))
	CharInventoryTextBlacksmithFacility = append(CharInventoryTextBlacksmithFacility, []rune("Utiliser"))
	CharInventoryTextBlacksmithFacility = append(CharInventoryTextBlacksmithFacility, []rune("Recycler"))
	CharInventoryTextBlacksmithFacility = append(CharInventoryTextBlacksmithFacility, []rune("Annuler"))
	CharInventoryTextBlacksmithFacility = append(CharInventoryTextBlacksmithFacility, []rune("Description :"))

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

	ByeDisplay = append(ByeDisplay, []rune(" .--..--..--..--..--..--..--..--..--..--..--..--..--..--..--. "))
	ByeDisplay = append(ByeDisplay, []rune("/ .. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\"))
	ByeDisplay = append(ByeDisplay, []rune("\\ \\/\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ \\/ /"))
	ByeDisplay = append(ByeDisplay, []rune(" \\/ /`--'`--'`--'`--'`--'`--'`--'`--'`--'`--'`--'`--'`--'\\/ / "))
	ByeDisplay = append(ByeDisplay, []rune(" / /\\                                                    / /\\ "))
	ByeDisplay = append(ByeDisplay, []rune("/ /\\ \\   ____                 _   _                  _  / /\\ \\"))
	ByeDisplay = append(ByeDisplay, []rune("\\ \\/ /  / ___| ___   ___   __| | | |__  _   _  ___  | | \\ \\/ /"))
	ByeDisplay = append(ByeDisplay, []rune(" \\/ /  | |  _ / _ \\ / _ \\ / _` | | '_ \\| | | |/ _ \\ | |  \\/ / "))
	ByeDisplay = append(ByeDisplay, []rune(" / /\\  | |_| | (_) | (_) | (_| | | |_) | |_| |  __/ |_|  / /\\ "))
	ByeDisplay = append(ByeDisplay, []rune("/ /\\ \\  \\____|\\___/ \\___/ \\__,_| |_.__/ \\__, |\\___| (_) / /\\ \\"))
	ByeDisplay = append(ByeDisplay, []rune("\\ \\/ /                                  |___/           \\ \\/ /"))
	ByeDisplay = append(ByeDisplay, []rune(" \\/ /                                                    \\/ / "))
	ByeDisplay = append(ByeDisplay, []rune(" / /\\.--..--..--..--..--..--..--..--..--..--..--..--..--./ /\\ "))
	ByeDisplay = append(ByeDisplay, []rune("/ /\\ \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\.. \\/\\ \\"))
	ByeDisplay = append(ByeDisplay, []rune("\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `'\\ `' /"))
	ByeDisplay = append(ByeDisplay, []rune(" `--'`--'`--'`--'`--'`--'`--'`--'`--'`--'`--'`--'`--'`--'`--' "))

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
	CreateDisplayVariables()
	time.Sleep(time.Second * 1)
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
//  /\┃/\      (Équipement)│┃ Coup de poing humain │ 240 │ 100 │ Attaque │ 105 ┃ /\┃/\
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
