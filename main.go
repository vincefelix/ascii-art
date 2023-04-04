package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//--Stocker les arguments dans une variable
	args := os.Args[1:]

	//--vérifier le nombre d'arguments
	if len(args) != 1 {
		fmt.Println("Error: review your arguments")

		//--vérifier si l'argument est vide
	} else if args[0] ==""{
		return
	}else {

		 //--------------------------------------1er étape: Lire le fichier avec les graphiques------------------------------------------------//
		//                                       -----------------------------------------------                                              //
		
		file, err := os.Open("standard.txt")
		if err != nil {
			fmt.Println("Erreur: nous ne parvenons pas à lire le fichier source standard", err)
			} else {
			// Stocker les caractères des graphiques dans une variable
			longtext := bufio.NewScanner(file)
			
			// Mettre les données stockées dans un tableau de string
			var tab []string
			for longtext.Scan() {
				tab = append(tab, longtext.Text())
			}
			
			 //--------------------------2ème étape: stocker chaque ensemble de caractère pour chaque ascii dans un slice de string---------------//
			//                           ------------------------------------------------------------------------------------------              //
			var vinc [][]string
			for i := 1; i < len(tab); i += 9 {
				vinc = append(vinc, tab[i:i+8])
			}
 //-----------------------------------------------------3ème étape: gérer l'affichage------------------------------------------------------//
//                                                      ------------------------------                                                    //

//--vérifier si l'argument contient un caractère affichable

if !IsPrintable(args[0]) { // l'argument ne contient pas de caractère affichable
fmt.Print(args[0])

} else { // l'argument contient un caractère affichable

splitext := strings.Split(args[0], "\\n") // séparer le string en cas de présence d'un "newline"s

var num int //varibale pour déterminer l'index dans le tableau des caractères 

space := []string{"      ", "      ", "      ", "      ", "      ", "      ", "      ", "      "} //representation du caractère 32 (espace)

//--Afficher le string de l'argument sous le format ascii-art
for _, v := range splitext {
	
	//--récolter les caractères asci-art à afficher
					var result [][]string 
					for _, y := range v {
						if y <0 || y > 127 {
							break
						} else {
						num = int(y - 32) //la position correspondant au caractère selon le tableau de caractères dans vinc
						if y == 32 {
							result = append(result, space)
							} else {
								result = append(result, vinc[num])
								}
							}
						}
						printres(result) // affiche la version graphique des caractères récoltées
						}
					}
				}
			}
		}
		
		//-----------------------------------------------------------Fonctions utilisées-------------------------------------------------------//
	   //                                                           ---------------------                                                     //
		
		//printres gére l'affichage d'un tableau à 2 dimensions contenant des caractères
		func printres(result [][]string) {
			if result == nil {
				return
			} else {
	var temp string // stocker les cartères à afficher par ligne
	for i := 0; i < 8; i++ { // parcourir la colonne
		for j := 0; j < len(result); j++ { // parcourir la ligne
			temp += result[j][i]
		}
		if i != 7 { // ne pas ajouter de newline au dernier caractère
			temp += "\n"
		}
	}
	fmt.Println(temp) //affiche la ligne
}
}

/*
IsPrintable permet de vérifier si le string comporte un caractère affichable ou pas
-il renvoie true s'il rencontre un seul caractère affichable 
-et renvoie false s'il n'en voit aucun.
*/
func IsPrintable(s string) bool {
	a := []rune(s) //convertir le tableau string en tableau de runes
	b := len(s)
	rep:= true
	for i := 0; i <= b-1; i++ {
		if a[i] <0 || a[i]>127 {
			fmt.Println("Attention le caractère saisi n'est pas compris dans l'ASCII")
			break
		}else if a[i] < 32 || a[i] > 126 { //intervalle des caractères affichable
			rep = false
		} else {
			rep=true
			break
		}
	}
	return rep
}
