#!/bin/bash

# Get the current version from the git tag
VERSION=$(git describe --tags $(git rev-list --tags --max-count=1))
# Get the number of commits as build number
BUILD=$(git rev-list --count HEAD)
# Get the short version of the current git commit hash (optional)
GITHASH=$(git rev-parse --short HEAD)

# Ensure the build directory exists
mkdir -p server/build

# Build the application with the version and build number
go build -o server/build/educado.out -ldflags "-X main.Version=$VERSION -X main.Build=$BUILD -X main.GitHash=$GITHASH"

# Copy the .env file to the build directory
cp ../../.env server/build/
cp ../../gcp_credentials.json server/build/


echo "Built version $VERSION (Build $BUILD) and copied .env to server/build/"
