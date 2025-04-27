# Slogx Package

## Introduction

Slogx is an enhanced logging package built on top of Go's standard library `log/slog`. It provides rich extensions to
standard logging capabilities including:

- Multiple log formats (Text/JSON/Colorful output with Tint)
- Log rotation support via Lumberjack
- Configurable log levels, output destinations, and time formats
- Built-in support for console output and file logging
- Developer-friendly source code location tracking
- Terminal color output customization

## Features

- **Format Support**: Text/JSON/Tint/Dev formats
- **Rotation**: Automatic log rotation with Lumberjack
- **Flexibility**:
    - Configure output destinations (console/file)
    - Custom time formats via `WithTimeLayout()`
    - Enable source code location tracking with `WithAddSource()`
- **Color Control**: Disable color output with `WithNoColor()`
- **Level Control**: Set from `LevelDebug` to `LevelFatal`

## Installation

## Usage

### Basic Setup

```go
import "github.com/origadmin/toolkits/slogx"

func main() {
  logger := slogx.New(
    slogx.WithFile("app.log"),
    slogx.WithFormat(slogx.FormatJSON),
    slogx.WithLumberjack(&lumberjack.Logger{
    MaxSize: 50, // MB
    MaxAge:  28, // days
    }),
  )
}
```

### Advanced Examples

#### Colored Console Output

```go
logger := slogx.New(
  slogx.WithFormat(slogx.FormatTint),
  slogx.WithConsoleOnly(),
)
```

#### Developer-friendly Format

```go
logger := slogx.New(
  slogx.WithFormat(slogx.FormatDev),
  slogx.WithAddSource(),
  slogx.WithTimeLayout("2006-01-02 15:04:05.000"),
)
```

#### JSON Log Format with Rotation

```go
logger := slogx.New(
  slogx.WithFile("app.json.log"),
  slogx.WithFormat(slogx.FormatJSON),
  slogx.WithLumberjack(&lumberjack.Logger{
    Filename:   "app.json.log",
    MaxSize:    10, // megabytes per file
    MaxBackups: 5,
    MaxAge:     30, // days
  }),
)
```

## Configuration Options
| Option              | Description                                                 |
|---------------------|-------------------------------------------------------------|
| `WithFile`          | Specify log file path                                       |
| `WithLumberjack`    | Configure log rotation using Lumberjack                     |
| `WithFormat`        | Choose log format (FormatText/JSON/Tint/Dev)                |
| `WithTimeLayout`    | Set custom time format (default: 2006-01-02 15:04:05 MST)   |
| `WithAddSource`     | Include file/line number information in logs                |
| `WithConsole`       | Enable console output (can be combined with file output)    |
| `WithConsoleOnly`   | Output logs to console only (no file)                       |
| `WithLevel`         | Set log level (LevelDebug/Info/Warning/Error/Fatal)         |
| `WithNoColor`       | Disable terminal color output                               |
| `WithDefault`       | Set as global default logger                                |

## Format Details
- **FormatDev**: Developer-friendly format with color coding and source location
- **FormatTint**: Terminal-friendly colored text output
- **FormatJSON**: Machine-readable structured logging
- **FormatText**: Plain text format with basic timestamp and level info
