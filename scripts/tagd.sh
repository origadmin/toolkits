#!/bin/bash

# Record the original working directory
ORIGINAL_DIR=$(pwd)
MODULE_NAME=$1


# Define a function to check for the existence of a go.mod file in a directory and perform corresponding actions
check_go_mod() {
    local dir="$1"
    local go_mod_name="go.mod"
    local updated=1  # Assume not updated by default

    # Change to the directory
    cd "$dir" || return 1

    # Check if the go.mod file exists in the directory
    if [ -f "$go_mod_name" ]; then
        echo "Checking directories $dir"
        # Only mark as updated if tests pass
        updated=0
    fi

    # Return to the original working directory
    cd "$ORIGINAL_DIR" || return 1
    # Return the update status
    return $updated
}

# Define a function to traverse directories and apply the check_go_mod_and_act function
tag_go_mod() {
    local module_name="$1"

    # If a module_name is specified, process only that directory
    if [ -n "$module_name" ]; then
        if check_go_mod "./$module_name"; then
            echo "MODULE_NAME: $module_name"
            HEAD_TAG=$("$ORIGINAL_DIR"/scripts/git/head_tag.sh "$module_name")
            HEAD_TAG=$("$ORIGINAL_DIR"/scripts/git/head_tag.sh "$module_name")
            echo "HEAD_TAG: $HEAD_TAG"
            NEXT_TAG=$("$ORIGINAL_DIR"/scripts/git/next_tag.sh "$module_name")
            echo "NEXT_TAG: $NEXT_TAG"
        fi
    else
        # Skip the root directory ('.')
        find . -mindepth 1 -type d -not -path './.*' -print0 | while IFS= read -r -d '' dir; do
            if check_go_mod "$dir"; then
              # Construct the commit message, including the module name
              # the module name must be the directory name without the beginning './'
              # otherwise, the commit message will be incorrect
              module_name=$(echo "$dir" | sed "s/^.\///")  # Drop the beginning './'

              echo "MODULE_NAME: $module_name"
              HEAD_TAG=$("$ORIGINAL_DIR"/scripts/git/head_tag.sh "$module_name")
              echo "HEAD_TAG: $HEAD_TAG"
              NEXT_TAG=$("$ORIGINAL_DIR"/scripts/git/next_tag.sh "$module_name")
              echo "NEXT_TAG: $NEXT_TAG"

              if [ -z "$HEAD_TAG" ]; then
                  git tag -a "$NEXT_TAG" -m "Bumped version to $NEXT_TAG"
              fi

            fi
        done
    fi
}

tag_go_mod "$MODULE_NAME"
