package app

type Stats struct {
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	SpAtk   int `json:"special-attack"`
	SpDef   int `json:"special-defense"`
	Speed   int `json:"speed"`
}

type BattleStats struct {
	HP       int `json:"hp"`
	Attack   int `json:"attack"`
	Defense  int `json:"defense"`
	SpAtk    int `json:"special-attack"`
	SpDef    int `json:"special-defense"`
	Speed    int `json:"speed"`
	Accuracy int `json:"accuracy"`
	Evasion  int `json:"evasion"`
}
