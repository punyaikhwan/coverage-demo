#!/bin/bash
set -e

# Create a directory for coverage reports if it doesn't exist
mkdir -p coverage

# Run tests with coverage
go test ./... -coverprofile=coverage/coverage.out

# Convert Go coverage to HTML for local visualization (optional)
go tool cover -html=coverage/coverage.out -o coverage/coverage.html

# Install gocov and gocov-xml if not already installed
go install github.com/axw/gocov/gocov@latest
go install github.com/AlekSi/gocov-xml@latest

# Convert Go coverage to Cobertura XML format (SonarCloud compatible)
gocov convert coverage/coverage.out | gocov-xml > coverage/coverage.xml

echo "Coverage report generated at coverage/coverage.xml"
echo "This file can be used with SonarCloud for coverage reporting"
