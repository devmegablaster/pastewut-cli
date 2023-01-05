package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devmegablaster/pastewut-cli/pkg/models"
)

func PostBackend(path string, pastewut models.PasteWut) (*http.Response, error) {
  requestBody, err := json.Marshal(&models.PasteWut{
    Content: pastewut.Content,
  })
  if err != nil {
    return nil, err
  }

  resp, err := http.Post(BackendUrl+path, "application/json", bytes.NewBuffer(requestBody) )
  if err != nil {
    return nil, err
  }

  return resp, nil
}

func DecodeResponse(resp *http.Response) (models.PasteWut, error) {
  var result models.PasteWut
  err := json.NewDecoder(resp.Body).Decode(&result)
  if err != nil {
    return result, err
  }

  return result, nil
}

func GetBackend(path string) (*http.Response, error) {
  resp, err := http.Get(BackendUrl + path)
  
  if err != nil {
    return nil, err
  }

  return resp, nil
}

func HandleStatusErrors(status int) error {
  if status == 400 {
    return fmt.Errorf("Bad Request")
  }

  if status == 401 {
    return fmt.Errorf("Unauthorized")
  }

  if status == 404 {
    return fmt.Errorf("Not Found")
  }

  return nil
}
