package utils

import (
	"github.com/urfave/cli"

  "github.com/devmegablaster/pastewut-cli/pkg/actions"
)

func CliCommands(app *cli.App) {
  app.Commands = []cli.Command{
    {
      Name: "Create PasteWut",
      Aliases: []string{"n", "new"},
      Usage: "Paste text to PasteWut and get a random link to share it.",
      Action: actions.CreatePasteWut,
    },
    {
      Name: "Get PasteWut",
      Aliases: []string{"g", "get"},
      Usage: "Get the text from a PasteWut link.",
      Action: actions.GetPasteWut,
    },
  }
}
