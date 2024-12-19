package versiongroup

type VersionGroup string

const (
	RedBlue                      VersionGroup = "red-blue"
	Yellow                       VersionGroup = "yellow"
	GoldSilver                   VersionGroup = "gold-silver"
	Crystal                      VersionGroup = "crystal"
	RubySapphire                 VersionGroup = "ruby-sapphire"
	Emerald                      VersionGroup = "emerald"
	FireRedLeafGreen             VersionGroup = "fire-red-leaf-green"
	DiamondPearl                 VersionGroup = "diamond-pearl"
	Platinum                     VersionGroup = "platinum"
	HeartGoldSoulSilver          VersionGroup = "heart-gold-soul-silver"
	BlackWhite                   VersionGroup = "black-white"
	Colosseum                    VersionGroup = "colosseum"
	XD                           VersionGroup = "xd"
	Black2White2                 VersionGroup = "black-2-white-2"
	XY                           VersionGroup = "x-y"
	OmegaRubyAlphaSapphire       VersionGroup = "omega-ruby-alpha-sapphire"
	SunMoon                      VersionGroup = "sun-moon"
	UltraSunUltraMoon            VersionGroup = "ultra-sun-ultra-moon"
	LetsGoPikachuLetsGoEevee     VersionGroup = "lets-go-pikachu-lets-go-eevee"
	SwordShield                  VersionGroup = "sword-shield"
	TheIsleOfArmor               VersionGroup = "the-isle-of-armor"
	TheCrownTundra               VersionGroup = "the-crown-tundra"
	BrilliantDiamondShiningPearl VersionGroup = "brilliant-diamond-shining-pearl"
	LegendsArceus                VersionGroup = "legends-arceus"
	ScarletViolet                VersionGroup = "scarlet-violet"
	TheTealMask                  VersionGroup = "the-teal-mask"
	TheIndigoDisk                VersionGroup = "the-indigo-disk"
)

func FromString(s string) VersionGroup {
	switch s {
	case "red-blue":
		return RedBlue
	case "yellow":
		return Yellow
	case "gold-silver":
		return GoldSilver
	case "crystal":
		return Crystal
	case "ruby-sapphire":
		return RubySapphire
	case "emerald":
		return Emerald
	case "fire-red-leaf-green":
		return FireRedLeafGreen
	case "diamond-pearl":
		return DiamondPearl
	case "platinum":
		return Platinum
	case "heart-gold-soul-silver":
		return HeartGoldSoulSilver
	case "black-white":
		return BlackWhite
	case "colosseum":
		return Colosseum
	case "xd":
		return XD
	case "black-2-white-2":
		return Black2White2
	case "x-y":
		return XY
	case "omega-ruby-alpha-sapphire":
		return OmegaRubyAlphaSapphire
	case "sun-moon":
		return SunMoon
	case "ultra-sun-ultra-moon":
		return UltraSunUltraMoon
	case "lets-go-pikachu-lets-go-eevee":
		return LetsGoPikachuLetsGoEevee
	case "sword-shield":
		return SwordShield
	case "the-isle-of-armor":
		return TheIsleOfArmor
	case "the-crown-tundra":
		return TheCrownTundra
	case "brilliant-diamond-shining-pearl":
		return BrilliantDiamondShiningPearl
	case "legends-arceus":
		return LegendsArceus
	case "scarlet-violet":
		return ScarletViolet
	case "the-teal-mask":
		return TheTealMask
	case "the-indigo-disk":
		return TheIndigoDisk
	}
	return RedBlue
}
