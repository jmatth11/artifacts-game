package types

type MapContentType string

const (
  MonsterContentType = MapContentType("monster")
  ResourceContentType = MapContentType("resource")
  WorkshopContentType = MapContentType("workshop")
  BankContentType = MapContentType("bank")
  GrandExchangeContentType = MapContentType("grand_exchange")
  TasksMasterContentType = MapContentType("tasks_master")
)

type MapGetAllRequest struct {
  ContentCode string `json:"content_code"`
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
  Data MapGetAllData `json:"data"`
  Total int `json:"total"`
  Page int `json:"page"`
  Size int `json:"size"`
  Pages int `json:"pages"`
}
