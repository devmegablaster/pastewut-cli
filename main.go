package main

import (
  "github.com/urfave/cli"

  "github.com/devmegablaster/pastewut-cli/pkg/utils"
)

func main() {
  utils.LoadEnv()

  app := cli.NewApp()
  utils.CliInfo(app)
  utils.CliCommands(app)
  utils.StartCli(app)
}
