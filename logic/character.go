package logic

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/jmatth11/artifacts-game/actions"
	"github.com/jmatth11/artifacts-game/types"
)

var skillList = []types.SkillEnum{
  types.SkillWoodcutting, types.SkillMining, types.SkillFishing,
}

func waitForCooldown(cm *types.CharacterManager) {
  if cm.Cooldown.StartedAt.Equal(types.ZeroTime) {
    return
  }
  //now := time.Now().UTC()
  //diff := now.Sub(cm.Cooldown.StartedAt)
  //log.Printf("diff time %d\n", diff)
  //if diff > 0 {
  //  time.Sleep(diff)
  //}
  time.Sleep(time.Second * time.Duration(cm.Cooldown.RemainingSeconds))
  cm.Cooldown.StartedAt = types.ZeroTime
  cm.Cooldown.RemainingSeconds = 0
}

func UpdateCharacterDetails(cm *types.CharacterManager) error {
  waitForCooldown(cm)
  resp, err := actions.CharacterDetails(cm.Client)
  if err != nil {
    return err
  }
  if resp.Type == types.ServerCodeOK {
    cm.Character = resp.Data
    cm.SetState(types.GameStateMove)
  } else {
    log.Printf("Client %s: update character response code: %d\n", cm.Client.Name, resp.Type)
    cm.SetState(types.GameStateNoAction)
  }
  return nil
}

func MoveToSpot(cm *types.CharacterManager, task types.Event) error {
  waitForCooldown(cm)
  mapReq := types.MapGetAllRequest{
    ContentCode: task.DetailTask,
    ContentType: task.Task,
    Page: 1,
    Size: 10,
  }
  resp, err := actions.MapGetAll(cm.Client, mapReq)
  if err != nil {
    return err
  }
  if resp.Type == types.ServerCodeNotFound {
    // TODO pull out as error type
    return fmt.Errorf("client %s: resource not found", cm.Client.Name)
  }
  var foundSpot types.MapGetAllData
  for _, spot := range resp.Data {
    if spot.Content.Code == task.DetailTask {
      foundSpot = spot
      break
    }
  }
  moveResp, err := actions.Move(cm.Client, foundSpot.X, foundSpot.Y)
  if err != nil {
    return err
  }
  if moveResp.Type == types.ServerCodeOK {
    log.Printf("Client %s: moved to (%d, %d)\n", cm.Client.Name, foundSpot.X, foundSpot.Y)
    cm.Character = moveResp.Character
    cm.Cooldown = moveResp.Cooldown
  } else if moveResp.Type == types.ServerCodeNotFound {
    cm.SetState(types.GameStateNoAction)
  }
  switch task.Task {
    case types.ResourceContentType:
      fallthrough
    case types.MonsterContentType:
      cm.SetState(types.GameStateFarm)
    case types.BankContentType:
      cm.SetState(types.GameStateDeposit)
  }
  return nil
}

func UpgradeFarming(cm *types.CharacterManager, task *types.Event) error {
  if cm.ResourceActionCount >= 3 {
    if task.Task == types.MonsterContentType {
      task.Task = types.ResourceContentType
    } else if task.Task == types.ResourceContentType {
      task.Task = types.MonsterContentType
    }
    cm.ResourceActionCount = 0
  }
  switch task.Task {
    case types.MonsterContentType: {
      req := types.MonsterListRequest{
      	Drop:     "",
      	MaxLevel: cm.Character.Level,
      	MinLevel: int(math.Max(float64(cm.Character.Level-5), 0)),
      	Page:     1,
      	Size:     10,
      }
      resp, err := actions.GetAllMonsters(cm.Client, req)
      if err != nil {
        return err
      }
      if resp.Type == types.ServerCodeOK && len(resp.Data) > 0 {
        item := resp.Data[0]
        task.DetailTask = item.Code
      } else {
        cm.SetState(types.GameStateNoAction)
      }
    }
    case types.ResourceContentType: {
      req := types.ResourceRequest{
      	Drop:     "",
      	MaxLevel: cm.Character.Level,
      	MinLevel: int(math.Max(float64(cm.Character.Level-5), 0)),
      	Page:     1,
      	Size:     10,
        Skill: skillList[cm.ResourceActionCount],
      }
      resp, err := actions.GetAllResources(cm.Client, req)
      if err != nil {
        return err
      }
      if resp.Type == types.ServerCodeOK && len(resp.Data) > 0 {
        item := resp.Data[0]
        task.DetailTask = item.Code
      } else {
        log.Printf("Client %s: upgrade code errored, no action\n", cm.Client.Name)
        cm.SetState(types.GameStateNoAction)
      }
    }
  }
  cm.ResourceActionCount++
  cm.SetState(types.GameStateLoad)
  return nil
}

