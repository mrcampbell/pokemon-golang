package main

import (
	"fmt"

	"github.com/mrcampbell/pokemon-golang/app"
)

func main() {
	fmt.Println("Hello, World!")
	species := app.SpeciesFromFile(1)
	fmt.Println(species)
}
