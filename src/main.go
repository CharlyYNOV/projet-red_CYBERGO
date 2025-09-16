package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strings"
)

type Item struct {
	Name        string
	Description string
}

type Character struct {
	name      string
	class     string
	level     int
	HPmax     int
	HPnormal  int
	Inventory []Item
	Skills	  []Skill
	Coins     int
}

type Enemy struct {
    name     string
    HPmax    int
    HPnormal int
    attack   int
	rewardLevel int
    rewardCoins int
    loot     Item
}

type Skill struct {
	name 	string
	description string
	shortcut     string
}

func main() {
	AfficherTitre()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Vous venez d'apparaitre au coeur d'un univers futuriste, à la fois au style cyberpunk et dystopique. La Terre est transformée par les progrès technologiques et les inégalités sociales, au sein de la mégalopole Cyber City. Dans cet univers, règnent le chaos et les technologies de pointe, la pauvreté et les disparités dominent, les toutes-puissantes mégacorporations imposent leur loi, supplantant les gouvernements traditionnels")
	fmt.Println("Le jeu propose un personnage principal déjà réglé, mais il est possible de créer le vôtre !")
	fmt.Println("Bienvenue dans CyberGo !")
	fmt.Println("=== Création du personnage ===")
	fmt.Println("Voulez-vous jouer avec le personnage principal ? (O / N)")
	fmt.Print("> ")

	var joueur Character

	if scanner.Scan() {
		choix := scanner.Text()
		if choix == "O" {
			fmt.Println()
			joueur = PersonnagePrincipal()
		} else if choix == "N" {
			fmt.Println()
			joueur = InitCharacter()
		} else {
			fmt.Println("Choix invalide, nous vous sélectionnons Ash par défaut.")
			joueur = PersonnagePrincipal()
		}
	}

	fmt.Println("\n=== Début de l'aventure ===")
	fmt.Printf("Bienvenue, %s le %s !\n", joueur.name, joueur.class)

	AfficherMenu(&joueur)
}

func AfficherTitre() {
	fmt.Println(`
				   █████████             █████                          █████████          
				  ███░░░░░███           ░░███                          ███░░░░░███         
 				 ███     ░░░  █████ ████ ░███████   ██████  ████████  ███     ░░░   ██████ 
				░███         ░░███ ░███  ░███░░███ ███░░███░░███░░███░███          ███░░███
				░███          ░███ ░███  ░███ ░███░███████  ░███ ░░░ ░███    █████░███ ░███
				░░███     ███ ░███ ░███  ░███ ░███░███░░░   ░███     ░░███  ░░███ ░███ ░███
				 ░░█████████  ░░███████  ████████ ░░██████  █████     ░░█████████ ░░██████ 
				  ░░░░░░░░░    ░░░░░███ ░░░░░░░░   ░░░░░░  ░░░░░       ░░░░░░░░░   ░░░░░░  
				               ███ ░███                                                    
				              ░░██████                                                     
				                ░░░░░░      
`)
}

func AfficherMenu(joueur *Character) {
	fmt.Println(`
	- Jouer (P)
	- Quitter (Q)
`)
	fmt.Println("\nChoisissez la destination...")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		choix := scanner.Text()
		if choix == "P" {
			fmt.Println()
			fmt.Println("Lancement du jeu...")
			fmt.Println()
			Jouer(joueur)
		} else if choix == "Q" {
			fmt.Println("Merci d'avoir joué, à bientôt !")
			os.Exit(0)
		} else {
			fmt.Println()
			fmt.Println("/!\\ Choix invalide, retour au menu. /!\\")
			AfficherMenu(joueur)
		} 
	}
}

func Jouer(joueur *Character) {
	fmt.Println("◊ Combats à mort (F)")
	fmt.Println("◊ Marchand (M)")
	fmt.Println("◊ Inventaire (I)")
	fmt.Println("◊ Détails du personnage (D)")
	fmt.Println("◊ Retour (B)")

	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		choix := scanner.Text()
		if choix == "I" {
			fmt.Println()
			accessInventory(joueur)
			fmt.Println()
		} else if choix == "B" {
			fmt.Println()
			AfficherMenu(joueur)
			fmt.Print("> ")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		} else if choix == "D" {
			fmt.Println()
			DisplayInfo(*joueur)
			fmt.Println()
		} else if choix == "M" {
			fmt.Println()
			Marchand(*joueur)
			fmt.Println()
		} else if choix == "F" {
			fmt.Println()
			ChoixCombat(*joueur)
			fmt.Println()
		} else {
			fmt.Println()
			fmt.Println("/!\\ Choix invalide, retour au menu. /!\\")
			fmt.Println()
			Jouer(joueur)
		}
	}
}

