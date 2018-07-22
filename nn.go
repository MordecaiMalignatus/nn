package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	DEFAULT_EXTENSION = ".md"
)

func main() {
	config := readConfigFile()

	createNewNote(config)
}

// ==========================
// Creating a new note.

func createNewNote(config Opts) {
	fileName := createFileName(config)
	prefabbedContent := defaultTextString()
  os.MkdirAll(config.InboxPath, os.ModePerm)

	err := ioutil.WriteFile(fileName, []byte(prefabbedContent), 0644)
	check(err)

	err = launchEditor(fileName)
	check(err)

	if noteWasChanged(fileName) {
		fmt.Println("Saving as " + fileName)
		config.Counter += 1
		writeConfig(config)
	} else {
		fmt.Println("Nothing was changed, discarding note...")
		err := os.Remove(fileName)
		check(err)
	}
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

func noteWasChanged(fileName string) bool {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	currentString := string(dat)

	return currentString != defaultTextString()
}

func createFileName(c Opts) string {
	title := os.Args[1:]
	if len(title) == 0 {
		return c.InboxPath + strconv.Itoa(c.Counter) + DEFAULT_EXTENSION
	} else {
		return c.InboxPath + strings.Join(title, "-") + DEFAULT_EXTENSION
	}
}

func defaultTextString() string {
	title := os.Args[1:]
	if len(title) == 0 {
		return fmt.Sprintf("# \n _(%s)_", getDate())
	} else {
		return fmt.Sprintf("# %s\n _(%s)", strings.Join(title, " "), getDate())
	}

}

// ==========================
// Config File

type Opts struct {
	InboxPath string
	Counter   int
}

func checkForConfig() {
	config := getConfigPath()
	_, err := os.Stat(config)
	if os.IsNotExist(err) {
		fmt.Println("Generating config file in ~/.config/nn...")
		fmt.Println("Default Inbox is ~/newNotes/, change in config if desired.")

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
