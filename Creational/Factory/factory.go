package main

import "fmt"

const (
	PubTypeMagazine  = "magazine"
	PubTypeNewspaper = "newspaper"
)

func newPublication(pubType, name string, pg int, pub string) (iPublication, error) {
	switch pubType {
	case "newspaper":
		return createNewspaper(name, pg, pub), nil
	case "magazine":
		return createMagazine(name, pg, pub), nil
	}

	return nil, fmt.Errorf("no such publication type: %s", pubType)
}
