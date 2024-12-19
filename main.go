package main

import (
	"fmt"

	"github.com/mrcampbell/pokemon-golang/pkg/app"
	"github.com/mrcampbell/pokemon-golang/pkg/cache"
)

func main() {
	fmt.Println("Hello, World!")
	species := app.SpeciesFromFile(1)
	fmt.Println(species)

	cache := cache.NewInMemoryCache()

	cache.Set("key", "value")
	value, _ := cache.Get("key")
	fmt.Println(value)
}
