#!/bin/bash

source "$(pwd)"/scripts/git/git_tag.sh
source "$(pwd)"/scripts/git/git_cmd.sh

# Record the original working directory
ORIGINAL_DIR=$(pwd)

# Function to check for the existence of a go.mod file in a directory
find_go_mod() {
    local dir="$1"
    local go_mod_name="go.mod"
    local updated=1  # Assume not updated by default

    # Change to the directory
    cd "$dir" || return 1

    # Check if the go.mod file exists in the directory
    if [ -f "$go_mod_name" ]; then
        echo "Find dependencies go.mod in directories:"
        echo " ->DIR: $dir"
        updated=0
    fi

    # Return to the original working directory for the next iteration
    cd "$ORIGINAL_DIR" || return 1

    # Return the update status
    return $updated
}

# Function to check if required functions are defined
function_checks() {
      if ! declare -f git_latest_tag > /dev/null; then
          echo "Error: git_latest_tag function is not defined"
          exit 1
      fi

      if ! declare -f git_tags_matching_pattern > /dev/null; then
         echo "Error: git_tags_matching_pattern function is not defined"
         exit 1
      fi

     # Check if get_latest_commit_hash is defined
     if ! declare -f get_head_version_tag > /dev/null; then
         echo "Error: get_head_version_tag function is not defined"
         exit 1
     fi

     if ! declare -f get_next_version_tag > /dev/null; then
         echo "Error: get_next_version_tag function is not defined"
         exit 1
     fi

    if ! declare -f get_head_version_tag > /dev/null; then
       echo "Error: get_head_version_tag function is not defined"
       exit 1
    fi
}

# Function to check the go.mod file and perform actions
check_go_mod_and_act() {
   local dir="$1"
   local module="main"
   local module_name  # Drop the beginning './'
   module_name="$(echo "$dir" | sed "s/\/$//")"

   if [ "$module_name" != "." ]; then
     module="$module_name"
   fi

   if find_go_mod "$dir"; then
        local HEAD_TAG
        local LATEST_TAG
        local NEXT_TAG
        echo " ->MODULE_NAME: $module"
        HEAD_TAG=$(get_head_version_tag "$module_name")
        echo " ->HEAD_TAG: $HEAD_TAG"
        LATEST_TAG=$(get_latest_version_tag "$module_name")
        echo " ->LATEST_TAG: $LATEST_TAG"
        NEXT_TAG=$(get_next_version_tag "$module_name")
        echo " ->NEXT_TAG: $NEXT_TAG"
        echo ""
   fi

}

# Define a function to traverse directories and apply the check_go_mod_and_act function
add_go_mod_tag() {
    echo "Checking for go.mod files..."
    function_checks
    local module_name="$1"

    echo ""
    echo "COMMIT_HASH: $(git_latest_commit_hash "$module_name")"
    echo ""

    # If a module_name is specified, process only that directory
    if [ "$module_name" == "." ]; then
       check_go_mod_and_act "."
    elif [ -n "$module_name" ]; then
      check_go_mod_and_act "$module_name"
    else
        # Skip the root directory ('.')
        find . -mindepth 1 -type d -not -path './.*' -print0 | while IFS= read -r -d '' dir; do
            # Construct the commit message, including the module name
            # the module name must be the directory name without the beginning './'
            # otherwise, the commit message will be incorrect
            check_go_mod_and_act "$dir"
        done
    fi
}

add_go_mod_tag "$1"
