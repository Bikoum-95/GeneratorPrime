package generateprime

import (
    "flag"
    "fmt"
    "os"
)

func GPrime() {
	name := flag.String("name", "", "Nom du pirate")
	prime := flag.String("prime", "", "Prime du pirate")
	img := flag.String("img", "", "Photo du pirate")
	flag.Parse()

	if *name == "" || *prime == "" || *img == "" {
		fmt.Println("Tous les paramètres sont obligatoires")
		os.Exit(1)
	}
    fmt.Println("Pirate enregistré avec succès.")
}

