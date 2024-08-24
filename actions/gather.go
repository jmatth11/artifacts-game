package actions

import (
	"encoding/json"
	"fmt"

	"github.com/jmatth11/artifacts-game/helpers"
	"github.com/jmatth11/artifacts-game/types"
)

func Gather(c types.Client) (types.GatheringResponse, error) {
  url := fmt.Sprintf("my/%s/action/gathering", c.Name)
  result, err := helpers.Post(c, url, nil)
  var resp types.GatheringResponse
  if err != nil {
    return resp, err
  }
  if result.Type == types.ServerCodeOK {
    a := struct {
      Data types.GatheringResponse `json:"data"`
    }{}
    if err := json.Unmarshal(result.Data, &a); err != nil {
      return resp, err
    }
    resp = a.Data
  }
  resp.Type = result.Type
  return resp, nil
}

func Fight(c types.Client) (types.FightResponse, error) {
  url := fmt.Sprintf("my/%s/action/fight", c.Name)
  result, err := helpers.Post(c, url, nil)
  var resp types.FightResponse
  if err != nil {
    return resp, err
  }
  if result.Type == types.ServerCodeOK {
    a := struct {
      Data types.FightResponse `json:"data"`
    }{}
    if err := json.Unmarshal(result.Data, &a); err != nil {
      return resp, err
    }
    resp = a.Data
  }
  resp.Type = result.Type
  return resp, nil
}

func CharacterDetails(c types.Client) (types.CharacterResponse, error) {
  url := fmt.Sprintf("characters/%s", c.Name)
  result, err := helpers.Get(c, url, nil)
  var resp types.CharacterResponse
  if err != nil {
    return resp, err
  }
  if result.Type == types.ServerCodeOK {
    a := struct {
      Data types.CharacterResponse `json:"data"`
    }{}
    if err := json.Unmarshal(result.Data, &a); err != nil {
      return resp, err
    }
    resp = a.Data
  }
  resp.Type = result.Type
  return resp, nil
}

func DepositItemToBank(c types.Client, item types.SimpleItem) (types.BankResponse, error) {
  url := fmt.Sprintf("my/%s/action/bank/deposit", c.Name)
  result, err := helpers.Post(c, url, item)
  var resp types.BankResponse
  if err != nil {
    return resp, err
  }
  if result.Type == types.ServerCodeOK {
    a := struct {
      Data types.BankResponse `json:"data"`
    }{}
    if err := json.Unmarshal(result.Data, &a); err != nil {
      return resp, err
    }
    resp = a.Data
  }
  resp.Type = result.Type
  return resp, nil
}
