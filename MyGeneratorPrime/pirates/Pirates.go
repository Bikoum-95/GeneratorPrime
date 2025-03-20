package pirates

import (
	"errors"
	"strings"
)

type Pirate struct {
	Name  string
	Prime string
	Img   string
}

func New(name, prime, img string) (*Pirate, error) {
	pirate := &Pirate{
		Name:  name,
		Prime: prime,
		Img:   img,
	}

	if name != strings.ToUpper(name) {
		return nil, errors.New("le nom du pirate doit être en majuscules")
	}

	return pirate, nil
}
