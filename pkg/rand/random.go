package rand

import (
	"math/rand/v2"

	"github.com/mrcampbell/pokemon-golang/pkg/app"
)

func RandBetween(min, max int) int {
	return rand.IntN(max-min) + min
}

// todo: use generics
func ShuffleLearnableMoves(slice []app.LearnableMove) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
