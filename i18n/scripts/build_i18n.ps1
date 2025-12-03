# I18n Data Download and Build Script (PowerShell Version)
# Automates all steps from README.md

param(
    [switch]$Clean,
    [switch]$Help
)

# Color definitions
$Colors = @{
    Red = "Red"
    Green = "Green" 
    Yellow = "Yellow"
    White = "White"
}

# Logging functions
function Write-Log {
    param(
        [string]$Message,
        [string]$Color = $Colors.White
    )
    Write-Host "[$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')] $Message" -ForegroundColor $Color
}

function Write-Info {
    param([string]$Message)
    Write-Log -Message $Message -Color $Colors.Green
}

function Write-Warn {
    param([string]$Message)
    Write-Log -Message $Message -Color $Colors.Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Log -Message $Message -Color $Colors.Red
}

# Check if command exists
function Test-Command {
    param([string]$Command)
    try {
        Get-Command $Command -ErrorAction Stop | Out-Null
        return $true
    }
    catch {
        return $false
    }
}

# Create directories
function New-Directories {
    Write-Info "Creating necessary directories..."
    @("tzdb", "data", "windows", "temp") | ForEach-Object {
        if (!(Test-Path $_)) {
            New-Item -ItemType Directory -Path $_ -Force | Out-Null
        }
    }
}

# Download file with retry
function Invoke-DownloadWithRetry {
    param(
        [string]$Url,
        [string]$Output,
        [int]$MaxRetries = 3
    )
    
    $retryCount = 0
    
    while ($retryCount -lt $MaxRetries) {
        try {
            Write-Info "Downloading: $Url"
            Invoke-WebRequest -Uri $Url -OutFile $Output -TimeoutSec 30 -ErrorAction Stop
            Write-Info "Download successful: $Output"
            return $true
        }
        catch {
            $retryCount++
            Write-Warn "Download failed, retry $retryCount/$MaxRetries: $Url"
            Start-Sleep -Seconds 2
        }
    }
    
    Write-Error "Download failed, max retries reached: $Url"
    return $false
}

# 1. Download tzdata from IANA
function Get-IanaTzdata {
    Write-Info "Starting IANA tzdata download..."
    
    Set-Location tzdb
    
    # Download tzcode
    Write-Info "Downloading tzcode-latest.tar.gz..."
    if (!(Invoke-DownloadWithRetry -Url "https://www.iana.org/time-zones/repository/tzcode-latest.tar.gz" -Output "tzcode-latest.tar.gz")) {
        Set-Location ..
        return $false
    }
    
    # Download tzdata
    Write-Info "Downloading tzdata-latest.tar.gz..."
    if (!(Invoke-DownloadWithRetry -Url "https://www.iana.org/time-zones/repository/tzdata-latest.tar.gz" -Output "tzdata-latest.tar.gz")) {
        Set-Location ..
        return $false
    }
    
    # Extract
    Write-Info "Extracting tzcode..."
    try {
        & tar -xzf tzcode-latest.tar.gz
    }
    catch {
        Write-Warn "tar command failed, trying with 7z..."
        if (Test-Command "7z") {
            & 7z x tzcode-latest.tar.gz
            & 7z x tzcode-latest.tar
        }
        else {
            Write-Error "Need tar or 7z command to extract files"
            Set-Location ..
            return $false
        }
    }
    
    Write-Info "Extracting tzdata..."
    try {
        & tar -xzf tzdata-latest.tar.gz
    }
    catch {
        Write-Warn "tar command failed, trying with 7z..."
        if (Test-Command "7z") {
            & 7z x tzdata-latest.tar.gz
            & 7z x tzdata-latest.tar
        }
        else {
            Write-Error "Need tar or 7z command to extract files"
            Set-Location ..
            return $false
        }
    }
    
    Set-Location ..
    Write-Info "IANA tzdata download and extraction completed"
    return $true
}

