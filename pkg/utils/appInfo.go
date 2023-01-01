package utils

import (
  "github.com/urfave/cli"
)

func CliInfo(app *cli.App) {
  app.Name = "PasteWut CLI"
  app.Usage = "A CLI for PasteWut. PasteWut is a pastebin service that allows you to paste text and get a link to share it."
  app.Version = "1.0.0"
  app.Author = "devmegablaster"
  app.Email = "devmegablaster@gmail.com"
}
