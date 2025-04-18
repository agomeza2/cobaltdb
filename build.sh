#!/bin/bash

# Create the build directory if it doesn't exist
mkdir -p build

cd src || exit 
# Build the project
go build -o ../build/cobalt . 

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build succeeded. Executable is in the 'build' directory."
else
    echo "Build failed."
fi
