package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

// Interaction is holding some var for creating interaction
type Interaction struct {
	Input  *os.File
	Answer string
}

// NewInteraction is creating a new Interaction
func NewInteraction() *Interaction {
	return &Interaction{
		Input: os.Stdin,
	}
}

// AskForInput is asking the user for a filename if it's not in func args
func (i *Interaction) AskForInput() string {
	exePath, _ := os.Executable()
	cwd := filepath.Dir(exePath)
	fmt.Printf("Filepath is not filled in. Actual path : %q\n", cwd)
	fmt.Println("Enter the path to file:")

	var input string
	fmt.Scanln(&input)
	return filepath.Join(cwd, input)
}

// Notify is a commun func to log some actions taken by the script
func (i *Interaction) Notify(action string, args ...interface{}) {
	switch action {
	case "context":
		i.printContext(args...)
	case "start":
		notifyStart(args...)
	case "end":
		notifyEnd(args...)
	}
}

func (i *Interaction) printContext(args ...interface{}) {
	list := (args[0]).([5]string)
	argsValue := args[1:]
	fmt.Println("------ Airtable Convertor")
	fmt.Println("| Convertor called with arguments :")

	for key := range list {
		fmt.Printf("|   - %s = %q\n", list[key], argsValue[key])
	}

	fmt.Println("******")
}

func notifyStart(args ...interface{}) {
	filename := (args[0]).(string)
	fmt.Printf("[Conversion started] File %q \n", filename)
}

func notifyEnd(args ...interface{}) {
	filename := (args[0]).(string)
	fmt.Printf("[Conversion terminated] File %q written\n", filename)
	fmt.Println("------ Airtable Convertor")
}
