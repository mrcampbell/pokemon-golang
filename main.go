package main

import (
	"fmt"

	"github.com/mrcampbell/pokemon-golang/pkg/cache"
	"github.com/mrcampbell/pokemon-golang/pkg/file"
	"github.com/mrcampbell/pokemon-golang/pokeapi/language"
	"github.com/mrcampbell/pokemon-golang/pokeapi/version"
)

const defaultVersion = version.Emerald
const defaultLanguage = language.English

func main() {
	fmt.Println("Hello, World!")
	speciesService := file.NewSpeciesService(defaultVersion)
	pikachu := speciesService.GetSpecies(25)
	fmt.Println(pikachu)

	moveService := file.NewMoveService(defaultLanguage)
	thunderbolt := moveService.GetMove(85)
	fmt.Println(thunderbolt)

	cache := cache.NewInMemoryCache()
	_ = moveService
	_ = cache

}
