package app

type MoveService interface {
	GetMove(id int) Move
}

type MoveLearnMethod string

const (
	MoveLearnMethodLevelUp MoveLearnMethod = "level-up"
	MoveLearnMethodMachine MoveLearnMethod = "machine"
	MoveLearnMethodTutor   MoveLearnMethod = "tutor"
	MoveLearnMethodEgg     MoveLearnMethod = "egg"
)

type MoveDamageClass string

const (
	MoveDamageClassStatus   MoveDamageClass = "status"
	MoveDamageClassPhysical MoveDamageClass = "physical"
	MoveDamageClassSpecial  MoveDamageClass = "special"
)

type MoveCategory string

const (
	MoveCategoryDamage           MoveCategory = "damage"
	MoveCategoryAilment          MoveCategory = "ailment"
	MoveCategoryNetGoodStats     MoveCategory = "net-good-stats"
	MoveCategoryHeal             MoveCategory = "heal"
	MoveCategoryDamageAndAilment MoveCategory = "damage+ailment"
	MoveCategorySwagger          MoveCategory = "swagger"
	MoveCategoryDamageAndLower   MoveCategory = "damage+lower"
	MoveCategoryDamageAndRaise   MoveCategory = "damage+raise"
	MoveCategoryDamageAndHeal    MoveCategory = "damage+heal"
	MoveCategoryOhko             MoveCategory = "ohko"
	MoveCategoryWholeFieldEffect MoveCategory = "whole-field-effect"
	MoveCategoryFieldEffect      MoveCategory = "field-effect"
	MoveCategoryForceSwitch      MoveCategory = "force-switch"
	MoveCategoryUnique           MoveCategory = "unique"
)

type Stat string

const (
	StatHP         Stat = "hp"
	StatAttack     Stat = "attack"
	StatDefense    Stat = "defense"
	StatSpecialAtk Stat = "special-attack"
	StatSpecialDef Stat = "special-defense"
	StatSpeed      Stat = "speed"
	StatAccuracy   Stat = "accuracy"
	StatEvasion    Stat = "evasion"
)

type StatChange struct {
	Change int
	Stat   Stat
}

type Move struct {
	ID            int
	Name          string
	Accuracy      int
	DamageClass   MoveDamageClass
	EffectChance  int
	EffectText    string
	FlavorText    string
	Ailment       Ailment
	AilmentChance int
	Category      MoveCategory
	CriticalRate  int
	Drain         int
	FlinchChance  int
	Healing       int
	MaxHits       int
	MaxTurns      int
	MinHits       int
	MinTurns      int
	StatChance    int
	Power         int
	PP            int
	Priority      int
	StatChanges   []StatChange
}

type LearnableMove struct {
	MoveID int
	Level  int
	Method MoveLearnMethod
}
