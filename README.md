# Logger

A customisable logger for Go.

## Installation

```shell
go get -u github.com/alexgokhale/logger
```

## Examples

This logger offers both default `Printf` style methods, and a `New` method to create a new logger with a prefix of your choice.

**One key difference between these methods and standard `Printf` is that a trailing newline _is_ added to the end of the string printed.**

### Default Methods

```go
name := "Alex"

// Print a critical message (coloured red)
// logger.Fatal acts identically to logger.Critical but calls `os.Exit(1)`
logger.Critical("%s broke something", name)
// Output: [CRIT] [2022-04-01T20:35:20.987] Alex broke something

// Print a warning message (coloured yellow)
logger.Warning("User '%s' missing, skipping entry", name)
// Output: [WARN] [2022-04-01T20:35:20.987] User 'Alex' missing, skipping entry

// Print an informational message (coloured blue)
logger.Info("Connecting to database...")
// Output: [INFO] [2022-04-01T20:35:20.987] Connecting to database...

// Print a success message (coloured green)
logger.Info("Connected successfully!")
// Output: [INFO] [2022-04-01T20:35:20.987] Connected successfully!

// Print a generic log message
logger.Log("Testing %d %d %d!", 1, 2, 3)
// Output: [2022-04-01T20:35:20.987] Testing 1 2 3!
```

### Custom Logger

```go
databaseLogger := logger.New("<DATABASE>")

databaseLogger.Info("Connecting...")
// Output: [INFO] [2022-04-01T20:35:20.987] <DATABASE> Connecting...
databaseLogger.Success("Connected successfully!")
// Output: [INFO] [2022-04-01T20:35:20.987] <DATABASE> Connected successfully!
```
