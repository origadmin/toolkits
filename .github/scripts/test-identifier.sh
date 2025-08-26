#!/bin/bash
#
# This script runs tests for all modules under the 'identifier' directory.

set -e # Exit immediately if a command exits with a non-zero status.
set -o pipefail # The pipeline's return status is the value of the last command to exit with a non-zero status.

# The root of the toolkits project, which is two levels up from this script's directory.
TOOLKITS_ROOT=$(cd "$(dirname "$0")/../.." && pwd)

# The path to the new ci.sh script.
CI_SCRIPT="$TOOLKITS_ROOT/.github/scripts/ci.sh"

# The target directory to test.
IDENTIFIER_DIR="$TOOLKITS_ROOT/identifier"

# Check if the identifier directory exists
if [ ! -d "$IDENTIFIER_DIR" ]; then
    echo "Error: The identifier directory was not found at $IDENTIFIER_DIR"
    exit 1
fi

# Make sure the ci.sh script is executable
chmod +x "$CI_SCRIPT"

# Run the tests for the identifier directory
echo "--- Running tests for all identifier modules ---"
"$CI_SCRIPT" test "$IDENTIFIER_DIR"

echo "--- Identifier tests completed successfully ---"
