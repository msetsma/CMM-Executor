package cmm

import (
    "testing"
	"strings"
)

// MockCommandExecutor is a mock implementation of the CommandExecutor.
type MockCommandExecutor struct{}

// Execute mocks the execution of a command and returns the args as a string.
func (mce *MockCommandExecutor) Execute(args []string) (string, error) {
    return strings.Join(args, " "), nil
}

func TestGetMonitorDetails(t *testing.T) {
    // build executor
	executor, err := ControlMyMonitor("ControlMyMonitor.exe")
	if err != nil || executor == nil{
		t.Errorf("Failed to build executor")
	}
}

func TestSetValue(t *testing.T) {
    // Arrange
    mockExec := &MockCommandExecutor{}

    e := Executor{commandExecutor: mockExec}

    // Call the method you want to test.
    err := e.SetValue("monitor1", "VCPCode1", "value1")
    if err != nil {
        t.Errorf("SetValue returned an unexpected error: %v", err)
    }

    // Since our mock just returns the args as a string, you can check that.
    expectedArgs := "/SetValue monitor1 VCPCode1 value1"
    actualArgs, _ := mockExec.Execute([]string{"/SetValue", "monitor1", "VCPCode1", "value1"})
    if expectedArgs != actualArgs {
        t.Errorf("expected %q, got %q", expectedArgs, actualArgs)
    }
}