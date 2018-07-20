package main

import (
	"encoding/json"
	"errors"
	"flag"
  "strconv"
	"fmt"
	"io/ioutil"
	"os"
  "strings"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"
)

const (
  DEFAULT_EXTENSION = ".md"
)

func main() {
  flags := parseFlags()
  config := readConfigFile()

  createNewNote(config, flags) 
}

// ==========================
// Handling Flags

func parseFlags() Flags {
	filename := flag.String("f", getDate(), "Manually set the filename for a newly created card")
	flag.Parse()

	return Flags{*filename}
}

type Flags struct {
	filepath string
}

// ==========================
// Config File

type Opts struct {
	InboxPath string
	Counter  int
}

func checkForConfig() {
	config := getConfigPath()
	_, err := os.Stat(config)
	if os.IsNotExist(err) {
    defaultInbox := getHomeDir() + "/newNotes/"
		defaultOpts := Opts{InboxPath: defaultInbox, Counter: 0}

		j, err := json.MarshalIndent(defaultOpts, "", "  ")
		check(err)

		err = ioutil.WriteFile(config, j, 0644)
		check(err)
	}
}

func readConfigFile() Opts {
	checkForConfig()

	data, err := ioutil.ReadFile(getConfigPath())
	check(err)

	var opts Opts
	err = json.Unmarshal(data, &opts)
	check(err)

	return opts
}

func writeConfig(opts Opts) {
  config := getConfigPath()
  j, err := json.MarshalIndent(opts, "", "  ")
  check(err)

  err = ioutil.WriteFile(config, j, 0644)
  check(err)
}

// ==========================
// Creating a new note.

func createNewNote(config Opts, flags Flags) {
  fileName := createFileName(config)
  prefabbedContent := defaultTextString()
  createInboxDirIfNotExists(config)

  err := ioutil.WriteFile(fileName, []byte(prefabbedContent), 0644)
  check(err)

  err = launchEditor(fileName)
  check(err)

  config.Counter += 1
  writeConfig(config)
}

func getEditor() (string, error) {
	editor := os.Getenv("EDITOR")

	if editor != "" {
		return editor, nil
	} else {
		return "", errors.New("EDITOR not set.")
	}
}

func launchEditor(filename string) error {
  editorPath, err := getEditor()
  check(err)

	cmdArgs := append([]string{"--"}, filename)
	cmd := exec.Command(editorPath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()

	return nil
}

func createInboxDirIfNotExists(c Opts) {
  os.MkdirAll(c.InboxPath, os.ModePerm)
}

func createFileName(c Opts) string {
  title := flag.Args()
  if len(title) == 0 {
    return c.InboxPath + strconv.Itoa(c.Counter) + DEFAULT_EXTENSION 
  } else {
    return c.InboxPath + strings.Join(title, "-") + DEFAULT_EXTENSION
  }
} 

func defaultTextString() string {
  title := flag.Args()
  if len(title) == 0 {
    return fmt.Sprintf("# \n _(%s)_", getDate())
  } else {
    return fmt.Sprintf("# %s\n _(%s)", strings.Join(title, " "), getDate())
  }

}

// ==========================
// Util

func getDate() string {
	t := time.Now()
	return t.String()[0:10]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func getConfigPath() string {
	home := getHomeDir()
	return filepath.Join(home, ".config", "nn")
}
