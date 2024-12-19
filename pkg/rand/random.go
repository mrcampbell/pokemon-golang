package rand

import "math/rand/v2"

func RandBetween(min, max int) int {
	return rand.IntN(max-min) + min
}