func PersonnagePrincipal() Character {
	personnage := Character{
		name:     "Ash",
		class:    "Mercenaire",
		level:    1,
		HPmax:    200,
		HPnormal: 100,
		Coins : 15,
		Inventory: []Item{
			{"Seringue de soin", "Restaure 50 PV"},
			{"Fiole de Neurotoxine", "Empoisonne l'ennemi pendant 3s (Combat uniquement)", },
		},
		Skills: []Skill{
			{"Coup de poing", "Enlève 15 PV", "(C)"},
		},
	}
	DisplayFirstInfo(personnage)
	return personnage
}

func InitCharacter() Character {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Choisis un prénom: ")
	scanner.Scan()
	name := scanner.Text()

	var class string
	var HPmax int

	for {
		fmt.Print("Choisis une classe (Humain, Robot, Vagabond) : ")
		scanner.Scan()
		class = strings.Title(strings.ToLower(scanner.Text()))

		switch class {
		case "Humain":
			HPmax = 100
		case "Robot":
			HPmax = 120
		case "Vagabond":
			HPmax = 80
		default:
			fmt.Println("Classe invalide. Essaie encore.")
			continue
		}
		break
	}

	personnage := Character{
		name:     name,
		class:    class,
		level:    1,
		HPmax:    HPmax,
		HPnormal: HPmax / 2,
		Coins:    15,

		Inventory: []Item{
			{"Seringue de soin", "Restaure 50 PV"},
			{"Fiole de Neurotoxine", "Empoisonne l'ennemi pendant 3s"},
		},
		Skills: []Skill{
			{"Coup de poing", "Enlève 15 PV", "(C)"},
		},
	}

	fmt.Println()
	fmt.Println("Ton personnage a été créé !")
	fmt.Println()
	DisplayFirstInfo(personnage)

	return personnage
}

func DisplayFirstInfo(joueur Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom :", joueur.name)
	fmt.Println("Classe :", joueur.class)
	fmt.Println("Niveau :", joueur.level)
	fmt.Println("Points de vie :", joueur.HPnormal, "/", joueur.HPmax)
	fmt.Println("Pièces :", joueur.Coins)
	fmt.Println("Inventaire :")
	for _, item := range joueur.Inventory {
		fmt.Printf("- %s : %s\n", item.Name, item.Description)
	}
	fmt.Println("===============================")
}

func DisplayInfo(joueur Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom :", joueur.name)
	fmt.Println("Classe :", joueur.class)
	fmt.Println("Niveau :", joueur.level)
	fmt.Println("Points de vie :", joueur.HPnormal, "/", joueur.HPmax)
	fmt.Println("Pièces :", joueur.Coins)
	fmt.Println("Inventaire :")
	for _, item := range joueur.Inventory {
		fmt.Printf("- %s : %s\n", item.Name, item.Description)
	}
	fmt.Println("===============================")
	GoBack(joueur)
}

func accessInventory(joueur *Character) {
	fmt.Println("\n===== Inventaire de", joueur.name, "=====")

	if len(joueur.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		GoBack(*joueur)
	}

	for i, item := range joueur.Inventory {
		fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Description)
	}
	fmt.Print("\nChoisissez un numéro d'objet ou revenir au menu principal (B): ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        choix := scanner.Text()
		if choix == "B" {
			fmt.Println()
			Jouer(joueur)
			fmt.Println()
		}
        var index int
        _, err := fmt.Sscanf(choix, "%d", &index)
		
        if err != nil {
            fmt.Println("X Choix invalide")
            accessInventory(joueur)
        }
        index = index - 1
        if index < 0 || index >= len(joueur.Inventory) {
            fmt.Println("X Pas d’objet à cet index")
            accessInventory(joueur)
        }
		
        item := joueur.Inventory[index]
        if item.Name == "Seringue de soin" {
            joueur.HPnormal += 50
            if joueur.HPnormal > joueur.HPmax {
                joueur.HPnormal = joueur.HPmax
            }
            fmt.Println("Vous avez utilisé une seringue pour vous soigner ! PV :", joueur.HPnormal, "/", joueur.HPmax)
			joueur.Inventory = removeItem(joueur.Inventory, index)
        } else {
            fmt.Println("/!\\ Cet objet n'est pas utilisable./!\\")
        }
		var ennemi Enemy
		if item.Name == "Fiole de Neurotoxine" {
			Poison(*joueur, ennemi)
			fmt.Println("Vous avez utilisé une fiole contre votre ennemi !")
		}
    }
	GoBack(*joueur)
}

