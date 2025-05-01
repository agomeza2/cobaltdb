# Create the build directory if it doesn't exist
$buildDir = "build"
if (-Not (Test-Path $buildDir)) {
    New-Item -ItemType Directory -Path $buildDir
}

# Change to the src directory
Set-Location -Path "src"

# Build the project
$buildOutput = "../$buildDir/cobalt.exe"
$buildResult = go build -o $buildOutput .

# Check if the build was successful
if ($LASTEXITCODE -eq 0) {
    Write-Host "Build succeeded. Executable is in the 'build' directory."
} else {
    Write-Host "Build failed."
}
