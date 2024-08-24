package types

import "time"

type FightResult string

const (
  FightResultWin = FightResult("win")
  FightResultLose = FightResult("lose")
)

type MoveRequest struct {
  X int `json:"x"`
  Y int `json:"y"`
}

type Cooldown struct {
  TotalSeconds int `json:"total_seconds"`
  RemainingSeconds int `json:"remaining_seconds"`
  StartedAt time.Time `json:"started_at"`
  Expiration time.Time `json:"expiration"`
  Reason string `json:"reason"`
}

type Content struct {
  Type MapContentType `json:"type"`
  Code CodeName  `json:"code"`
}

type Destination struct {
  Name string `json:"name"`
  Skin string `json:"skin"`
  X int `json:"x"`
  Y int `json:"y"`
  Content Content `json:"content"`
}

// Details contains experience points and items.
type Details struct {
	Xp    int     `json:"xp"`
	Items []SimpleItem  `json:"items"`
}

// SimpleItem represents an individual item with a code and quantity.
type SimpleItem struct {
	Code     CodeName `json:"code"`
	Quantity int    `json:"quantity"`
}

// Character represents the character's stats and inventory.
type Character struct {
	Name                      string    `json:"name"`
	Skin                      string    `json:"skin"`
	Level                     int       `json:"level"`
	Xp                        int       `json:"xp"`
	MaxXp                     int       `json:"max_xp"`
	TotalXp                   int       `json:"total_xp"`
	Gold                      int       `json:"gold"`
	Speed                     int       `json:"speed"`
	MiningLevel               int       `json:"mining_level"`
	MiningXp                  int       `json:"mining_xp"`
	MiningMaxXp               int       `json:"mining_max_xp"`
	WoodcuttingLevel           int       `json:"woodcutting_level"`
	WoodcuttingXp              int       `json:"woodcutting_xp"`
	WoodcuttingMaxXp           int       `json:"woodcutting_max_xp"`
	FishingLevel              int       `json:"fishing_level"`
	FishingXp                 int       `json:"fishing_xp"`
	FishingMaxXp              int       `json:"fishing_max_xp"`
	WeaponcraftingLevel       int       `json:"weaponcrafting_level"`
	WeaponcraftingXp          int       `json:"weaponcrafting_xp"`
	WeaponcraftingMaxXp       int       `json:"weaponcrafting_max_xp"`
	GearcraftingLevel         int       `json:"gearcrafting_level"`
	GearcraftingXp            int       `json:"gearcrafting_xp"`
	GearcraftingMaxXp         int       `json:"gearcrafting_max_xp"`
	JewelrycraftingLevel      int       `json:"jewelrycrafting_level"`
	JewelrycraftingXp         int       `json:"jewelrycrafting_xp"`
	JewelrycraftingMaxXp      int       `json:"jewelrycrafting_max_xp"`
	CookingLevel              int       `json:"cooking_level"`
	CookingXp                 int       `json:"cooking_xp"`
	CookingMaxXp              int       `json:"cooking_max_xp"`
	Hp                        int       `json:"hp"`
	Haste                     int       `json:"haste"`
	CriticalStrike            int       `json:"critical_strike"`
	Stamina                   int       `json:"stamina"`
	AttackFire                int       `json:"attack_fire"`
	AttackEarth               int       `json:"attack_earth"`
	AttackWater               int       `json:"attack_water"`
	AttackAir                 int       `json:"attack_air"`
	DmgFire                   int       `json:"dmg_fire"`
	DmgEarth                  int       `json:"dmg_earth"`
	DmgWater                  int       `json:"dmg_water"`
	DmgAir                    int       `json:"dmg_air"`
	ResFire                   int       `json:"res_fire"`
	ResEarth                  int       `json:"res_earth"`
	ResWater                  int       `json:"res_water"`
	ResAir                    int       `json:"res_air"`
	X                        int       `json:"x"`
	Y                        int       `json:"y"`
	Cooldown                  int       `json:"cooldown"`
	CooldownExpiration        time.Time `json:"cooldown_expiration"`
	WeaponSlot                string    `json:"weapon_slot"`
	ShieldSlot                string    `json:"shield_slot"`
	HelmetSlot                string    `json:"helmet_slot"`
	BodyArmorSlot             string    `json:"body_armor_slot"`
	LegArmorSlot              string    `json:"leg_armor_slot"`
	BootsSlot                 string    `json:"boots_slot"`
	Ring1Slot                 string    `json:"ring1_slot"`
	Ring2Slot                 string    `json:"ring2_slot"`
	AmuletSlot                string    `json:"amulet_slot"`
	Artifact1Slot             string    `json:"artifact1_slot"`
	Artifact2Slot             string    `json:"artifact2_slot"`
	Artifact3Slot             string    `json:"artifact3_slot"`
	Consumable1Slot           string    `json:"consumable1_slot"`
	Consumable1SlotQuantity   int       `json:"consumable1_slot_quantity"`
	Consumable2Slot           string    `json:"consumable2_slot"`
	Consumable2SlotQuantity   int       `json:"consumable2_slot_quantity"`
	Task                      string    `json:"task"`
	TaskType                  string    `json:"task_type"`
	TaskProgress              int       `json:"task_progress"`
	TaskTotal                 int       `json:"task_total"`
	InventoryMaxItems         int       `json:"inventory_max_items"`
	Inventory                 []InventoryItem `json:"inventory"`
}

