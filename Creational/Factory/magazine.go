package main

import "fmt"

type magazine struct {
	publication
}

func (n magazine) String() string {
	return fmt.Sprintf("This is a magazine named %s", n.name)
}

func createMagazine(name string, pages int, publisher string) iPublication {
	return &magazine{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
