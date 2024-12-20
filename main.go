package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mrcampbell/pokemon-golang/pkg/env"
	"github.com/mrcampbell/pokemon-golang/pkg/file"
	"github.com/mrcampbell/pokemon-golang/pkg/game"
	"github.com/mrcampbell/pokemon-golang/pkg/http"
	"github.com/mrcampbell/pokemon-golang/pkg/sqlc"
	"github.com/mrcampbell/pokemon-golang/pokeapi/language"
	"github.com/mrcampbell/pokemon-golang/pokeapi/version"
)

const defaultVersion = version.Emerald
const defaultLanguage = language.English

func main() {
	ctx := context.Background()
	env.LoadEnv()

	conn, err := pgx.Connect(ctx, env.POSTGRES_URL)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)
	pokemon, err := queries.ListPokemon(ctx)
	if err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		panic(err)
	}
	fmt.Println(pokemon)

	speciesService := file.NewSpeciesService(defaultVersion)
	moveService := file.NewMoveService(defaultLanguage)
	pokemonService := game.NewPokemonService(moveService, speciesService)

	server := http.NewServer(&pokemonService, &speciesService, &moveService)
	server.Run(":8080")

	// cache := cache.NewInMemoryCache()
	// _ = cache

}
