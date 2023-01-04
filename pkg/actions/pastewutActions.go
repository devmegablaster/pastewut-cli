package actions

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"
	"github.com/atotto/clipboard"

	"github.com/devmegablaster/pastewut-cli/pkg/models"
)

func CreatePasteWut(c *cli.Context) {
  if c.Bool("clipboard") {
    CreatePasteWutFromClipboard(c)
    return
  }

  color.Yellow("Enter the text you want to paste!")
  var text string

  scanner := bufio.NewScanner(os.Stdin)
  var lines []string

  for {
    scanner.Scan()
    line := scanner.Text()
    if line == "" {
      break
    }
    lines = append(lines, line)
  }

  text = strings.Join(lines, "\n")

  postBody, _ := json.Marshal(&models.PasteWut{
    Content: text,
  })

  requestBody := bytes.NewBuffer(postBody)

  // Make a request to the API
  resp, err := http.Post(BackendUrl + "pastewut/", "application/json", requestBody)
  if err != nil {
    panic(err)
  }

  // Decode the response
  var result models.PasteWut
  json.NewDecoder(resp.Body).Decode(&result)

  // Print the response
  fmt.Println()
  color.Cyan("PasteWut Code: %s", color.New(color.FgHiGreen).Add(color.Bold).Add(color.Underline).Sprint(result.Code))
}

func GetPasteWut(c *cli.Context) {
  if c.Bool("clipboard") {
    GetPasteWutToClipboard(c)
    return
  }

  if c.String("output") != "" {
    GetPasteWutToFile(c)
    return
  }

  code := c.Args().Get(0)

  // Make a request to the API
  resp, err := http.Get(BackendUrl + "pastewut/" + code)
  if err != nil {
    panic(err)
  }

  // Decode the response
  var result models.PasteWut
  json.NewDecoder(resp.Body).Decode(&result)

  // Print the response
  fmt.Println()
  color.Green("Here are the contents of your PasteWut!")
  fmt.Println()
  color.Yellow(result.Content)
}

func CreatePasteWutFromClipboard(c *cli.Context) {
  text, err := clipboard.ReadAll()
  if err != nil {
    panic(err)
  }

  postBody, _ := json.Marshal(&models.PasteWut{
    Content: string(text),
  })

  requestBody := bytes.NewBuffer(postBody)

  // Make a request to the API
  resp, err := http.Post(BackendUrl + "pastewut", "application/json", requestBody)
  if err != nil {
    panic(err)
  }

  // Decode the response
  var result models.PasteWut
  json.NewDecoder(resp.Body).Decode(&result)

  // Print the response
  fmt.Println()
  color.Cyan("PasteWut Code: %s", color.New(color.FgHiGreen).Add(color.Bold).Add(color.Underline).Sprint(result.Code))
}

func GetPasteWutToClipboard(c *cli.Context) {
  code := c.Args().Get(0)

  // Make a request to the API
  resp, err := http.Get(BackendUrl + "pastewut/" + code)
  if err != nil {
    panic(err)
  }

  // Decode the response
  var result models.PasteWut
  json.NewDecoder(resp.Body).Decode(&result)

  // Copy the response to Clipboard
  clipboard.WriteAll(result.Content)

  // Print the response
  fmt.Println()
  color.Green("Here are the contents of your PasteWut!")
  fmt.Println()
  color.Yellow(result.Content)
  fmt.Println()
  color.Green("The contents of your PasteWut have been copied to your clipboard!")
}

func GetPasteWutToFile(c* cli.Context) {
  code :=  c.Args().Get(0)

  // Make a request to the API
  resp, err := http.Get(BackendUrl + "pastewut/" + code)
  if err != nil {
    panic(err)
  }

  // Decode the response
  var result models.PasteWut
  json.NewDecoder(resp.Body).Decode(&result)

  // Write the response to a File
  fileName := c.String("output")
  file, err := os.Create(fileName)
  if err != nil {
    panic(err)
  }

  _, err = file.WriteString(result.Content)
  defer file.Close()

  // Print the response
  fmt.Println()
  color.Green("Here are the contents of your PasteWut!")
  fmt.Println()
  color.Yellow(result.Content)
  fmt.Println()
  color.Green("The contents of your PasteWut have been written to %s!", fileName)
}
