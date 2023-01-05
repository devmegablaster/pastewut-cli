package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
  "github.com/atotto/clipboard"
)

func GetPasteWutContents() string {
  color.Yellow("Enter the text you want to paste!")

  scanner := bufio.NewScanner(os.Stdin)
  var lines []string

  for{
    scanner.Scan()
    line := scanner.Text()
    if line == "" {
      break
    }
    lines = append(lines, line)
  }

  return strings.Join(lines, "\n")
}

func GetContentOfFile(filename string) (string, error) {
  file, err := os.Open(filename)
  if err != nil {
    return "", fmt.Errorf("Error opening file: %s", err)
  }

  scanner := bufio.NewScanner(file)
  var lines []string

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return strings.Join(lines, "\n"), nil
}

func GetContentsOfClipboard() (string, error) {
  text, err := clipboard.ReadAll()
  if err != nil {
    return "", err
  }

  return string(text), nil
}
