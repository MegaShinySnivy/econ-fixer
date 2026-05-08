#!/bin/bash

# Create builds directory if it doesn't exist
mkdir -p builds

# Get the project name from go.mod (defaults to econfixer)
PROJECT_NAME=$(grep "^module " go.mod | cut -d' ' -f2)
if [ -z "$PROJECT_NAME" ]; then
    PROJECT_NAME="econfixer"
fi

echo "Compiling ${PROJECT_NAME}..."

# Build for Linux (default GOOS, produces native binary)
echo "Building for Linux..."
go build -o builds/${PROJECT_NAME}.bin .

# Build for macOS
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o builds/${PROJECT_NAME}-mac .

# Build for Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o builds/${PROJECT_NAME}-win64.exe .

echo "Build complete. Output files in ./builds/"