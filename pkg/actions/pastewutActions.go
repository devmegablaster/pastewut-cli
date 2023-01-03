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
	"golang.design/x/clipboard"

	"github.com/devmegablaster/pastewut-cli/pkg/models"
)

func CreatePasteWut(c *cli.Context) {
  color.Yellow("Enter the text you want to paste!")
  var text string

  reader := bufio.NewReader(os.Stdin)
  text, _ = reader.ReadString('\n')

  text = strings.TrimSuffix(text, "\n")

  postBody, _ := json.Marshal(&models.PasteWut{
    Content: text,
  })

  requestBody := bytes.NewBuffer(postBody)

  // Make a request to the API
  resp, err := http.Post("http://localhost:3000/pastewut", "application/json", requestBody)
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
  code := c.Args().Get(0)

  // Make a request to the API
  resp, err := http.Get("http://localhost:3000/pastewut/" + code)
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
  err := clipboard.Init()
  if err != nil {
    panic(err)
  }

  text := clipboard.Read(clipboard.FmtText)

  postBody, _ := json.Marshal(&models.PasteWut{
    Content: string(text),
  })

  requestBody := bytes.NewBuffer(postBody)

  // Make a request to the API
  resp, err := http.Post("http://localhost:3000/pastewut", "application/json", requestBody)
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
