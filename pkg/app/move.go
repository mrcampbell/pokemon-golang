package app

import "fmt"

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

type MoveService interface {
	GetMove(id int) Move
}

func (m Move) PrintableSummary() string {
	result := fmt.Sprintf("Name: %s\n", m.Name)
	result += fmt.Sprintf("\tPower: %d\n", m.Power)
	result += fmt.Sprintf("\tPP: %d\n", m.PP)
	result += fmt.Sprintf("\tAccuracy: %d\n", m.Accuracy)
	result += fmt.Sprintf("\tCategory: %s\n", m.Category)
	result += fmt.Sprintf("\tDamage Class: %s\n", m.DamageClass)
	result += fmt.Sprintf("\tPriority: %d\n", m.Priority)
	result += fmt.Sprintf("\tEffect Chance: %d\n", m.EffectChance)
	result += fmt.Sprintf("\tEffect: %s\n", m.EffectText)
	result += fmt.Sprintf("\tAilment: %s\n", m.Ailment)
	result += fmt.Sprintf("\tAilment Chance: %d\n", m.AilmentChance)
	result += fmt.Sprintf("\tCritical Rate: %d\n", m.CriticalRate)
	result += fmt.Sprintf("\tDrain: %d\n", m.Drain)
	result += fmt.Sprintf("\tFlinch Chance: %d\n", m.FlinchChance)
	result += fmt.Sprintf("\tHealing: %d\n", m.Healing)
	result += fmt.Sprintf("\tMax Hits: %d\n", m.MaxHits)
	result += fmt.Sprintf("\tMax Turns: %d\n", m.MaxTurns)
	result += fmt.Sprintf("\tMin Hits: %d\n", m.MinHits)
	result += fmt.Sprintf("\tMin Turns: %d\n", m.MinTurns)
	result += fmt.Sprintf("\tStat Chance: %d\n", m.StatChance)
	result += fmt.Sprintf("\tAccuracy: %d\n", m.Accuracy)
	result += fmt.Sprintf("\tCategory: %s\n", m.Category)
	result += fmt.Sprintf("\tDamage Class: %s\n", m.DamageClass)
	result += fmt.Sprintf("\tPriority: %d\n", m.Priority)
	result += fmt.Sprintf("\tEffect Chance: %d\n", m.EffectChance)
	result += fmt.Sprintf("\tEffect: %s\n", m.EffectText)
	result += fmt.Sprintf("\tAilment: %s\n", m.Ailment)
	result += fmt.Sprintf("\tAilment Chance: %d\n", m.AilmentChance)
	result += fmt.Sprintf("\tCritical Rate: %d\n", m.CriticalRate)
	result += fmt.Sprintf("\tDrain: %d\n", m.Drain)
	result += fmt.Sprintf("\tFlinch Chance: %d\n", m.FlinchChance)
	result += fmt.Sprintf("\tHealing: %d\n", m.Healing)
	result += fmt.Sprintf("\tMax Hits: %d\n", m.MaxHits)
	result += fmt.Sprintf("\tMax Turns: %d\n", m.MaxTurns)
	result += fmt.Sprintf("\tMin Hits: %d\n", m.MinHits)
	result += fmt.Sprintf("\tMin Turns: %d\n", m.MinTurns)
	result += fmt.Sprintf("\tStat Chance: %d\n", m.StatChance)
	result += "\t\tStat Changes: \n"
	for _, change := range m.StatChanges {
		result += fmt.Sprintf("\t\t\t%d to %s\n", change.Change, change.Stat)
	}
	return result

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

type LearnableMove struct {
	MoveID int
	Level  int
	Method MoveLearnMethod
}