# 2. Download data from TimeZoneDB
function Get-TimeZoneDB {
    Write-Info "Starting TimeZoneDB data download..."
    
    Set-Location temp
    
    if (!(Invoke-DownloadWithRetry -Url "https://timezonedb.com/files/TimeZoneDB.csv.zip" -Output "TimeZoneDB.csv.zip")) {
        Set-Location ..
        return $false
    }
    
    # Extract
    Write-Info "Extracting TimeZoneDB.csv.zip..."
    try {
        Expand-Archive -Path "TimeZoneDB.csv.zip" -DestinationPath . -Force
    }
    catch {
        Write-Warn "Expand-Archive failed, trying with 7z..."
        if (Test-Command "7z") {
            & 7z x TimeZoneDB.csv.zip
        }
        else {
            Write-Error "Need Expand-Archive or 7z command to extract files"
            Set-Location ..
            return $false
        }
    }
    
    # Move to data directory
    if (Test-Path "time_zone.csv") {
        Move-Item "time_zone.csv" "..\data\" -Force
        Write-Info "time_zone.csv moved to data directory"
    }
    
    Set-Location ..
    Write-Info "TimeZoneDB data download completed"
    return $true
}

# 3. Download CLDR DTD
function Get-CldrDtd {
    Write-Info "Starting CLDR DTD download..."
    
    if (!(Invoke-DownloadWithRetry -Url "https://raw.githubusercontent.com/unicode-org/cldr/main/common/dtd/ldmlSupplemental.dtd" -Output "windows/ldmlSupplemental.dtd")) {
        return $false
    }
    
    Write-Info "CLDR DTD download completed"
    return $true
}

# 4. Download Windows timezone XML
function Get-WindowsZones {
    Write-Info "Starting Windows timezone data download..."
    
    if (!(Invoke-DownloadWithRetry -Url "https://raw.githubusercontent.com/unicode-org/cldr/main/common/supplemental/windowsZones.xml" -Output "windows/windowsZones.xml")) {
        return $false
    }
    
    Write-Info "Windows timezone data download completed"
    return $true
}

# 5. Generate JSON files
function New-JsonFiles {
    Write-Info "Starting JSON file generation..."
    
    if (Test-Command "go") {
        Set-Location tz
        try {
            & go run make.go -generate-json
            Set-Location ..
            Write-Info "JSON file generation completed"
            return $true
        }
        catch {
            Write-Error "Go program execution failed: $_"
            Set-Location ..
            return $false
        }
    }
    else {
        Write-Error "Go command not found, cannot generate JSON files"
        return $false
    }
}

# 6. Clean temporary files
function Remove-TempFiles {
    Write-Info "Cleaning temporary files..."
    if (Test-Path "temp") {
        Remove-Item -Recurse -Force "temp"
    }
    Write-Info "Cleanup completed"
}

# Main function
function Start-Build {
    Write-Info "Starting I18n data automated build..."
    
    # Check necessary commands
    $requiredCommands = @("wget", "curl") # PowerShell has built-in Invoke-WebRequest
    if (!(Test-Command "wget") -and !(Test-Command "curl")) {
        Write-Warn "Recommend installing wget or curl for better download experience"
    }
    
    # Create directories
    New-Directories
    
    # Execute download steps
    if (!(Get-IanaTzdata)) { return }
    if (!(Get-TimeZoneDB)) { return }
    if (!(Get-CldrDtd)) { return }
    if (!(Get-WindowsZones)) { return }
    
    # Generate JSON
    New-JsonFiles
    
    # Cleanup
    Remove-TempFiles
    
    Write-Info "I18n data automated build completed!"
}

# Clean function
function Start-Clean {
    Write-Info "Cleaning all generated files..."
    @("tzdb", "data/*.csv", "windows/*.dtd", "windows/*.xml", "windows/*.json", "temp") | ForEach-Object {
        if (Test-Path $_) {
            Remove-Item -Recurse -Force $_ -ErrorAction SilentlyContinue
        }
    }
    Write-Info "Cleanup completed"
}

# Help information
function Show-Help {
    Write-Host "Usage: .\build_i18n.ps1 [options]"
    Write-Host "Options:"
    Write-Host "  -Clean    Clean all generated files"
    Write-Host "  -Help     Show this help information"
    Write-Host "  (none)    Execute complete automated build"
}

# Main entry point
if ($Help) {
    Show-Help
}
elseif ($Clean) {
    Start-Clean
}
else {
    Start-Build
}