func FarmResources(cm *types.CharacterManager, task types.Event) error {
  waitForCooldown(cm)
  switch task.Task {
    case types.MonsterContentType: {
      fightResp, err := actions.Fight(cm.Client)
      if err != nil {
        return err
      }
      if fightResp.Type == types.ServerCodeOK {
        cm.Character = fightResp.Character
        cm.Cooldown = fightResp.Cooldown
        fmt.Println(strings.Join(fightResp.Fight.Logs, "\n"))
      } else if fightResp.Type == types.ServerCodeResourceNotFoundOnMap {
        cm.SetState(types.GameStateMove)
      } else if fightResp.Type == types.ServerCodeInventoryFull {
        cm.SetState(types.GameStateDeposit)
      } else {
        log.Printf("Client %s: gather monster error resp code %d\n", cm.Client.Name, fightResp.Type)
      }
    }
    case types.ResourceContentType: {
      gatherResp, err := actions.Gather(cm.Client)
      if err != nil {
        return err
      }
      if gatherResp.Type == types.ServerCodeOK {
        cm.Character = gatherResp.Character
        cm.Cooldown = gatherResp.Cooldown
        log.Printf("Client %s: gather XP %d\n", cm.Client.Name, gatherResp.Details.Xp)
      } else if gatherResp.Type == types.ServerCodeResourceNotFoundOnMap {
        log.Printf("Client %s: gather resource not found\n", cm.Client.Name)
        cm.SetState(types.GameStateMove)
      } else if gatherResp.Type == types.ServerCodeInventoryFull {
        log.Printf("Client %s: gather inventory full\n", cm.Client.Name)
        cm.SetState(types.GameStateDeposit)
      } else {
        log.Printf("Client %s: gather resource error resp code %d\n", cm.Client.Name, gatherResp.Type)
      }
    }
  }
  return nil
}

func DeposityInventory(cm *types.CharacterManager) error {
  waitForCooldown(cm)
  // hard coded bank coordinates
  moveResp, err := actions.Move(cm.Client, 4, 1)
  if err != nil {
    return err
  }
  if moveResp.Type == types.ServerCodeOK {
    cm.Character = moveResp.Character
    cm.Cooldown = moveResp.Cooldown
    for _, item := range moveResp.Character.Inventory {
      if item.Quantity == 0 {
        continue
      }
      waitForCooldown(cm)
      depositItem := types.SimpleItem {
        Code: item.Code,
        Quantity: item.Quantity,
      }
      bankResp, err := actions.DepositItemToBank(cm.Client, depositItem)
      if err != nil {
        return err
      }
      if bankResp.Type == types.ServerCodeOK {
        cm.Character = bankResp.Character
        cm.Cooldown = bankResp.Cooldown
      } else if bankResp.Type == types.ServerCodeInsufficientQuantity {
        log.Printf("Client %s: Insufficient quantity: Item %s\n", cm.Client.Name, item.Code)
      }
    }
    cm.SetState(types.GameStateUpgrade)
  } else {
    cm.SetState(types.GameStateRetry)
  }
  return nil
}
