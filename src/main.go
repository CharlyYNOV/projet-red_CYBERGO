package main

import (
	"bufio"
	"fmt"
	"os"
)

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
			joueur = PersonnagePrincipal()
		} else if choix == "N" {
			joueur = InitCharacter()
		} else {
			fmt.Println("Choix invalide, nous vous sélectionnons Ash par défaut.")
			joueur = PersonnagePrincipal()
		}
	}

	fmt.Println("\n=== Début de l'aventure ===")
	fmt.Printf("Bienvenue, %s le %s !\n", joueur.name, joueur.class)

	AfficherMenu(joueur)
}


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
}

type Skill struct {
	name 	string
	description string
	shortcut     string
}

func AfficherTitre() {
	fmt.Println(`
  _____   __     __   ____     _____    ____       ____       ___
 / ____|  \ \   / /  | __ \   |  ___|  |  _ \     / ___|     / _ \
| |        \ \_/ /   |    /   | |___   | |_) |   | /        | | | |
| |         \   /    |  _ \   |  ___|  |  _  /   | |  ___   | | | |
| |____      | |     | |_) |  | |___   | | \ \   | |_|  _|  | |_| |
 \_____|     |_|     |____/   |_____|  |_|  \_\   \____|     \___/
`)
}

func AfficherMenu(joueur Character) {
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

func Jouer(joueur Character) {
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
			DisplayInfo(joueur)
			fmt.Println()
		} else if choix == "M" {
			fmt.Println()
			Marchand(joueur)
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
		Inventory: []Item{
			{"Seringue de soin", "Restaure 50 PV"},
		},
		Skills: []Skill{
			{"Coup de poing", "Enlève 15 PV", "(C)"},
			{"Poing chromé", "Enlève 40 PV", "(U)"},
		},
	}
	DisplayFirstInfo(personnage)
	return personnage
}

func InitCharacter() Character {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Choisis un prénom: ")
	scanner.Scan()

	personnage := Character{
		name:     scanner.Text(),
		class:    "Robot",
		level:    1,
		HPmax:    100,
		HPnormal: 40,
		Inventory: []Item{
			{"Seringue de soin", "Restaure 50 PV", },
			{"Fiole de Neurotoxine", "Empoisonne l'ennemi pendant 3s", },
		},
		Skills: []Skill{
			{"Coup de poing", "Enlève 15 PV", "(C)"},
			{"Poing chromé", "Enlève 40 PV", "(U)"},
		},
	}
	fmt.Println("Ton personnage a été créé !")
	DisplayFirstInfo(personnage)

	return personnage
}

func DisplayFirstInfo(joueur Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom :", joueur.name)
	fmt.Println("Classe :", joueur.class)
	fmt.Println("Niveau :", joueur.level)
	fmt.Println("Points de vie :", joueur.HPnormal, "/", joueur.HPmax)
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
	fmt.Println("Inventaire :")
	for _, item := range joueur.Inventory {
		fmt.Printf("- %s : %s\n", item.Name, item.Description)
	}
	fmt.Println("===============================")
	GoBack(joueur)
}

func accessInventory(joueur Character) {
	fmt.Println("\n===== Inventaire de", joueur.name, "=====")

	if len(joueur.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		GoBack(joueur)
	}

	for i, item := range joueur.Inventory {
		fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Description)
	}
	fmt.Print("\nChoisissez un numéro d’objet à utiliser : ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        choix := scanner.Text()

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
            fmt.Printf("Vous avez utilisé une seringue pour vous soigner ! PV :%d / %d", joueur.HPnormal, joueur.HPmax)

            joueur.Inventory = removeItem(joueur.Inventory, index)
        } else {
            fmt.Println("/!\\ Cet objet n’est pas utilisable./!\\")
        }
    }
	GoBack(joueur)
}

func removeItem(inventory []Item, index int) []Item {

	if index < 0 || index >= len(inventory) {
		return inventory
	}
	newInventory := append(inventory[:index], inventory[index+1:]...)
	fmt.Println(newInventory)

	return newInventory
}

func GoBack(joueur Character) {

	fmt.Println("\nRevenir au menu principal (B).")
	fmt.Println()
	fmt.Print("> ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println()
	Jouer(joueur)
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
	fmt.Println("- Seringue de soin (gratuit) O1 / N1")
	GoBack(joueur)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		choix := scanner.Text()

		if choix == "O1" {
			item := Item{Name: "Seringue de soin", Description: "Restaure 50 PV"}
			fmt.Println("\nTu as acheté une seringue, sois prudent !\n")
			joueur.Inventory = addItem(joueur.Inventory, item)
			Jouer(joueur)
		} else if choix == "N1" {
			GoBack(joueur)
		}
	}
}

func Poison(joueur Character) {

}

func LearnSkill(joueur Character) {
	
}

func addItem(inventory []Item, item Item) []Item {
	return append(inventory, item)
}

func isDead(joueur *Character) {
    if joueur.HPnormal <= 0 {
        fmt.Println(`
W       W   AAAAAAA   SSSSSSS  TTTTTTTT   EEEEEEE   DDDDDD  
W       W   A     A   S           T       E         D     D 
W   W   W   AAAAAAA    SSSSS      T       EEEEE     D     D 
W W W W W   A     A         S     T       E         D     D 
W W   W W   A     A   S     S     T       E         D     D 
W       W   A     A    SSSSS      T       EEEEEEE   DDDDDD  
   `)
        fmt.Printf("\n %s est mort...\n", joueur.name)
        joueur.HPnormal = joueur.HPmax / 2
        fmt.Printf(" %s a été ressuscité avec %d PV !\n", joueur.name, joueur.HPnormal)
    }
}
