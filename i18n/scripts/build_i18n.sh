#!/bin/bash

# I18n Data Download and Build Script
# Automates all steps from README.md

set -e

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if command exists
check_command() {
    if ! command -v $1 &> /dev/null; then
        log_error "$1 command not found, please install it first"
        exit 1
    fi
}

# Create directories
create_directories() {
    log_info "Creating necessary directories..."
    mkdir -p tzdb
    mkdir -p data
    mkdir -p windows
    mkdir -p temp
}

# Download file with retry
download_with_retry() {
    local url=$1
    local output=$2
    local max_retries=3
    local retry_count=0
    
    while [ $retry_count -lt $max_retries ]; do
        if wget --timeout=30 --tries=1 -O "$output" "$url"; then
            log_info "Download successful: $output"
            return 0
        else
            retry_count=$((retry_count + 1))
            log_warn "Download failed, retry $retry_count/$max_retries: $url"
            sleep 2
        fi
    done
    
    log_error "Download failed, max retries reached: $url"
    return 1
}

# 1. Download tzdata from IANA
download_iana_tzdata() {
    log_info "Starting IANA tzdata download..."
    
    cd tzdb
    
    # Download tzcode
    log_info "Downloading tzcode-latest.tar.gz..."
    download_with_retry "https://www.iana.org/time-zones/repository/tzcode-latest.tar.gz" "tzcode-latest.tar.gz"
    
    # Download tzdata
    log_info "Downloading tzdata-latest.tar.gz..."
    download_with_retry "https://www.iana.org/time-zones/repository/tzdata-latest.tar.gz" "tzdata-latest.tar.gz"
    
    # Extract
    log_info "Extracting tzcode..."
    gzip -dc tzcode-latest.tar.gz | tar -xf -
    
    log_info "Extracting tzdata..."
    gzip -dc tzdata-latest.tar.gz | tar -xf -
    
    cd ..
    log_info "IANA tzdata download and extraction completed"
}

# 2. Download data from TimeZoneDB
download_timezonedb() {
    log_info "Starting TimeZoneDB data download..."
    
    cd temp
    download_with_retry "https://timezonedb.com/files/TimeZoneDB.csv.zip" "TimeZoneDB.csv.zip"
    
    # Extract
    log_info "Extracting TimeZoneDB.csv.zip..."
    unzip -o TimeZoneDB.csv.zip
    
    # Move to data directory
    if [ -f "time_zone.csv" ]; then
        mv time_zone.csv ../data/
        log_info "time_zone.csv moved to data directory"
    fi
    
    cd ..
    log_info "TimeZoneDB data download completed"
}

# 3. Download CLDR DTD
download_cldr_dtd() {
    log_info "Starting CLDR DTD download..."
    
    download_with_retry "https://raw.githubusercontent.com/unicode-org/cldr/main/common/dtd/ldmlSupplemental.dtd" "windows/ldmlSupplemental.dtd"
    
    log_info "CLDR DTD download completed"
}

# 4. Download Windows timezone XML
download_windows_zones() {
    log_info "Starting Windows timezone data download..."
    
    download_with_retry "https://raw.githubusercontent.com/unicode-org/cldr/main/common/supplemental/windowsZones.xml" "windows/windowsZones.xml"
    
    log_info "Windows timezone data download completed"
}

# 5. Generate JSON files
generate_json_files() {
    log_info "Starting JSON file generation..."
    
    # Call Go program
    if command -v go &> /dev/null; then
        cd tz
        go run make.go -generate-json
        cd ..
        log_info "JSON file generation completed"
    else
        log_error "Go command not found, cannot generate JSON files"
        return 1
    fi
}

# 6. Clean temporary files
cleanup() {
    log_info "Cleaning temporary files..."
    rm -rf temp
    log_info "Cleanup completed"
}

# Main function
main() {
    log_info "Starting I18n data automated build..."
    
    # Check necessary commands
    check_command wget
    check_command gzip
    check_command tar
    check_command unzip
    
    # Create directories
    create_directories
    
    # Execute download steps
    download_iana_tzdata
    download_timezonedb
    download_cldr_dtd
    download_windows_zones
    
    # Generate JSON
    generate_json_files
    
    # Cleanup
    cleanup
    
    log_info "I18n data automated build completed!"
}

# Handle command line arguments
case "${1:-}" in
    "clean")
        log_info "Cleaning all generated files..."
        rm -rf tzdb data/*.csv windows/*.dtd windows/*.xml windows/*.json
        log_info "Cleanup completed"
        ;;
    "help"|"-h"|"--help")
        echo "Usage: $0 [options]"
        echo "Options:"
        echo "  clean    Clean all generated files"
        echo "  help     Show this help message"
        echo "  (none)   Execute complete automated build"
        ;;
    *)
        main
        ;;
esac