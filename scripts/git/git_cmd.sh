#!/bin/bash

# Helper function to get tags matching a pattern
git_tags_matching_pattern() {
  local pattern=$1
  git tag -l "$pattern"
}

# Helper function to get the latest tag from a list of tags
git_latest_tag() {
  local tags=$1
  # Sort the tags and pick the last one as the latest
  echo "$tags" | sort -V | tail -n1
}

git_latest_commit_hash() {
  git rev-parse HEAD
}

git_tags_on_commit_hash(){
  local hash=$1
  git tag --points-at "$hash"
}

git_add_tags_on_commit_hash(){
  local next_tag=$1
  git tag -a "$next_tag" -m "Bumped version to $next_tag"
}