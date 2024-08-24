package types

type Monster struct {
	Name         string  `json:"name"`
	Code         CodeName  `json:"code"`
	Level        int     `json:"level"`
	HP           int     `json:"hp"`
	AttackFire   int     `json:"attack_fire"`
	AttackEarth  int     `json:"attack_earth"`
	AttackWater  int     `json:"attack_water"`
	AttackAir    int     `json:"attack_air"`
	ResFire      int     `json:"res_fire"`
	ResEarth     int     `json:"res_earth"`
	ResWater     int     `json:"res_water"`
	ResAir       int     `json:"res_air"`
	MinGold      int     `json:"min_gold"`
	MaxGold      int     `json:"max_gold"`
	Drops        []Drops  `json:"drops"`
}

type MonsterList struct {
  ServerCodeInfo
  Data []Monster `json:"data"`
  Total int `json:"total"`
  Page int `json:"page"`
  Size int `json:"size"`
  Pages int `json:"pages"`
}

type MonsterRequest struct {
  Code CodeName `json:"code"`
}

type MonsterListRequest struct {
  Drop string `json:"drop"`
  MaxLevel int `json:"max_level"`
  MinLevel int `json:"min_level"`
  Page int `json:"page"`
  Size int `json:"size"`
}
