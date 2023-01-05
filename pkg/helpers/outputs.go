package helpers

import (
	"fmt"
  "os"

	"github.com/fatih/color"
)

func PrintPasteWutCode(code string) {
  fmt.Println()
  color.Cyan("PasteWut Code: %s", color.New(color.FgHiGreen).Add(color.Bold).Add(color.Underline).Sprint(code))
}

func PrintPasteWutContent(content string) {
  fmt.Println()
  color.Green("Here are the contents of your PasteWut!")
  fmt.Println()
  color.Yellow(content)
}

func WriteContentToFile(fileName, content string) error {
  file, err := os.Create(fileName)
  if err != nil {
    return err
  }

  _, err = file.WriteString(content)
  if err != nil {
    return err
  }

  return nil
}

func HandleError(err error) {
  fmt.Println(color.RedString("Error: %s", err))
}
