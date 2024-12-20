package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kr/pretty"
	"github.com/mrcampbell/pokemon-golang/pkg/db"
	"github.com/mrcampbell/pokemon-golang/pkg/env"
	"github.com/mrcampbell/pokemon-golang/pkg/file"
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

	speciesService := file.NewSpeciesService(defaultVersion)
	moveService := file.NewMoveService(defaultLanguage)
	pokemonService := db.NewPokemonService(conn, queries, moveService, speciesService)

	// pikachu := pokemonService.CreatePokemon(25, 50)
	// print(pikachu.PrintableSummary())
	// _, err = pokemonService.SavePokemon(ctx, pikachu)
	// if err != nil {
	// 	fmt.Printf("Error saving pokemon: %v\n", err)
	// }

	id, err := uuid.Parse("19aaa6ec-b51f-4042-86a0-ffc8b64c19dc")
	if err != nil {
		fmt.Printf("Error parsing UUID: %v\n", err)
	}

	p, err := pokemonService.GetPokemon(ctx, id)
	if err != nil {
		fmt.Printf("Error getting pokemon: %v\n", err)
	}
	pretty.Println(p)

	// server := http.NewServer(&pokemonService, &speciesService, &moveService)
	// server.Run(":8080")

	// cache := cache.NewInMemoryCache()
	// _ = cache
}
