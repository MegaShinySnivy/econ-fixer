#!/bin/bash

# Determine binary name based on OS
if [ "${{ runner.os }}" = "macOS" ]; then
  BINARY="econfixer-darwin"
elif [ "${{ runner.os }}" = "Linux" ]; then
  BINARY="econfixer-linux"
else
  echo "Unsupported OS: ${{ runner.os }}"
  exit 1
fi

# Build and move binary with platform-specific name
go build -o "$BINARY" .
mv "$BINARY" econfixer