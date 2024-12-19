package app

import "github.com/mrcampbell/pokemon-golang/pokeapi"

type MoveLearnMethod string

const (
	MoveLearnMethodLevelUp MoveLearnMethod = "level-up"
	MoveLearnMethodMachine MoveLearnMethod = "machine"
	MoveLearnMethodTutor   MoveLearnMethod = "tutor"
	MoveLearnMethodEgg     MoveLearnMethod = "egg"
)

type Move struct {
	ID   int
	Name string
}

type LearnableMove struct {
	MoveID int
	Level  int
	Method MoveLearnMethod
}

func LearnableMovesFromSource(source []pokeapi.LearnableMove) []LearnableMove {
	result := make([]LearnableMove, len(source))
	for i, move := range source {
		result[i] = LearnableMove{
			MoveID: move.MoveID,
			Level:  move.Level,
			Method: MoveLearnMethod(move.Method),
		}
	}
	return result
}
