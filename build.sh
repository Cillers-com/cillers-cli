#!/bin/bash

# Name of your app
APP_NAME="cillers"

# Supported platforms
PLATFORMS=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
    "windows/amd64"
)

# Version number
VERSION=$(git describe --tags --always --dirty)

# Build directory
BUILD_DIR="build"

# Create build directory if it doesn't exist
mkdir -p $BUILD_DIR

# Loop through platforms
for PLATFORM in "${PLATFORMS[@]}"; do
    # Split platform into OS and architecture
    IFS='/' read -r -a PLATFORM_SPLIT <<< "$PLATFORM"
    GOOS=${PLATFORM_SPLIT[0]}
    GOARCH=${PLATFORM_SPLIT[1]}
    
    # Set output binary name
    if [ $GOOS = "windows" ]; then
        OUTPUT_NAME=$APP_NAME'_'$GOOS'_'$GOARCH'_'$VERSION'.exe'
    else
        OUTPUT_NAME=$APP_NAME'_'$GOOS'_'$GOARCH'_'$VERSION
    fi
    
    # Build binary
    echo "Building for $GOOS $GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $BUILD_DIR/$OUTPUT_NAME
    
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done

echo "Build process completed!"
