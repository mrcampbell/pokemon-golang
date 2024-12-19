package generation

type Generation string

const (
	Generation1 Generation = "generation-i"
	Generation2 Generation = "generation-ii"
	Generation3 Generation = "generation-iii"
	Generation4 Generation = "generation-iv"
	Generation5 Generation = "generation-v"
	Generation6 Generation = "generation-vi"
	Generation7 Generation = "generation-vii"
	Generation8 Generation = "generation-viii"
	Generation9 Generation = "generation-ix"
)

func FromString(s string) Generation {
	switch s {
	case "generation-i":
		return Generation1
	case "generation-ii":
		return Generation2
	case "generation-iii":
		return Generation3
	case "generation-iv":
		return Generation4
	case "generation-v":
		return Generation5
	case "generation-vi":
		return Generation6
	case "generation-vii":
		return Generation7
	case "generation-viii":
		return Generation8
	case "generation-ix":
		return Generation9
	}
	return Generation1
}

func FromInt(i int) Generation {
	switch i {
	case 1:
		return Generation1
	case 2:
		return Generation2
	case 3:
		return Generation3
	case 4:
		return Generation4
	case 5:
		return Generation5
	case 6:
		return Generation6
	case 7:
		return Generation7
	case 8:
		return Generation8
	case 9:
		return Generation9
	}
	return Generation1
}
