package parse

import (
	"MyGeneratorPrime/pirates"
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCSV(filepath string) []pirates.Pirate {
	file, err := os.Open(filepath)
	var pr []pirates.Pirate

	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Print(err)
	}
	for _, record := range records {
		if len(record) < 3 {
			fmt.Println("Ligne non trouvée")
			continue
		}
		name := record[0]
		prime := record[1]
		img := record[2]

		pr = append(pr, pirates.Pirate{Name: name, Prime: prime, Img: img})
		if err != nil {
			fmt.Println("Erreur tototo :", err)
			continue
		}
	}
	fmt.Println("Pirate ajouté :", pr)
	return pr
}
