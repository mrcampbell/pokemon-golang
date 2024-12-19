package version

type Version string

const (
	Red              Version = "red"
	Blue             Version = "blue"
	Yellow           Version = "yellow"
	Gold             Version = "gold"
	Silver           Version = "silver"
	Crystal          Version = "crystal"
	Ruby             Version = "ruby"
	Sapphire         Version = "sapphire"
	Emerald          Version = "emerald"
	FireRed          Version = "firered"
	LeafGreen        Version = "leafgreen"
	Diamond          Version = "diamond"
	Pearl            Version = "pearl"
	Platinum         Version = "platinum"
	HeartGold        Version = "heartgold"
	SoulSilver       Version = "soulsilver"
	Black            Version = "black"
	White            Version = "white"
	Colosseum        Version = "colosseum"
	XD               Version = "xd"
	Black2           Version = "black-2"
	White2           Version = "white-2"
	X                Version = "x"
	Y                Version = "y"
	OmegaRuby        Version = "omega-ruby"
	AlphaSapphire    Version = "alpha-sapphire"
	Sun              Version = "sun"
	Moon             Version = "moon"
	UltraSun         Version = "ultra-sun"
	UltraMoon        Version = "ultra-moon"
	LetsGoPikachu    Version = "lets-go-pikachu"
	LetsGoEevee      Version = "lets-go-eevee"
	Sword            Version = "sword"
	Shield           Version = "shield"
	TheIsleOfArmor   Version = "the-isle-of-armor"
	TheCrownTundra   Version = "the-crown-tundra"
	BrilliantDiamond Version = "brilliant-diamond"
	ShiningPearl     Version = "shining-pearl"
	LegendsArceus    Version = "legends-arceus"
	Scarlet          Version = "scarlet"
	Violet           Version = "violet"
	TheTealMask      Version = "the-teal-mask"
	TheIndigoDisk    Version = "the-indigo-disk"
)

func FromString(s string) Version {
	switch s {
	case "red":
		return Red
	case "blue":
		return Blue
	case "yellow":
		return Yellow
	case "gold":
		return Gold
	case "silver":
		return Silver
	case "crystal":
		return Crystal
	case "ruby":
		return Ruby
	case "sapphire":
		return Sapphire
	case "emerald":
		return Emerald
	case "firered":
		return FireRed
	case "leafgreen":
		return LeafGreen
	case "diamond":
		return Diamond
	case "pearl":
		return Pearl
	case "platinum":
		return Platinum
	case "heartgold":
		return HeartGold
	case "soulsilver":
		return SoulSilver
	case "black":
		return Black
	case "white":
		return White
	case "colosseum":
		return Colosseum
	case "xd":
		return XD

	case "black-2":
		return Black2
	case "white-2":
		return White2
	case "x":
		return X
	case "y":
		return Y
	case "omega-ruby":
		return OmegaRuby
	case "alpha-sapphire":
		return AlphaSapphire
	case "sun":
		return Sun
	case "moon":
		return Moon
	case "ultra-sun":
		return UltraSun
	case "ultra-moon":
		return UltraMoon
	case "lets-go-pikachu":
		return LetsGoPikachu
	case "lets-go-eevee":
		return LetsGoEevee
	case "sword":
		return Sword
	case "shield":
		return Shield
	case "the-isle-of-armor":
		return TheIsleOfArmor
	case "the-crown-tundra":
		return TheCrownTundra
	case "brilliant-diamond":
		return BrilliantDiamond
	case "shining-pearl":
		return ShiningPearl
	case "legends-arceus":
		return LegendsArceus
	case "scarlet":
		return Scarlet
	case "violet":
		return Violet
	case "the-teal-mask":
		return TheTealMask
	case "the-indigo-disk":
		return TheIndigoDisk
	}
	return Red
}
