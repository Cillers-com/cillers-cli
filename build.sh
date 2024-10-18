#!/bin/bash

# Set the project directory
PROJECT_DIR="$(pwd)"

# Set the output file name
OUTPUT_FILE="cillers"

# Build the Go project
echo "Building Go project..."
go build -o "$OUTPUT_FILE" "$PROJECT_DIR/main.go"

if [ $? -eq 0 ]; then
    echo "Build successful. Output file: $OUTPUT_FILE"
else
    echo "Build failed."
    exit 1
fi

# Install the binary to /usr/local/bin
echo "Installing $OUTPUT_FILE to /usr/local/bin..."
sudo mv "$OUTPUT_FILE" /usr/local/bin/

if [ $? -eq 0 ]; then
    echo "Installation successful. You can now run '$OUTPUT_FILE' from anywhere."
else
    echo "Installation failed. You may need to run this script with sudo privileges."
    exit 1
fi