func accessInventoryCombat(joueur *Character, ennemi *Enemy) {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Printf("\n===== Inventaire de %s =====\n", joueur.name)
    if len(joueur.Inventory) == 0 {
        fmt.Println("Votre inventaire est vide.")
        return
    }
    for i, item := range joueur.Inventory {
        fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Description)
    }

    fmt.Print("\nChoisissez un numéro d'objet ou revenir au combat (B): ")
    if scanner.Scan() {
        choix := scanner.Text()
        if choix == "B" || choix == "b" {
            return
        }
        var index int
        _, err := fmt.Sscanf(choix, "%d", &index)
        if err != nil || index < 1 || index > len(joueur.Inventory) {
            fmt.Println("X Choix invalide")
            return
        }

        item := joueur.Inventory[index-1]

        if item.Name == "Seringue de soin" {
            joueur.HPnormal += 50
            if joueur.HPnormal > joueur.HPmax {
                joueur.HPnormal = joueur.HPmax
            }
            fmt.Printf("%s utilise une seringue et récupère 50 PV ! (%d/%d)\n", joueur.name, joueur.HPnormal, joueur.HPmax)
            joueur.Inventory = removeItem(joueur.Inventory, index-1)
			return

        } else if item.Name == "Fiole de Neurotoxine" {
            fmt.Printf("%s utilise une fiole de Neurotoxine contre %s !\n", joueur.name, ennemi.name)
            for i := 1; i <= 3; i++ {
                time.Sleep(1 * time.Second)
                ennemi.HPnormal -= 10
                if ennemi.HPnormal < 0 {
                    ennemi.HPnormal = 0
                }
                fmt.Printf("Tour %d : %s perd 10 PV à cause du poison. PV restants : %d/%d\n", i, ennemi.name, ennemi.HPnormal, ennemi.HPmax)
                if ennemi.HPnormal <= 0 {
                    IsDeadEnemy(*ennemi)
                    break
                }
            }
            joueur.Inventory = removeItem(joueur.Inventory, index-1)
			return
        } else {
            fmt.Println("/!\\ Cet objet n'est pas utilisable en combat. /!\\")
        }
    }
}

func removeItem(inventory []Item, index int) []Item {

	if index < 0 || index >= len(inventory) {
		return inventory
	}
	return append(inventory[:index], inventory[index+1:]...)
}

func Marchand(joueur Character) {
	fmt.Println(`
	(\_._/) 
    	( o o )   
    	(  -  ) 
       c("===")ɔ
    	 || || 
    	 || || 
        ==   ==
	`)
	fmt.Println("Bienvenue,", joueur.name, "! J'ai tout ce qu'il te faut ici.")
	fmt.Println("Tu as", joueur.Coins, "pièces.")

	fmt.Println("\n--- Objets ---")
	fmt.Println("- Seringue de soin (3 pièces) O1 / N")
	fmt.Println("- Fiole de Neurotoxine (3 pièces) O2 / N")

	fmt.Println()
	fmt.Println("- Coup de poing (gratuit, déjà appris) O3 / N")
	fmt.Println("- Bras mécanique (25 pièces + Pièces mécaniques) O4 / N")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	if scanner.Scan() {
		choix := scanner.Text()
		if choix == "O1" {
			if joueur.Coins >= 3 {
				item := Item{Name: "Seringue de soin", Description: "Restaure 50 PV"}
				joueur.Inventory = addItem(joueur.Inventory, item)
				joueur.Coins -= 3
				fmt.Println("\nTu as acheté une seringue, sois prudent !\n")
			} else {
				fmt.Println("\nTu n'as pas assez de pièces pour acheter une seringue, reviens vite !")
			}
		} else if choix == "N" {
			fmt.Println()
			fmt.Println("\nÀ très bientot", joueur.name, "!")
			fmt.Println()
			Jouer(&joueur)
		} else if choix == "O2" {
			if joueur.Coins >= 3 {
				item := Item{Name: "Fiole de Neurotoxine", Description: "Empoisonne l'ennemi pendant 3s"}
				joueur.Inventory = addItem(joueur.Inventory, item)
				joueur.Coins -= 3
				fmt.Println("\nJe me demande bien ce que tu comptes en faire.\n")
			} else {
				fmt.Println("\nTu n'as pas assez de pièces pour acheter une fiole toxique, reviens vite !")
			}
		} else if choix == "O3" {
			fmt.Println("\nTu maîtrises déjà le Coup de poing, mais maintenant il est utilisable en combat.\n")
		} else if choix == "O4" {
			if joueur.Coins >= 25 && hasItem(joueur.Inventory, "Pièces mécaniques") {
				skill := Skill{"Uppercut chromé", "Enlève 40 PV avec ton bras cybernétique", "(U)"}
				joueur.Skills = append(joueur.Skills, skill)
				joueur.Coins -= 25
				joueur.HPnormal += 30
				joueur.HPmax += 50
				fmt.Println("\nTu as débloqué une nouvelle compétence : Bras mécanique !\n")
			} else {
				fmt.Println("\nTu n’as pas les conditions nécessaires (25 pièces + Pièces mécaniques).\n")
			}
		}
		Jouer(&joueur)
	}
}

