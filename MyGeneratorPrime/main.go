package main

import (
	test "MyGeneratorPrime/generate_prime"
	"MyGeneratorPrime/parse"
	"MyGeneratorPrime/pdf"
	"MyGeneratorPrime/pirates"
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		return
	}

	switch args[0] {
	case "-cli":
		if len(args) != 7 {
			log.Fatal("Erreur veuillez entrer un argument. ")
			return
		}
		if args[1] == "-name" && args[3] == "-prime" && args[5] == "-img" {
			name := args[2]
			prime := args[4]
			img := args[6]

			pirate, err := pirates.New(name, prime, img)
			pdf.CreatePdf(
				pdf.PdfSaver{OutputDir: "/OutputDir"},
				pdf.Pirate{Name: pirate.Name, Prime: pirate.Prime, Img: pirate.Img},
			)

			if err != nil {
				return
			}
		} else {
			fmt.Println("veuillez entrer -cli -prime -name -img")
			return
		}

	case "-file":

		filepath := args[1]
		pirates := parse.ParseCSV(filepath)

		for _, pirate := range pirates {
			pdf.CreatePdf(
				pdf.PdfSaver{OutputDir: "/OutputDir" + pirate.Name},
				pdf.Pirate{Name: pirate.Name, Prime: pirate.Prime, Img: pirate.Img},
			)
		}
	default:
		fmt.Println("Veuillez utiliser -cli ou -file")
	}
	test.GPrime()
}
