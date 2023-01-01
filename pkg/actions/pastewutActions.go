package actions

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/devmegablaster/pastewut-cli/pkg/models"
	"github.com/urfave/cli"
)

func CreatePasteWut(c *cli.Context) {
  fmt.Println("Enter The Text You Want To Paste:")
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
  fmt.Println("PasteWut Code: " + result.Code)
}

