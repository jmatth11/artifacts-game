package types

type SkillEnum string

const (
  SkillMining = SkillEnum("mining")
  SkillWoodcutting = SkillEnum("woodcutting")
  SkillFishing = SkillEnum("fishing")
)

type Drops struct {
  Code CodeName `json:"code"`
  Rate int `json:"rate"`
  MinQuantity int `json:"min_quantity"`
  MaxQuantity int `json:"max_quantity"`
}

type Resource struct {
  Name string `json:"name"`
  Code CodeName `json:"code"`
  Skill string `json:"skill"`
  Level int `json:"level"`
  Drops []Drops `json:"drops"`
}

type ResourceList struct {
  ServerCodeInfo
  Data []Resource `json:"data"`
  Total int `json:"total"`
  Page int `json:"page"`
  Size int `json:"size"`
  Pages int `json:"pages"`
}

type ResourceRequest struct {
  Drop string `json:"drop"`
  MaxLevel int `json:"max_level"`
  MinLevel int `json:"min_level"`
  Page int `json:"page"`
  Size int `json:"size"`
  Skill SkillEnum `json:"skill"`
}
