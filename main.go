package main

import (
	"log"
	"strings"
	"time"

	"github.com/jmatth11/artifacts-game/logic"
	"github.com/jmatth11/artifacts-game/types"
	"github.com/joho/godotenv"
)

func character_loop(c chan bool, cm *types.CharacterManager, task types.Event) {
  cm.State = types.GameStateLoad
  for {
    switch cm.State {
      case types.GameStateLoad: {
        log.Printf("Client %s: Loading\n", cm.Client.Name)
        err := logic.UpdateCharacterDetails(cm)
        if err != nil {
          log.Println(err)
          cm.State = types.GameStateRetry
        }
      }
      case types.GameStateMove: {
        log.Printf("Client %s: Moving\n", cm.Client.Name)
        err := logic.MoveToSpot(cm, task)
        if err != nil {
          log.Println(err)
          cm.State = types.GameStateRetry
        }
      }
      case types.GameStateFarm: {
        log.Printf("Client %s: Farming\n", cm.Client.Name)
        err := logic.FarmResources(cm, task)
        if err != nil {
          log.Println(err)
          cm.State = types.GameStateRetry
        }
      }
      case types.GameStateDeposit: {
        log.Printf("Client %s: Bank Deposit\n", cm.Client.Name)
        err := logic.DeposityInventory(cm)
        if err != nil {
          log.Println(err)
          cm.State = types.GameStateRetry
        }
      }
      case types.GameStateUpgrade: {
        log.Printf("Client %s: Task Upgrade\n", cm.Client.Name)
        err := logic.UpgradeFarming(cm, &task)
        if err != nil {
          log.Println(err)
          cm.State = types.GameStateRetry
        }
      }
      case types.GameStateNoAction: {
        log.Printf("Client %s: No Action\n", cm.Client.Name)
        // wait for new action
        task = <-cm.Event
      }
      case types.GameStateRetry: {
        // wait 60 seconds then retry
        log.Printf("Client %s: Retrying in 60s because of error.\n", cm.Client.Name)
        time.Sleep(time.Second * 60)
        cm.State = types.GameStateLoad
      }
    }
  }
}

func main() {
  envFile, err := godotenv.Read(".env")
  if err != nil {
    panic(err)
  }
  players := strings.Split(envFile["PLAYERS"], ",")
  characters := make([]types.CharacterManager, 0, 5)
  mainChannel := make(chan bool)
  events:= []types.Event {
    {
      Task: types.ResourceContentType,
      DetailTask: types.AshTree,
    },
    {
      Task: types.ResourceContentType,
      DetailTask: types.CopperRocks,
    },
    {
      Task: types.MonsterContentType,
      DetailTask: types.Chicken,
    },
    {
      Task: types.ResourceContentType,
      DetailTask: types.AshTree,
    },
    {
      Task: types.ResourceContentType,
      DetailTask: types.CopperRocks,
    },
  }
  for idx, p := range players {
    charInfo := types.CharacterManager {
      Client: types.Client{
        Token: envFile["TOKEN"],
        Name: p,
      },
      State: types.GameStateLoad,
    }
    characters = append(characters, charInfo)
    go character_loop(mainChannel, &charInfo, events[idx])
  }
  <-mainChannel
}
