package element

type Element string

const (
	None     Element = "none"
	Normal   Element = "normal"
	Fire     Element = "fire"
	Water    Element = "water"
	Electric Element = "electric"
	Grass    Element = "grass"
	Ice      Element = "ice"
	Fighting Element = "fighting"
	Poison   Element = "poison"
	Ground   Element = "ground"
	Flying   Element = "flying"
	Psychic  Element = "psychic"
	Bug      Element = "bug"
	Rock     Element = "rock"
	Ghost    Element = "ghost"
	Fairy    Element = "fairy"
	Dark     Element = "dark"
	Steel    Element = "steel"
)

func ElementFromString(s string) Element {
	switch s {
	case "none":
		return None
	case "normal":
		return Normal
	case "fire":
		return Fire
	case "water":
		return Water
	case "electric":
		return Electric
	case "grass":
		return Grass
	case "ice":
		return Ice
	case "fighting":
		return Fighting
	case "poison":
		return Poison
	case "ground":
		return Ground
	case "flying":
		return Flying
	case "psychic":
		return Psychic
	case "bug":
		return Bug
	case "rock":
		return Rock
	case "ghost":
		return Ghost
	case "fairy":
		return Fairy
	case "dark":
		return Dark
	case "steel":
		return Steel
	default:
		return ""
	}
}
