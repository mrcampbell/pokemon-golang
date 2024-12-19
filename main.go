package main

import (
	"fmt"

	"github.com/mrcampbell/pokemon-golang/pkg/app"
	"github.com/mrcampbell/pokemon-golang/pkg/cache"
)

func main() {
	fmt.Println("Hello, World!")
	species := app.Load(1)
	fmt.Println(species)

	// for later use
	_ = cache.NewInMemoryCache()

}
