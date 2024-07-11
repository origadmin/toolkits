#!/bin/bash

# Get the module name as a parameter
MODULE_NAME=$1

# Function to find the latest tag matching the module name
find_latest_tag() {
    current_commit=$(git rev-parse HEAD)
    module_name=$1
    latest_tag=""

    # Get all tags that point to the current commit
    tags_for_current_commit=$(git tag --points-at "$current_commit")

    # If a module name is provided, filter out matching tags
    if [ -n "$module_name" ]; then
        # Convert the tag list into an array
        IFS=$'\n' read -d '' -r -a tags_array <<< "$tags_for_current_commit"

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
        done <<< "$tags_for_current_commit"
    fi

    # Output the latest tag
    echo "$latest_tag"
}

# Call the function
LATEST_TAG=$(find_latest_tag "$MODULE_NAME")
echo "$LATEST_TAG"