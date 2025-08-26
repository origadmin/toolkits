#!/bin/bash
#
# Copyright (c) 2024 OrigAdmin. All rights reserved.
#
# This script combines testing, dependency updating, and git tagging functionalities
# for Go modules within the repository.

set -e
set -o pipefail

# --- Script Variables ---
ORIGINAL_DIR=$(pwd)

# --- Git Helper Functions ---

# get_matching_tags finds tags matching a given pattern.
get_matching_tags() {
    local pattern=$1
    git tag -l "$pattern"
}

# get_latest_tag finds the latest tag from a list of version tags.
get_latest_tag() {
    local tags=$1
    sort -V <<<"$tags" | tail -n1
}

# get_current_commit_hash returns the short hash of the current commit.
get_current_commit_hash() {
    git rev-parse --short HEAD
}

# get_tags_for_commit returns tags pointing to a specific commit hash.
get_tags_for_commit() {
    local hash=$1
    git tag --points-at "$hash"
}

# create_new_tag creates a new annotated git tag.
create_new_tag() {
    local tag_name=$1
    local message="Bumped version to $tag_name"
    git tag -a "$tag_name" -m "$message"
    echo "Created new tag: $tag_name"
}

# get_next_module_version determines the next semantic version for a module.
get_next_module_version() {
    local module_path=$1
    local pattern="v*"
    [[ "$module_path" != "." ]] && pattern="${module_path}/v*"

    local tags
    tags=$(get_matching_tags "$pattern")
    
    local latest_tag
    latest_tag=$(get_latest_tag "$tags")

    local next_tag
    if [[ -z "$latest_tag" ]]; then
        next_tag="v0.1.0"
    else
        local version_part=${latest_tag##*/}
        IFS='.' read -r -a version_array <<<"${version_part#v}"
        local major=${version_array[0]:-0}
        local minor=${version_array[1]:-0}
        local patch=${version_array[2]:-0}
        ((patch++))
        next_tag="v$major.$minor.$patch"
    fi

    if [[ "$module_path" != "." ]]; then
        next_tag="$module_path/$next_tag"
    fi
    echo "$next_tag"
}

# get_head_version_tag checks if the current HEAD is already tagged.
get_head_version_tag() {
    local module_path=$1
    local commit_hash
    commit_hash=$(get_current_commit_hash)
    
    local tags_on_commit
    tags_on_commit=$(get_tags_for_commit "$commit_hash")
    
    local latest_matching_tag=""
    while IFS= read -r tag; do
        local pattern="v[0-9]+\.[0-9]+\.[0-9]+"
        if [[ "$module_path" != "." ]]; then
            pattern="^${module_path}/${pattern}$"
        else
            # Match root tags like v1.2.3, not module/v1.2.3
            if [[ "$tag" == */* ]]; then
                continue
            fi
            pattern="^${pattern}$"
        fi

        if [[ $tag =~ $pattern ]]; then
            if [[ -z "$latest_matching_tag" ]] || [[ "$tag" > "$latest_matching_tag" ]]; then
                latest_matching_tag=$tag
            fi
        fi
    done <<<"$tags_on_commit"

    echo "$latest_matching_tag"
}


# --- Core Logic Functions ---

# run_tests executes tests for a given module directory.
run_tests() {
    local module_dir=$1
    echo "--- Testing module: $module_dir ---"
    cd "$module_dir"
    go fmt ./...
    go mod tidy
    go vet ./...
    go test ./...
    cd "$ORIGINAL_DIR"
    echo "--- Tests passed for: $module_dir ---"
}

# update_deps updates dependencies for a given module directory.
update_deps() {
    local module_dir=$1
    echo "--- Updating dependencies for module: $module_dir ---"
    cd "$module_dir"
    go get -u all
    go mod tidy
    
    local module_name=${module_dir#./}
    if git status --porcelain | grep -q -E "go.(mod|sum)"; then
        git add go.mod go.sum
        local commit_message="feat($module_name): Update go.mod and go.sum"
        echo "Committing changes with message: $commit_message"
        git commit -m "$commit_message"
    else
        echo "No dependency changes to commit in $module_name"
    fi

    cd "$ORIGINAL_DIR"
}

# create_tags creates a new version tag for a module if it's not already tagged.
create_tags() {
    local module_dir=$1
    local module_name=${module_dir#./}
    
    echo "--- Checking tags for module: $module_name ---"

    local head_tag
    head_tag=$(get_head_version_tag "$module_name")

    if [[ -n "$head_tag" ]]; then
        echo "HEAD is already tagged as $head_tag for module $module_name. Nothing to do."
        return
    fi

    local next_tag
    next_tag=$(get_next_module_version "$module_name")
    
    echo "Next tag will be: $next_tag"
    create_new_tag "$next_tag"
}

# --- Main Logic ---

# traverse_and_execute finds all Go modules and runs a specified command on them.
traverse_and_execute() {
    local command_func=$1
    local target_dir=${2:-.} # Default to current directory if not specified

    if [[ -f "${target_dir}/go.mod" ]]; then
        "$command_func" "$target_dir"
        return
    fi

    find "$target_dir" -name "go.mod" -print0 | while IFS= read -r -d '' go_mod_file; do
        local module_dir
        module_dir=$(dirname "$go_mod_file")
        # Skip vendor directories and hidden directories
        if [[ "$module_dir" == *"/vendor/"* ]] || [[ "$module_dir" == *"/.*"* ]]; then
            continue
        fi
        "$command_func" "$module_dir"
    done
}

main() {
    if [[ $# -eq 0 ]]; then
        echo "Error: No command specified."
        echo "Usage: $0 {test|update|tag} [target_directory]"
        exit 1
    fi

    local command=$1
    local target_dir=${2:-.}

    # Switch to main branch for tagging to ensure consistency
    local current_branch=""
    if [[ "$command" == "tag" ]]; then
        current_branch=$(git branch --show-current)
        echo "Stashing changes and switching to main branch for tagging..."
        git stash save --include-untracked "ci-script-stash"
        git checkout main
    fi

    case "$command" in
        test)
            traverse_and_execute run_tests "$target_dir"
            ;;
        update)
            traverse_and_execute update_deps "$target_dir"
            ;;
        tag)
            traverse_and_execute create_tags "$target_dir"
            echo "Pushing new tags..."
            git push origin --tags
            ;;
        *)
            echo "Error: Unknown command '$command'"
            echo "Usage: $0 {test|update|tag} [target_directory]"
            # Restore branch if we switched
            if [[ -n "$current_branch" ]]; then
                git checkout "$current_branch"
                git stash pop || true
            fi
            exit 1
            ;;
    esac

    # Restore original branch and unstash changes if we switched
    if [[ -n "$current_branch" ]]; then
        echo "Switching back to $current_branch and restoring stashed changes..."
        git checkout "$current_branch"
        git stash pop || true
    fi

    echo "--- Command '$command' completed successfully. ---"
}

# Execute the main function with all script arguments
main "$@"
