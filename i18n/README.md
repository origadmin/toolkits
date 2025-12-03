# I18n

## 🚀 Automated Build (Recommended)

### Quick Start

We provide multiple automated build methods, choose any of the following:

#### Method 1: Using Shell Script (Linux/macOS)
```bash
# Execute complete build
./scripts/build_i18n.sh

# Clean files
./scripts/build_i18n.sh clean

# Show help
./scripts/build_i18n.sh help
```

#### Method 2: Using PowerShell Script (Windows)
```powershell
# Execute complete build
.\scripts\build_i18n.ps1

# Clean files
.\scripts\build_i18n.ps1 -Clean

# Show help
.\scripts\build_i18n.ps1 -Help
```

#### Method 3: Using Command Tool
```bash
# Execute complete build
go run cmd/make/main.go -generate-json

# Clean files
go run cmd/make/main.go -clean

# Show version
go run cmd/make/main.go -version
```

#### Method 4: Using Makefile (requires make command)
```bash
# Execute complete build
make -f scripts/Makefile.i18n all

# Or execute individual steps
make -f scripts/Makefile.i18n download-iana
make -f scripts/Makefile.i18n download-timezonedb
make -f scripts/Makefile.i18n download-cldr
make -f scripts/Makefile.i18n download-windows
make -f scripts/Makefile.i18n generate-json

# Clean files
make -f scripts/Makefile.i18n clean

# Show help
make -f scripts/Makefile.i18n help
```

## 📋 Manual Build Steps (Deprecated)

> ⚠️ **Note**: The following steps are manual build method and not recommended. Please prioritize using the automated scripts above.

### 1. Get tzdata from https://data.iana.org

#### 1.1 Download tzdata from https://data.iana.org

```bash
mkdir tzdb
cd tzdb
wget https://www.iana.org/time-zones/repository/tzcode-latest.tar.gz
wget https://www.iana.org/time-zones/repository/tzdata-latest.tar.gz
gzip -dc tzcode-latest.tar.gz | tar -xf -
gzip -dc tzdata-latest.tar.gz | tar -xf -
```

#### 1.2 Download tzdata from https://timezonedb.com

```bash
wget https://timezonedb.com/files/TimeZoneDB.csv.zip -O TimeZoneDB.csv.zip
zip -d TimeZoneDB.csv.zip
```

#### 1.3 Download ldmlSupplemental.dtd from https://github.com/unicode-org/cldr/blob/main/common/dtd/ldmlSupplemental.dtd

#### 1.4 Download windowsZones.xml from https://github.com/unicode-org/cldr/blob/main/common/supplemental/windowsZones.xml

## 🛠️ Automated Script Features

### ✅ Implemented Features
- **Automatic Download**: Download timezone data from multiple sources
- **Retry Mechanism**: Automatic retry on network failures
- **Error Handling**: Detailed error messages and logging
- **Cross-platform**: Support for Linux, macOS, and Windows
- **Cleanup Function**: One-click cleanup of all generated files
- **Dependency Check**: Automatic checking of required command-line tools

### 📁 Generated File Structure
```
toolkits/i18n/
├── tzdb/                    # IANA timezone data source
├── data/
│   ├── country.csv          # Country data
│   └── time_zone.csv        # Timezone data
├── windows/
│   ├── ldmlSupplemental.dtd # CLDR DTD file
│   ├── windowsZones.xml     # Windows timezone mapping
│   └── windows_zones.json   # Converted JSON
├── tz/
│   ├── country.json         # Country data JSON
│   └── time_zone.json       # Timezone data JSON
├── scripts/
│   ├── build_i18n.sh        # Linux/macOS build script
│   ├── build_i18n.ps1       # Windows PowerShell script
│   └── Makefile.i18n        # Make build file
└── cmd/make/
    ├── main.go              # Command-line tool
    └── go.mod               # Go module file
```

## 🔧 Dependency Requirements

### Required Tools
- **wget** or **curl** (for downloading files)
- **tar** (for extracting .tar.gz files)
- **gzip** (for extraction)
- **unzip** (for extracting .zip files)
- **go** (for generating JSON files)

### Optional Tools
- **7z** (alternative to tar and unzip on Windows)
- **make** (for Makefile build)

## 🚨 Troubleshooting

### Common Issues

1. **Network Connection Issues**
   ```bash
   # Check network connection
   ping www.iana.org
   ping timezonedb.com
   ```

2. **Permission Issues**
   ```bash
   # Linux/macOS
   chmod +x scripts/build_i18n.sh
   
   # Windows (PowerShell administrator mode)
   Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
   ```

3. **Missing Dependencies**
   ```bash
   # Ubuntu/Debian
   sudo apt-get install wget tar gzip unzip golang-go
   
   # macOS
   brew install wget golang
   
   # Windows (using Chocolatey)
   choco install wget golang
   ```

## 📝 Contributing

If you find issues with the automated scripts or have improvement suggestions, please submit an Issue or Pull Request.