func hasItem(inventory []Item, itemName string) bool{
	for _, item := range inventory {
		if item.Name == itemName {
			return true
		}
	}
	return false
}

func hasSkill(skills []Skill, skillName string) bool {
    for _, s := range skills {
        if s.name == skillName {
            return true
        }
    }
    return false
}

func Poison(joueur Character, ennemi Enemy) {
	fmt.Printf("%s utilise une fiole de Neurotoxine contre %s !\n", joueur.name, ennemi.name)

	for i:= 1; i < 3; i++ {
		time.Sleep(1 * time.Second)
		damage := 10
		ennemi.HPnormal -= damage
		fmt.Printf("Tour %d : %s perd %d PV à cause du poison. PV restants : %d/%d\n", i, ennemi.name, damage, ennemi.HPnormal, ennemi.HPmax)

		if ennemi.HPnormal <= 0 {
			IsDeadEnemy(ennemi)
			return
		}
	}
	fmt.Printf("L'effet du poison sur %s est terminé. \n", ennemi.name)
}

func LearnSkill(joueur Character) {
    newSkills := []Skill{
        {"Coup de poing", "Enlève 15 PV", "(C)"},
        {"Uppercut chromé", "Enlève 40 PV", "(U)"},
    }
    joueur.Skills = append(joueur.Skills, newSkills...)
    fmt.Println("Vous avez appris de nouvelles compétences :")
    for _, skill := range newSkills {
        fmt.Println("-", skill.name)
    }
}

func addItem(inventory []Item, item Item) []Item {
    if !canAddItem(inventory) {
        fmt.Printf("Impossible d'ajouter %s : inventaire plein.\n", item.Name)
        return inventory
    }
    return append(inventory, item)
}

func IsDeadJoueur(joueur Character, ennemi Character) {
    if joueur.HPnormal <= 0 {
        fmt.Println(`
				 ▄█     █▄     ▄████████    ▄████████     ███        ▄████████ ████████▄  
				███     ███   ███    ███   ███    ███ ▀█████████▄   ███    ███ ███   ▀███ 
				███     ███   ███    ███   ███    █▀     ▀███▀▀██   ███    █▀  ███    ███ 
				███     ███   ███    ███   ███            ███   ▀  ▄███▄▄▄     ███    ███ 
				███     ███ ▀███████████ ▀███████████     ███     ▀▀███▀▀▀     ███    ███ 
				███     ███   ███    ███          ███     ███       ███    █▄  ███    ███ 
				███ ▄█▄ ███   ███    ███    ▄█    ███     ███       ███    ███ ███   ▄███ 
 				▀███▀███▀    ███    █▀   ▄████████▀     ▄████▀     ██████████ ████████▀  
   `)
        fmt.Printf("\n %s est mort...\n", joueur.name)
        joueur.HPnormal = joueur.HPmax / 2
        fmt.Printf(" %s a été ressuscité avec %d PV !\n", joueur.name, joueur.HPnormal)
	}
	Jouer(&joueur)
}

