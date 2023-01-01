package utils

import (
	"github.com/urfave/cli"

  "github.com/devmegablaster/pastewut-cli/pkg/actions"
)

func CliCommands(app *cli.App) {
  app.Commands = []cli.Command{
    {
      Name: "Create a PasteWut",
      Aliases: []string{"p", "paste"},
      Usage: "Paste text to PasteWut and get a random link to share it.",
      Action: actions.CreatePasteWut,
    },
  }
}
