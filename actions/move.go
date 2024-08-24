package actions

import (
	"encoding/json"
	"fmt"

	"github.com/jmatth11/artifacts-game/helpers"
	"github.com/jmatth11/artifacts-game/types"
)

func Move(c types.Client, x, y int) (types.MoveResponse, error) {
  moveUrl := fmt.Sprintf("my/%s/action/move", c.Name)
  moveReq := types.MoveRequest{
    X: x,
    Y: y,
  }
  result, err := helpers.Post(c, moveUrl, moveReq)
  var resp types.MoveResponse
  if err != nil {
    return resp, err
  }
  if result.Type == types.ServerCodeOK {
    a := struct {
      Data types.MoveResponse `json:"data"`
    }{}
    if err := json.Unmarshal(result.Data, &a); err != nil {
      return resp, err
    }
    resp = a.Data
  }
  resp.Type = result.Type
  return resp, nil
}
