package main

import (
	"fmt"

	"github.com/jackc/pgx"
	"github.com/mrcampbell/pokemon-golang/pkg/env"
	"github.com/mrcampbell/pokemon-golang/pkg/file"
	"github.com/mrcampbell/pokemon-golang/pkg/game"
	"github.com/mrcampbell/pokemon-golang/pkg/http"
	"github.com/mrcampbell/pokemon-golang/pokeapi/language"
	"github.com/mrcampbell/pokemon-golang/pokeapi/version"
)

const defaultVersion = version.Emerald
const defaultLanguage = language.English

func main() {
	// ctx := context.Background()
	env.LoadEnv()

	conn, err := pgx.Connect(pgx.ConnConfig{
		Database: env.POSTGRES_DB,
		Host:     env.POSTGRES_HOST,
		User:     env.POSTGRES_USER,
		Password: env.POSTGRES_PASSWORD,
		Port:     env.POSTGRES_PORT,
	})
	if err != nil {
		fmt.Errorf("Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer conn.Close()

	speciesService := file.NewSpeciesService(defaultVersion)
	moveService := file.NewMoveService(defaultLanguage)
	pokemonService := game.NewPokemonService(moveService, speciesService)

	server := http.NewServer(&pokemonService, &speciesService, &moveService)
	server.Run(":8080")

	// cache := cache.NewInMemoryCache()
	// _ = cache

}
