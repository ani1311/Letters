#!/bin/bash

# Set the name of the binary to be built
BINARY_NAME="main"

# Build the binary
go build -o $BINARY_NAME

# Run the binary
./$BINARY_NAME

# Cleanup the binary
rm $BINARY_NAME
