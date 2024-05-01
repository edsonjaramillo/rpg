#!/bin/bash

# The name of the package. This is used in the name of the output file.
package_name="rpg"

# The directory where the output files will be placed.
dist_dir="dist"

# An array of the platforms we want to build for. Each platform is a string in the format OS/ARCH.
platforms=("darwin/amd64" "linux/386" "linux/amd64" "windows/386" "windows/amd64")

# Loop over each platform
for platform in "${platforms[@]}"
do
    # Split the platform into OS and ARCH
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    # Construct the output name using the package name and the platform
    output_name=$package_name'-'$GOOS'-'$GOARCH

    echo 'Building for '$GOOS'-'$GOARCH

    # If the OS is Windows, add a .exe extension to the output name
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi    

    # Build the program for the specified platform and put the output in the dist directory
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $dist_dir/$output_name cmd/rpg/main.go

    # If the build failed, print an error message and exit the script
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done