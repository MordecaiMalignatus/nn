package main

import (
	"errors"
	"fmt"
  "io/ioutil"
  "strings"
	"os"
	"os/exec"
	"time"
  "flag"
)

const CONFIG_FILE = "~/.config/nn"

func main() {
  fmt.Println(parseFlags())
}

func parseFlags() Opts {
  filename := flag.String("f", getDate(), "Manually set the filename for a newly created card")
  flag.Parse()

  return Opts { *filename }
}

func getEditor() (string, error) {
	editor := os.Getenv("EDITOR")

	if editor != "" {
		return editor, nil
	} else {
		return "", errors.New("EDITOR not set.")
	}
}

type Opts struct {
  filename string
}

func readConfigFile() (map[string]string, error) {
  checkForConfig()

  data, err := ioutil.ReadFile(CONFIG_FILE)
  check(err)

  configString := string(data)
  ret := make(map[string]string)
  
  for _, line := range strings.Split(configString, "\n") {
    separated := strings.Split(line, " = ")
    ret[separated[0]] = separated[1]
  }

  return ret, nil
}

func launchEditor(editorPath string, filename []string) error {
	cmdArgs := append([]string{"--"}, filename...)
	cmd := exec.Command(editorPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	fmt.Println(cmd.Run())

	return nil
}

func defaultTextString() string {
	return fmt.Sprintf("# \n _(%s)_", getDate())
}

func getDate() string {
  t := time.Now()
  return t.String()[0:10]
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
