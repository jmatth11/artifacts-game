package actions

import (
	"fmt"

	"github.com/jmatth11/artifacts-game/helpers"
	"github.com/jmatth11/artifacts-game/types"
)

func Gather(c types.Client) error {
  url := fmt.Sprintf("my/%s/action/gathering", c.Name)
  result, err := helpers.Post(c, url, nil)
  if err != nil {
    return err
  }
  fmt.Println(string(result))
  return nil
}
