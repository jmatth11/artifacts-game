package types

import "time"

type GameStateEnum int

const (
  GameStateRetry = GameStateEnum(-1)
  GameStateNoAction = GameStateEnum(0)
  GameStateLoad = GameStateEnum(1)
  GameStateFarm = GameStateEnum(2)
  GameStateDeposit = GameStateEnum(3)
  GameStateMove = GameStateEnum(4)
  GameStateUpgrade = GameStateEnum(5)
)

var ZeroTime time.Time

type GameState struct {
  State GameStateEnum
}
