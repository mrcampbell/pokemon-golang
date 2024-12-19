package main

import (
	"fmt"

	"github.com/mrcampbell/pokemon-golang/pkg/file"
	"github.com/mrcampbell/pokemon-golang/pkg/game"
	"github.com/mrcampbell/pokemon-golang/pokeapi/language"
	"github.com/mrcampbell/pokemon-golang/pokeapi/version"
)

const defaultVersion = version.Emerald
const defaultLanguage = language.English

func main() {
	fmt.Println("Hello, World!")
	speciesService := file.NewSpeciesService(defaultVersion)

	moveService := file.NewMoveService(defaultLanguage)

	pokemonService := game.NewPokemonService(moveService, speciesService)

	pikachu := pokemonService.CreatePokemon(25, 25)
	fmt.Println(pikachu.PrintableSummary())
	// cache := cache.NewInMemoryCache()
	// _ = cache

}
