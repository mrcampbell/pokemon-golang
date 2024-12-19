package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mrcampbell/pokemon-golang/pokeapi"
)

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

// todo: move to service, just storing here for simplicity and speed of development
func LoadMove(id int) Move {
	source, err := readFile(fmt.Sprintf("%d", id))
	if err != nil {
		panic(err)
	}

	result := Move{
		ID:   source.ID,
		Name: source.Name,
	}

	return result
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

func readFile(path string) (pokeapi.Move, error) {
	const path_dir = "data/moves/"
	data, err := os.ReadFile(path_dir + path + ".json")
	if err != nil {
		return pokeapi.Move{}, err
	}

	pa_move := pokeapi.Move{}
	err = json.Unmarshal(data, &pa_move)
	if err != nil {
		return pokeapi.Move{}, err
	}

	return pa_move, nil
}
