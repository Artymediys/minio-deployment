#!/bin/bash

# Define architectures
ARCHS=("amd64" "arm64")

# Output directory
OUT_DIR="./"

# Compile for each architecture
for ARCH in ${ARCHS[@]}; do
    echo "Building for $ARCH..."
    GOOS=linux GOARCH=$ARCH go build -o $OUT_DIR/minio-deployment-$ARCH
done

echo "Builds completed."
