#!/bin/bash

source "$(pwd)"/scripts/git/git_cmd.sh

function_checks() {
     # Check if git_latest_commit_hash is defined
     if ! declare -f git_latest_commit_hash > /dev/null; then
         echo "Error: git_latest_commit_hash function is not defined"
         exit 1
     fi

     if ! declare -f git_latest_tag > /dev/null; then
         echo "Error: git_latest_tag function is not defined"
         exit 1
     fi

     if ! declare -f git_tags_matching_pattern > /dev/null; then
         echo "Error: git_tags_matching_pattern function is not defined"
         exit 1
     fi
}

# Function to get the next version tag based on existing tags
get_next_version_tag() {
    function_checks

    local module_name=$1

    # Determine the pattern based on whether module_name is provided
    local pattern="v*"
    if [ "$module_name" != "." ]; then
      pattern="${module_name}/v*"
    fi

    # Get all tags that match the pattern
    local tags=""
    tags=$(git_tags_matching_pattern "$pattern")

    # Get the latest tag
    local latest_tag=""
    latest_tag=$(git_latest_tag "$tags")

    local next_tag
    if [[ -z "$latest_tag" ]]; then
        # Default to v0.0.1
        next_tag="v0.0.1"
    else
        # Correctly extract the version part from the tag
        local version_part=${latest_tag##"$module_name/"} # Changed to use ## instead of #
        IFS='.' read -r -a version_array <<< "${version_part#v}" # Also remove leading 'v'
        local major=${version_array[0]}
        local minor=${version_array[1]}
        local patch=${version_array[2]}
        ((patch++))
        next_tag="v$major.$minor.$patch"
    fi

    # If module_name is provided, prepend it to the tag
    if [[ "$module_name" != "." ]]; then
        next_tag="$module_name/$next_tag"
    fi

    echo "$next_tag"
}


# Function to find the latest tag matching the module name
get_head_version_tag() {
    function_checks

    module_name=$1
    latest_tag=""

    # Get the current commit hash
    local commit_hash
    commit_hash=$(git_latest_commit_hash)

   # Determine the pattern based on whether module_name is provided
#    local pattern="v*"
#    if [ "$module_name" != "." ]; then
#      pattern="${module_name}/v*"
#    fi

    # Get all tags that match the pattern
#    local tags=""
#    tags=$(git_tags_matching_pattern "$pattern")

    # Get the latest tag
#    local latest_tag=""
#    latest_tag=$(git_latest_tag "$tags")
#
#    # If the latest tag is newer than the current commit, update latest_tag
#    if [ -z "$latest_tag" ]; then
#      latest_tag="$commit_hash"
#    fi
    # Get all tags that point to the current commit
    tags_on_commit=$(git_tags_on_commit_hash "$commit_hash")

    if [ "$module_name" == "." ]; then
      # Convert the tag list into an array
        IFS=$'\n' read -d '' -r -a tags_array <<< "$tags_on_commit"

        # Iterate over all tags, looking for tags that match the module name
        for tag in "${tags_array[@]}"; do
            if [[ "$tag" == "$module_name"* ]]; then
                # If it's the first matching tag or newer than the existing tag, update latest_tag
                if [[ -z "$latest_tag" ]] || [[ "$tag" > "$latest_tag" ]]; then
                    latest_tag=$tag
                fi
            fi
        done
    # If a module name is provided, filter out matching tags
    elif [ -n "$module_name" ]; then
        # Convert the tag list into an array
        IFS=$'\n' read -d '' -r -a tags_array <<< "$tags_on_commit"

        # Iterate over all tags, looking for tags that match the module name
        for tag in "${tags_array[@]}"; do
            if [[ "$tag" == "$module_name"* ]]; then
                # If it's the first matching tag or newer than the existing tag, update latest_tag
                if [[ -z "$latest_tag" ]] || [[ "$tag" > "$latest_tag" ]]; then
                    latest_tag=$tag
                fi
            fi
        done
    else
        # Read tags line by line
        while IFS= read -r tag; do
            # Check if the tag matches the vx.y.z format
            if [[ $tag =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
                # If it's the first matching tag or newer than the existing tag, update latest_tag
                if [[ -z "$latest_tag" ]] || [[ "$tag" > "$latest_tag" ]]; then
                    latest_tag=$tag
                fi
            fi
        done <<< "$tags_on_commit"
    fi

    # Output the latest tag
    echo "$latest_tag"
}

get_latest_version_tag() {
  local module_name=$1

  # Determine the pattern based on whether module_name is provided
  local pattern="v*"
  if [ "$module_name" != "." ]; then
    pattern="${module_name}/v*"
  fi

  local tags=""
  tags=$(git_tags_matching_pattern "$pattern")

  # Get the latest tag
  local latest_tag=""
  latest_tag=$(git_latest_tag "$tags")
  echo "$latest_tag"
}

# Call the function with the module name
# git_next_version_tag "$MODULE_NAME"
