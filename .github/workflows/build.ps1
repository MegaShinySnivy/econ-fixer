# Determine binary name based on OS
if ($env:RUNNER_OS -eq 'Windows') {
    $BINARY = "econfixer.exe"
} elseif ($env:RUNNER_OS -eq 'macOS') {
    $BINARY = "econfixer-darwin"
} elseif ($env:RUNNER_OS -eq 'Linux') {
    $BINARY = "econfixer-linux"
} else {
    Write-Host "Unsupported OS: $env:RUNNER_OS" -ForegroundColor Red
    exit 1
}

# Build and move binary with platform-specific name
go build -o "$BINARY" .
Move-Item -Path "$BINARY" -Destination "econfixer" -Force