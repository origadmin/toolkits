#!/bin/bash

# Gets the module name as a parameter
MODULE_NAME=$1

# Helper function to get tags matching a pattern
get_tags_matching_pattern() {
    local pattern=$1
    git tag -l "$pattern"
}

# Helper function to get the latest tag from a list of tags
get_latest_tag() {
    local tags=$1
    # Sort the tags and pick the last one as the latest
    echo "$tags" | sort -V | tail -n1
}

# Function to get the next version tag based on existing tags
get_next_version_tag() {
    local module_name=$1

    # Determine the pattern based on whether module_name is provided
    local pattern="${module_name:+$module_name/}v*"

    # Get all tags that match the pattern
    local tags=$(get_tags_matching_pattern "$pattern")

    # Get the latest tag
    local latest_tag=$(get_latest_tag "$tags")

    if [[ -z "$latest_tag" ]]; then
        # Default to v0.0.1
        local next_tag="v0.0.1"
    else
        # Correctly extract the version part from the tag
        local version_part=${latest_tag##"$module_name/"} # Changed to use ## instead of #
        IFS='.' read -r -a version_array <<< "${version_part#v}" # Also remove leading 'v'
        local major=${version_array[0]}
        local minor=${version_array[1]}
        local patch=${version_array[2]}
        ((patch++))
        local next_tag="v$major.$minor.$patch"
    fi

    # If module_name is provided, prepend it to the tag
    if [[ -n "$module_name" ]]; then
        next_tag="$module_name/$next_tag"
    fi

    echo "$next_tag"
}

# Call the function with the module name
NEXT_TAG=$(get_next_version_tag "$MODULE_NAME")
echo "$NEXT_TAG"
