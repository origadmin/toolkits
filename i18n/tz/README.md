# TimeZone Package

## Overview

The `tz` package provides comprehensive timezone functionality including:

- **Timezone Detection**: Automatically detect system timezone on Windows and Unix systems
- **Timezone Conversion**: Convert time between different timezones
- **Timezone Lookup**: Find timezones by name, country code, or abbreviation
- **Data Management**: Load and manage timezone and country data from embedded JSON files

## Features

### ✅ Core Functionality
- **Cross-platform timezone detection** (Windows/Unix)
- **Lazy loading** of timezone data for performance
- **Caching** of timezone locations
- **Comprehensive timezone information**
- **Time conversion** between timezones
- **DST (Daylight Saving Time) support**

### ✅ Data Sources
- **IANA Timezone Database** via tzdata
- **TimeZoneDB.com** CSV data
- **Windows Zone Mapping** from CLDR
- **Country information** with timezone mapping

## API Reference

### Core Functions

```go
// Get current system timezone location
loc, err := tz.GetLocation()
locStr := tz.GetLocationString()

// Find timezones
tz, found := tz.FindTimeZoneByZoneName("Asia/Shanghai")
timezones, found := tz.FindTimeZonesByCountryCode("CN")
timezones, found := tz.FindTimeZonesByAbbreviation("CST")

// Get timezone information
info, err := tz.GetTimeZoneInfo("Asia/Shanghai")

// Convert time between timezones
converted, err := tz.ConvertTime(time.Now(), "Asia/Shanghai", "America/New_York")

// Check DST
isDST, err := tz.IsDST(time.Now(), "America/New_York")

// Get timezone offset
offset, err := tz.GetOffset("Asia/Shanghai")
```

### Data Access

```go
// Get all data
countries := tz.GetCountries()
timezones := tz.GetTimeZones()
```

## Data Structures

### TimeZone
```go
type TimeZone struct {
    ZoneName     string         `json:"zone_name"`
    CountryCode  string         `json:"country_code"`
    Abbreviation string         `json:"abbreviation"`
    GmtOffset    int          `json:"gmt_offset"` // in seconds
    Dst          bool         `json:"dst"`        // daylight saving time
    Location     *time.Location `json:"-"`
}
```

### TimeZoneInfo
```go
type TimeZoneInfo struct {
    Name         string         `json:"name"`
    Country      string         `json:"country"`
    Offset       int            `json:"offset"`       // in seconds
    OffsetString string         `json:"offset_string"` // "+08:00" format
    IsDST        bool          `json:"is_dst"`
    Abbreviation string         `json:"abbreviation"`
    Location     *time.Location `json:"-"`
}
```

## Performance Features

### Lazy Loading
Timezone data is loaded only when first accessed to improve startup time.

### Caching
Timezone locations are cached to avoid repeated loading from system.

### Memory Efficient
Duplicate timezone entries are filtered out during initialization.

## Platform Support

### Unix/Linux/macOS
- Reads from `/etc/timezone` file
- Falls back to `/etc/localtime` symlink parsing
- Default: `Asia/Shanghai`

### Windows
- Uses `tzutil.exe` command
- Maps Windows timezone IDs to IANA timezone names
- Default: `Asia/Shanghai`

## Build and Data Generation

### Generate JSON from CSV
```bash
cd tz
go run make.go -generate-json
```

### Clean Generated Files
```bash
go run ../../cmd/make/main.go -clean
```

## Error Handling

The package defines standard errors:
- `ErrTimeZoneNotFound`: Timezone not found
- `ErrInvalidTimeZone`: Invalid timezone format
- `ErrConversionFailed`: Timezone conversion failed
- `ErrLocationNotFound`: Location not found

## Testing

Run tests with:
```bash
go test ./...
go test -v ./...
```

## Dependencies

- Go 1.21+
- Standard library only (no external dependencies)

## File Structure

```
tz/
├── types.go           # Type definitions and interfaces
├── errors.go          # Error definitions
├── tz.go             # Core functionality
├── tz_unix.go        # Unix-specific implementation
├── tz_windows.go     # Windows-specific implementation
├── time_zone.go       # TimeZone types and data
├── country.go        # Country types and data
├── make.go           # Data generation tool
├── *_test.go         # Test files
├── data/             # CSV data files
├── windows/          # Windows timezone data
├── *.json           # Embedded data files
└── README.md        # This file
```

## Usage Examples

### Basic Usage
```go
package main

import (
    "fmt"
    "time"
    "github.com/origadmin/toolkits/i18n/tz"
)

func main() {
    // Get current timezone
    loc, err := tz.GetLocation()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Current timezone: %s\n", loc.String())
    
    // Convert time
    now := time.Now()
    converted, err := tz.ConvertTime(now, "UTC", "Asia/Shanghai")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("UTC: %s\n", now.UTC())
    fmt.Printf("Shanghai: %s\n", converted)
}
```

### Advanced Usage
```go
// Get detailed timezone information
info, err := tz.GetTimeZoneInfo("America/New_York")
if err != nil {
    panic(err)
}

fmt.Printf("Timezone: %s\n", info.Name)
fmt.Printf("Country: %s\n", info.Country)
fmt.Printf("Offset: %s\n", info.OffsetString)
fmt.Printf("DST: %t\n", info.IsDST)

// Find all timezones for a country
timezones, found := tz.FindTimeZonesByCountryCode("US")
if found {
    fmt.Printf("Found %d timezones for US\n", len(timezones))
    for _, tz := range timezones {
        fmt.Printf("- %s (%s)\n", tz.ZoneName, tz.Abbreviation)
    }
}
```