func IsDeadEnemy(ennemi Enemy) {
	if ennemi.HPnormal <= 0 {
        fmt.Println(`
				 ____  __ _  ____  _  _  _  _      ____  ____   __   ____ 
				(  __)(  ( \(  __)( \/ )( \/ )    (    \(  __) / _\ (    \
				 ) _) /    / ) _) / \/ \ )  /      ) D ( ) _) /    \ ) D (
				(____)\_)__)(____)\_)(_/(__/      (____/(____)\_/\_/(____/                          
	`)
        fmt.Printf("\n %s est mort...\n", ennemi.name)
    }
}

func CreateEnemies() (Enemy, Enemy) {
    brasDroit := Enemy{
        name:        "Bras droit de Garmadon",
        HPmax:       120,
        HPnormal:    120,
        attack:      20,
        rewardLevel: 2,
        rewardCoins: 20,
        loot:        Item{"Pièces mécaniques", "Composants pour débloquer des équipements"},
    }

    garmadon := Enemy{
        name:        "Garmadon",
        HPmax:       200,
        HPnormal:    200,
        attack:      30,
        rewardLevel: 3,
        rewardCoins: 30,
        loot:        Item{"", ""},
    }

    return brasDroit, garmadon
}

func ChoixCombat(joueur Character) {
    brasDroit, garmadon := CreateEnemies()

    fmt.Println("\nChoisis ton adversaire :")
    fmt.Println("1 - Bras droit de Garmadon (Récompense : +2 niveaux, +20 pièces, Pièces mécaniques)")
    fmt.Println("2 - Garmadon (Niveau requis : 5. Récompense : +3 niveaux, +30 pièces)")
    fmt.Print("> ")

    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        choix := scanner.Text()
        if choix == "1" {
            Combat(&joueur, &brasDroit)
        } else if choix == "2" {
            if joueur.level >= 5 {
				fmt.Println()
				fmt.Print(`
*******#@@@@@%#********#@@@@@@*+++++++++
********%@@@@@%********#@@@@@#++++++++++
********#%@%%%#*#@@@@#**#@%%@*++++++++++
*********#%%%%**%@@@@%**#%%%#+++++++++++
**********%%%%@@@@%%@@@@%%%#+++=++++++++
********#**#%%%%@@@@@%%%##*=+#%*=+++++++
******#%@@%#%%%%%%######%%%%%%%%#=-+++++
*****%@@@@@@@%%%@%%%%%%%%%@@@@@%#+++++++
+******%%@@@@@%%%%%%%%%%%%@@@@%*++++++++
++******%@@@@@%%#+*##+*%%%%@@%%%#+++++==
+++++*#%%%%%%%%%%%%%%%#*%%%%%%%%%%*+====
++++++*%%%%%%@%%%#%%%%%%%%%*+++++=======
+++++++++++++*%%@@%##%@%%#+=========-===
***++++++++*+++*%@@@%@@%#**#**####%%*===
+++*##**#%%@@@%%%%#%%%%%%%%%%%%%%#*+====
*++++**#%@@@@@@%%%%%%%%%%%%%%%%*----====
**+++*+#%@@@@@@%%%%%%%%%%%%%@%%%+::--===
***%@@@@%@@@@@@%%%%%%#%@@@%##%%%##=::-==
#*%@@@@@@@@@@@@%%@%%%%@@@@@%*%%%#%%*--==
##@@@@@@@@@@@@@%%%%%%%@@@@@%##%%%%%#+-==
######%%#%@@@@@%%%%%%%@@@@@@%*##%@@%%+==
####****#@@@@@@@%@%%%%@@@@@@@%@@@%@@%+=+
#####***#@@@@@@@%@%@%%@@@@@@@@@%##%%#*#=
########%@@@@@@%%%%%%%@@@@@@@@@@@@%*++*%
			`)
                Combat(&joueur, &garmadon)
            } else {
                fmt.Println("X Ton niveau est trop faible pour affronter Garmadon (niveau requis : 5).")
            }
        } else {
            fmt.Println("X Choix invalide.")
            ChoixCombat(joueur)
        }
    }
	GoBack(joueur)
}