// InventoryItem represents an item in the character's inventory.
type InventoryItem struct {
	Slot     int    `json:"slot"`
  Code     CodeName `json:"code"`
	Quantity int    `json:"quantity"`
}

// Fight represents the structure of a fight result.
type Fight struct {
	Xp                  int                `json:"xp"`
	Gold                int                `json:"gold"`
	Drops               []SimpleItem             `json:"drops"`
	Turns               int                `json:"turns"`
	MonsterBlockedHits  BlockedHits        `json:"monster_blocked_hits"`
	PlayerBlockedHits   BlockedHits        `json:"player_blocked_hits"`
	Logs                []string           `json:"logs"`
	Result              FightResult        `json:"result"`
}

// BlockedHits contains hit information for a specific type of attack.
type BlockedHits struct {
	Fire  int `json:"fire"`
	Earth int `json:"earth"`
	Water int `json:"water"`
	Air   int `json:"air"`
	Total int `json:"total"`
}

// Item represents the structure of an item with its details and crafting information.
type Item struct {
	Name        string   `json:"name"`
	Code        CodeName   `json:"code"`
	Level       int      `json:"level"`
	Type        string   `json:"type"`
	Subtype     string   `json:"subtype"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
	Craft       Craft    `json:"craft"`
}

// Effect represents the effect of an item, with a name and value.
type Effect struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Craft represents crafting information for an item, including required skill, level, and materials.
type Craft struct {
	Skill     string  `json:"skill"`
	Level     int     `json:"level"`
	Items     []SimpleItem `json:"items"`
	Quantity  int     `json:"quantity"`
}

type MoveResponse struct {
  ServerCodeInfo
  Cooldown Cooldown `json:"cooldown"`
  Destination Destination `json:"destination"`
  Character Character `json:"character"`
}

type GatheringResponse struct {
  ServerCodeInfo
	Cooldown    Cooldown    `json:"cooldown"`
	Details     Details     `json:"details"`
	Character   Character   `json:"character"`
}

type FightResponse struct {
  ServerCodeInfo
	Cooldown    Cooldown    `json:"cooldown"`
	Fight       Fight       `json:"fight"`
	Character   Character   `json:"character"`
}

type BankResponse struct {
  ServerCodeInfo
	Cooldown    Cooldown    `json:"cooldown"`
  Item       Item `json:"item"`
	Bank       []SimpleItem       `json:"bank"`
	Character   Character   `json:"character"`
}

type CharacterResponse struct {
  ServerCodeInfo
  Data Character `json:"data"`
}
