package actions

import (
	"encoding/json"

	"github.com/jmatth11/artifacts-game/helpers"
	"github.com/jmatth11/artifacts-game/types"
)

func GetAllResources(c types.Client, req types.ResourceRequest) (types.ResourceList, error) {
  url := "resources"
  result, err := helpers.Get(c, url, req)
  var resp types.ResourceList
  if err != nil {
    return resp, err
  }
  if result.Type == types.ServerCodeOK {
    if err := json.Unmarshal(result.Data, &resp); err != nil {
      return resp, err
    }
  }
  resp.Type = result.Type
  return resp, nil
}

func GetAllMonsters(c types.Client, req types.MonsterListRequest) (types.MonsterList, error) {
  url := "monsters"
  result, err := helpers.Get(c, url, req)
  var resp types.MonsterList
  if err != nil {
    return resp, err
  }
  if result.Type == types.ServerCodeOK {
    if err := json.Unmarshal(result.Data, &resp); err != nil {
      return resp, err
    }
  }
  resp.Type = result.Type
  return resp, nil
}

