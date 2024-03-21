package cmm

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	setValue         = "/SetValue"
	setValueIfNeeded = "/SetValueIfNeeded"
	changeValue      = "/ChangeValue"
	switchValue      = "/SwitchValue"
	getValue         = "/GetValue"
	turnOff          = "/TurnOff"
	turnOn           = "/TurnOn"
	switchOffOn      = "/SwitchOffOn"
	saveConfig       = "/SaveConfig"
	loadConfig       = "/LoadConfig"
	sText            = "/stext"
	sTab             = "/stab"
	sComma           = "/scomma"
	sHtml            = "/shtml"
	sVerHtml         = "/sverhtml"
	sXml             = "/sxml"
	sJson            = "/sjson"
	sMonitors        = "/smonitors"
)

// CommandExecutor interface defines the method for executing commands
type CommandExecutor interface {
	Execute(args []string) (string, error)
}

// executor struct uses a CommandExecutor to execute commands
type Executor struct {
	commandExecutor CommandExecutor
}

// SystemCommandExecutor is a concrete implementation of CommandExecutor
// that executes system commands using the executable name stored in it
type systemCommandExecutor struct {
	executable string
}

// Execute runs a command using SystemCommandExecutor's executable
func (sce *systemCommandExecutor) Execute(args []string) (string, error) {
	var out bytes.Buffer
	cmd := exec.Command(sce.executable, args...)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing command: %w", err)
	}
	return out.String(), nil
}

// Constructor for executor
func ControlMyMonitor(executable string) (*Executor, error) {
	if executable == "" {
		return nil, errors.New("executable cannot be empty")
	}
	if !isExecutableInPath(executable) {
		return nil, errors.New("executable missing from PATH")
	}
	// Instantiate SystemCommandExecutor with the executable name
	ce := &systemCommandExecutor{executable: executable}
	return &Executor{commandExecutor: ce}, nil
}

func isExecutableInPath(executable string) bool {
	// Retrieve the PATH environment variable
	path := os.Getenv("PATH")
	paths := strings.Split(path, string(os.PathListSeparator))

	// Check if the executable file exists in any of the directories in PATH
	for _, dir := range paths {
		fullPath := filepath.Join(dir, executable)
		if _, err := os.Stat(fullPath); err == nil {
			return true
		}
	}
	return false
}

func (e Executor) SetValue(monitor, VCPCode, value string) error {
	_, err := e.commandExecutor.Execute([]string{setValue, monitor, VCPCode, value})
	return err
}

func (e Executor) SetValueIfNeeded(monitor, VCPCode, value string) error {
	_, err := e.commandExecutor.Execute([]string{setValueIfNeeded, monitor, VCPCode, value})
	return err
}

func (e Executor) GetValue(monitor, VCPCode string) (string, error) {
	_, err := e.commandExecutor.Execute([]string{getValue, monitor, VCPCode})
	if err != nil {
		return "", err
	}
	return e.commandExecutor.Execute([]string{"echo", "$LASTEXITCODE"})
}

func (e Executor) TurnOff(monitor string) error {
	_, err := e.commandExecutor.Execute([]string{turnOff, monitor})
	return err
}

func (e Executor) TurnOn(monitor string) error {
	_, err := e.commandExecutor.Execute([]string{turnOn, monitor})
	return err
}

func (e Executor) SwitchOffOn(monitor string) error {
	_, err := e.commandExecutor.Execute([]string{switchOffOn, monitor})
	return err
}

func (e Executor) SwitchValue(monitor string, VCPCode string, values []string) error {
	command := []string{switchValue, monitor, VCPCode}
	command = append(command, values...)
	_, err := e.commandExecutor.Execute(command)
	return err
}

func (e Executor) SaveConfig(filename, monitor string) error {
	_, err := e.commandExecutor.Execute([]string{saveConfig, filename, monitor})
	return err
}

func (e Executor) LoadConfig(filename, monitor string) error {
	_, err := e.commandExecutor.Execute([]string{loadConfig, filename, monitor})
	return err
}

// Generalized Execute Function
func (e Executor) executeCommand(command []string, returnContents bool) (string, error) {
    if returnContents {
		// set filename to ""
		command[1] = "\"\""
		command = append(command, "|", "more")
    }
    return e.commandExecutor.Execute(command)
}

func (e Executor) SText(filename string, monitor string, returnContents bool) (string, error) {
    return e.executeCommand([]string{sText, filename, monitor}, returnContents)
}

func (e Executor) STab(filename string, monitor string, returnContents bool) (string, error) {
	return e.executeCommand([]string{sTab, filename, monitor}, returnContents)
}

func (e Executor) SComma(filename string, monitor string, returnContents bool) (string, error) {
	return e.executeCommand([]string{sComma, filename, monitor}, returnContents)
}

func (e Executor) SHtml(filename string, monitor string, returnContents bool) (string, error) {
	return e.executeCommand([]string{sHtml, filename, monitor}, returnContents)
}

func (e Executor) SVerHtml(filename string, monitor string, returnContents bool) (string, error) {
	return e.executeCommand([]string{sVerHtml, filename, monitor}, returnContents)
}

func (e Executor) SXml(filename string, monitor string, returnContents bool) (string, error) {
	return e.executeCommand([]string{sXml, filename, monitor}, returnContents)
}

func (e Executor) SJson(filename string, monitor string, returnContents bool) (string, error) {
	return e.executeCommand([]string{sJson, filename, monitor}, returnContents)
}

func (e Executor) SMonitors(filename string, returnContents bool) (string, error) {
	return e.executeCommand([]string{sMonitors, filename}, returnContents)
}
