package main

import (
	"fmt"
	"os"
)

/* 
j'ai téléchargé le fichier ou il y a la police standard a adopté lors du traitement
si nous parvenons a bien trimer les parties concernées et à les replacer avec leur ascii correspondant 
on trouvera l'exo in sha Allah
-j'espère que ce bout de code vous donnera un aperçu sur ce qui nous attend
 */  


func main() {
	// storing the args
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error: review your arguments")
	} else {
		//Reading the file with the graphics
		file, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Error while reading the file:", err)
		} else {
			// storing the graphics in a variable
			longtext :=string(file)
			// Trimming graphics
			fmt.Println(longtext[:39]) // exclamaton mark
			fmt.Println(longtext[41:96]) // double quote mark
			fmt.Println(longtext[98:193]) // hash mark

		}
	}
}
