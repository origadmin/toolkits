#!/bin/bash

# Use git describe to find the latest tag
NEXT_TAG=$(git describe --tags $(git rev-list --tags --max-count=1))

if [[ -z "$NEXT_TAG" ]]; then
  NEXT_TAG="v0.0.1"
else
  # Remove 'v' prefix using bash parameter expansion
  NEXT_TAG=${NEXT_TAG#v}
  IFS='.' read -r -a VERSION_ARRAY <<< "$NEXT_TAG"
  MAJOR=${VERSION_ARRAY[0]}
  MINOR=${VERSION_ARRAY[1]}
  PATCH=${VERSION_ARRAY[2]}
  (( PATCH++ ))
  NEXT_TAG=v$MAJOR.$MINOR.$PATCH
  # Output the latest tag for use in subsequent steps
  echo "$NEXT_TAG"
fi

