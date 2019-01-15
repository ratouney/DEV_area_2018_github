package yolopack

import (
	"fmt"
)

// Logique : Oui, lol Go veux un commentaire
// et il doit commencer avec le nom de la fonction
func Logique() {
	fmt.Printf("Hello there\n")
}

// Ah oui, et que les fonctions/variables qui commencent
// avec une majuscule sont exportés
func logique() {
	fmt.Printf("Same but different\n")
}