func Combat(joueur *Character, ennemi *Enemy) {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Printf("\nCombat : %s VS %s\n", joueur.name, ennemi.name)

    for joueur.HPnormal > 0 && ennemi.HPnormal > 0 {
        fmt.Printf("\n%s : %d/%d PV | %s : %d/%d PV\n",
            joueur.name, joueur.HPnormal, joueur.HPmax,
            ennemi.name, ennemi.HPnormal, ennemi.HPmax)

        fmt.Println("\nChoisis ton action :")
        fmt.Println("1 - Coup de poing (-15 PV)")

        if hasSkill(joueur.Skills, "Uppercut chromé") {
            fmt.Println("2 - Uppercut chromé (-40 PV)")
        }

        fmt.Println("3 - Inventaire")
        fmt.Print("> ")

        if scanner.Scan() {
            choix := scanner.Text()
            if choix == "1" {
                ennemi.HPnormal -= 15
                fmt.Printf("%s frappe %s ! (-15 PV)\n", joueur.name, ennemi.name)
            } else if choix == "2" && hasSkill(joueur.Skills, "Uppercut chromé") {
                ennemi.HPnormal -= 40
                fmt.Printf("%s utilise Uppercut chromé sur %s ! (-40 PV)\n", joueur.name, ennemi.name)
            } else if choix == "3" {
                accessInventoryCombat(joueur, ennemi)
                if ennemi.HPnormal <= 0 {
					IsDeadEnemy(*ennemi)
					joueur.level += ennemi.rewardLevel
					joueur.Coins += ennemi.rewardCoins
					if ennemi.loot.Name !="" {
						joueur.Inventory = append(joueur.Inventory, ennemi.loot)
					}
					fmt.Println("Bravo", joueur.name, "tu as vaincu", ennemi.name, ", tu obtiens donc" ,ennemi.rewardLevel, "de niveaux et" ,ennemi.rewardCoins, " de pièces !")
					if ennemi.loot.Name != "" {
            			fmt.Printf("💎 Tu obtiens : %s\n", ennemi.loot.Name)
					}
					break
            	} else {
                fmt.Println("X Choix invalide.")
                continue
            	}
        	}
        	if ennemi.HPnormal <= 0 {
            	IsDeadEnemy(*ennemi)
            	joueur.level += ennemi.rewardLevel
           		joueur.Coins += ennemi.rewardCoins
            	if ennemi.loot.Name != "" {
                	joueur.Inventory = append(joueur.Inventory, ennemi.loot)
            	}
            	fmt.Printf(" Tu gagnes %d niveaux et %d pièces !\n", ennemi.rewardLevel, ennemi.rewardCoins)
            	if ennemi.loot.Name != "" {
                fmt.Printf(" Tu obtiens : %s\n", ennemi.loot.Name)
           		}
            	break
        	}

        	fmt.Printf("%s attaque %s ! (-%d PV)\n", ennemi.name, joueur.name, ennemi.attack)
        	joueur.HPnormal -= ennemi.attack

        	if joueur.HPnormal <= 0 {
            IsDeadJoueur(*joueur, *joueur)
        	}
    	}
	}
}

func GoBack(joueur Character) {

	fmt.Println("\nRevenir au menu principal (B).")
	fmt.Println()
	fmt.Print("> ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println()
	Jouer(&joueur)
}

func canAddItem(inventory []Item) bool {
    return len(inventory) < 10
}

func AddItem(inventory []Item, item Item) []Item {
    if !canAddItem(inventory) {
        fmt.Printf("Impossible d'ajouter %s : inventaire plein.\n", item.Name)
        return inventory
    }
    return append(inventory, item)
}

func TestAchat() {
    var joueur Character
    joueur.Coins = 10
    joueur.Inventory = []Item{}

    if joueur.Coins >= 3 {
        item := Item{Name: "Seringue de soin", Description: "Restaure 50 PV"}
        if canAddItem(joueur.Inventory) {
            joueur.Inventory = AddItem(joueur.Inventory, item)
            joueur.Coins -= 3
            fmt.Println("\nTu as acheté une seringue, sois prudent !\n")
        } else {
            fmt.Println("\nTon inventaire est plein, tu ne peux pas acheter cette seringue.")
        }
    } else {
        fmt.Println("\nTu n'as pas assez de pièces pour acheter une seringue.")
    }

    if joueur.Coins >= 3 {
        item := Item{Name: "Fiole de Neurotoxine", Description: "Empoisonne l'ennemi pendant 3s"}
        if canAddItem(joueur.Inventory) {
            joueur.Inventory = AddItem(joueur.Inventory, item)
            joueur.Coins -= 3
            fmt.Println("\nTu as acheté une fiole de Neurotoxine, sois prudent !\n")
        } else {
            fmt.Println("\nTon inventaire est plein, tu ne peux pas acheter cette fiole de Neurotoxine.")
        }
    	} else {
        fmt.Println("\nTu n'as pas assez de pièces pour acheter une fiole de Neurotoxine.")
    }
}