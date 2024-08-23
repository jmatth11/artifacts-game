package main

import (
	"fmt"

	"github.com/jmatth11/artifacts-game/actions"
	"github.com/jmatth11/artifacts-game/types"
	"github.com/joho/godotenv"
)

func main() {
  envFile, err := godotenv.Read(".env")
  if err != nil {
    panic(err)
  }
  c := types.Client {
    Token: envFile["TOKEN"],
    Name: "spudbud",
  }

  //resp, err := actions.Move(c, 2, 0)
  //if err != nil {
  //  fmt.Println(err)
  //  return
  //}
  //fmt.Println("%v+", resp)

  if err := actions.Gather(c); err != nil {
    fmt.Println(err)
  }
}
