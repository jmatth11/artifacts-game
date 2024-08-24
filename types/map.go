package types

type MapContentType string
type CodeName string

const (
  MonsterContentType = MapContentType("monster")
  ResourceContentType = MapContentType("resource")
  WorkshopContentType = MapContentType("workshop")
  BankContentType = MapContentType("bank")
  GrandExchangeContentType = MapContentType("grand_exchange")
  TasksMasterContentType = MapContentType("tasks_master")
)

const (
  AshTree = CodeName("ash_tree")
  CopperRocks = CodeName("copper_rocks")
  CoalRocks = CodeName("coal_rocks")
  Chicken = CodeName("chicken")
  BlueSlime = CodeName("blue_slime")
  RedSlime = CodeName("red_slime")
)

type MapGetAllRequest struct {
  ContentCode CodeName `json:"content_code"`
  ContentType MapContentType `json:"content_type"`
  Page int `json:"page"`
  Size int `json:"size"`
}

type MapGetAllData struct {
  Name string `json:"name"`
  Skin string `json:"skin"`
  X int `json:"x"`
  Y int `json:"y"`
  Content Content `json:"content"`
}

type MapGetAllResponse struct {
  ServerCodeInfo
  Data []MapGetAllData `json:"data"`
  Total int `json:"total"`
  Page int `json:"page"`
  Size int `json:"size"`
  Pages int `json:"pages"`
}
