package actions

import (
	"encoding/json"

	"github.com/jmatth11/artifacts-game/helpers"
	"github.com/jmatth11/artifacts-game/types"
)

func MapGetAll(c types.Client, req  types.MapGetAllRequest) (types.MapGetAllResponse, error) {
  url := "maps"
  result, err := helpers.Get(c, url, req)
  var resp types.MapGetAllResponse
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
