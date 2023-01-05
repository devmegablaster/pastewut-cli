package actions

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/urfave/cli"

	"github.com/devmegablaster/pastewut-cli/pkg/helpers"
	"github.com/devmegablaster/pastewut-cli/pkg/models"
)

func CreatePasteWut(c *cli.Context) {
  if c.Bool("clipboard") {
    CreatePasteWutFromClipboard(c)
    return
  }

  if c.String("file") != "" {
    CreatePasteWutFromFile(c)
    return
  }

  var text string

  text = helpers.GetPasteWutContents()

  pastewut := models.PasteWut{
    Content: text,
  }

  resp, err := helpers.PostBackend("pastewut", pastewut)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Decode the response
  result, err := helpers.DecodeResponse(resp)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Print the response
  helpers.PrintPasteWutCode(result.Code)
}

func GetPasteWut(c *cli.Context) {
  code := c.Args().Get(0)
  if code == "" {
    helpers.HandleError(fmt.Errorf("No code provided!"))
    return
  }

  if c.Bool("clipboard") {
    GetPasteWutToClipboard(c)
    return
  }

  if c.String("output") != "" {
    GetPasteWutToFile(c)
    return
  }

  // Make a request to the API
  resp, err := helpers.GetBackend("pastewut/" + code)
  if err != nil {
    fmt.Println(color.RedString("Error: %s", err))
    return
  }

  // Handle 404
  err = helpers.HandleStatusErrors(resp.StatusCode)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Decode the response
  result, err := helpers.DecodeResponse(resp)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Print the response
  helpers.PrintPasteWutContent(result.Content)
}

func CreatePasteWutFromClipboard(c *cli.Context) {
  text, err := clipboard.ReadAll()
  if err != nil {
    fmt.Println(color.RedString("Error: %s", err))
    return
  }

  pastewut := models.PasteWut{
    Content: string(text),
  }

  resp, err := helpers.PostBackend("pastewut", pastewut)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Decode the response
  result, err := helpers.DecodeResponse(resp)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Print the response
  helpers.PrintPasteWutCode(result.Code)
}

func CreatePasteWutFromFile(c *cli.Context) {
  text, err := helpers.GetContentOfFile(c.String("file"))
  if err != nil {
    helpers.HandleError(err)
    return
  }

  pastewut := models.PasteWut{
    Content: text,
  }

  resp, err := helpers.PostBackend("pastewut", pastewut)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Decode the response
  result, err := helpers.DecodeResponse(resp)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Print the response
  helpers.PrintPasteWutCode(result.Code)
}

func GetPasteWutToClipboard(c *cli.Context) {
  code := c.Args().Get(0)

  // Make a request to the API
  resp, err := helpers.GetBackend("pastewut/" + code)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Handle 404
  err = helpers.HandleStatusErrors(resp.StatusCode)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Decode the response
  result, err := helpers.DecodeResponse(resp)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Copy the response to Clipboard
  clipboard.WriteAll(result.Content)

  // Print the response
  helpers.PrintPasteWutContent(result.Content)
  fmt.Println()
  color.Green("The contents of your PasteWut have been copied to your clipboard!")
}

func GetPasteWutToFile(c* cli.Context) {
  code :=  c.Args().Get(0)

  // Make a request to the API
  resp, err := helpers.GetBackend("pastewut/" + code)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Handle 404
  err = helpers.HandleStatusErrors(resp.StatusCode)

  // Decode the response
  result, err := helpers.DecodeResponse(resp)
  if err != nil {
    helpers.HandleError(err)
    return
  }

  // Write the response to a File
  err = helpers.WriteContentToFile(c.String("output"), result.Content)

  // Print the response
  helpers.PrintPasteWutContent(result.Content)
  fmt.Println()
  color.Green("The contents of your PasteWut have been written to %s!", c.String("output"))
}
