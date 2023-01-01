package utils

import (
	"os"

	"github.com/urfave/cli"
)

func StartCli(app *cli.App) {
  if err := app.Run(os.Args); err != nil {
    panic(err)
  }
}
