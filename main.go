package main

import (
	"fmt"

	"github.com/mrcampbell/pokemon-golang/pkg/app"
	"github.com/mrcampbell/pokemon-golang/pkg/cache"
)

func main() {
	fmt.Println("Hello, World!")
	// species := app.LoadSpecies(1)
	// fmt.Println(species)
	move := app.LoadMove(28)
	fmt.Println(move)

	// for later use
	_ = cache.NewInMemoryCache()

}
