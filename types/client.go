package types

type Client struct {
  Name string
  Token string
}

type Event struct {
  Task MapContentType
  DetailTask CodeName
}

type CharacterManager struct {
  State GameStateEnum
  PrevState GameStateEnum
  Event chan Event
  Cooldown Cooldown
  Character Character
  Client Client
  ResourceActionCount int
}

func (cm *CharacterManager) SetState(state GameStateEnum) {
  cm.PrevState = cm.State
  cm.State = state
}

