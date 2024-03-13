# cmm_executor - A Go Wrapper for ControlMyMonitor

`cmm_executor` is a Go package providing a convenient wrapper around the [ControlMyMonitor Utility](https://www.nirsoft.net/utils/control_my_monitor.html) by NirSoft. This package allows developers to programmatically control their monitors for tasks like changing settings, saving/loading configurations, and managing power states, directly from Go applications.

## Getting Started

### Prerequisites

- Go 1.15 or later.
- ControlMyMonitor (1.37) executable must be installed and accessible in your system's PATH.

### Installation

To install `cmm_executor`, run the following command in your terminal:

```bash
go get github.com/msetsma/CMM-Executor
```

## Usage

Import `cmm_executor` into your Go project:

```go
import "cmm_executor"
```

### Creating an Executor

To use the package, you must first create an `Executor` instance by providing the path to the ControlMyMonitor executable:

```go
executor, err := cmm_executor.ControlMyMonitor("ControlMyMonitor.exe")
if err != nil {
    log.Fatalf("Failed to create executor: %v", err)
}
```

### Executing Commands

With an `Executor` instance, you can now execute commands to control your monitor. Here are a few examples:

#### Turning Off a Monitor

```go
if err := executor.TurnOff("monitorID"); err != nil {
    log.Printf("Failed to turn off monitor: %v", err)
}
```

#### Changing a Value

```go
if err := executor.SetValue("monitorID", "VCPCode", "newValue"); err != nil {
    log.Printf("Failed to set value: %v", err)
}
```

## Supported Commands

- **SetValue**: Sets a new value for a specified setting. Requires monitor ID, VCP code, and new value.

- **SetValueIfNeeded**: Sets a new value only if different from the current one. Requires monitor ID, VCP code, and new value.

- **GetValue**: Retrieves the current value of a setting. Requires monitor ID and VCP code.

- **TurnOff**: Turns off the monitor. Requires monitor ID.

- **TurnOn**: Turns on the monitor. Requires monitor ID.

- **SwitchOffOn**: Toggles the monitor's power state off, then on. Requires monitor ID.

- **SaveConfig**: Saves current settings to a file. Requires filename and monitor ID.

- **LoadConfig**: Loads settings from a file. Requires filename and monitor ID.

- **SText**: Exports settings to a text file. Requires filename and monitor ID.

- **STab**: Exports settings in tab-delimited format. Requires filename and monitor ID.

- **SComma**: Exports settings in CSV format. Requires filename and monitor ID.

- **SHtml**: Exports settings to HTML file. Requires filename and monitor ID.

- **SVerHtml**: Exports settings to vertical HTML file. Requires filename and monitor ID.

- **SXml**: Exports settings to XML file. Requires filename and monitor ID.

- **SJson**: Exports settings to JSON file. Requires filename and monitor ID.

- **SMonitors**: Exports all detected monitors' settings. Requires filename.


## Dependencies

- ControlMyMonitor by NirSoft must be installed on the host system.

## Contributing

Contributions to `cmm_executor` are welcome. Please follow the standard Git workflow - fork the repo, make your changes, and submit a pull request.

## TODO 

1. Implement the direct return of the `S` command results. In the utility you are able to add `| more` to the end of the commands and it will return the results directly to the terminal instead of writing to a file.