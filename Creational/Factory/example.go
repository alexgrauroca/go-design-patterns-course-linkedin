package main

import "fmt"

func main() {
	mag1, _ := newPublication(PubTypeMagazine, "Tyme", 50, "The Tymes")
	mag2, _ := newPublication(PubTypeMagazine, "Lyfe", 40, "Lyfe Inc")
	news1, _ := newPublication(PubTypeNewspaper, "The Herald", 60, "Heralders")
	news2, _ := newPublication(PubTypeNewspaper, "The Standard", 30, "Standarders")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub iPublication) {
	fmt.Println("----------------------")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Println("----------------------")
}
