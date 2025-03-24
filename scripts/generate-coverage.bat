@echo off
setlocal

REM Create a directory for coverage reports if it doesn't exist
if not exist coverage mkdir coverage

REM Run tests with coverage
go test ./... -coverprofile=coverage/coverage.out

REM Convert Go coverage to HTML for local visualization (optional)
go tool cover -html=coverage/coverage.out -o coverage/coverage.html

REM Install gocov and gocov-xml if not already installed
go install github.com/axw/gocov/gocov@latest
go install github.com/AlekSi/gocov-xml@latest

REM Get GOPATH
FOR /F "tokens=*" %%g IN ('go env GOPATH') do (SET GOPATH=%%g)

REM Convert Go coverage to Cobertura XML format (SonarCloud compatible)
"%GOPATH%\bin\gocov" convert coverage/coverage.out | "%GOPATH%\bin\gocov-xml" > coverage/coverage.xml

echo Coverage report generated at coverage/coverage.xml
echo This file can be used with SonarCloud for coverage reporting

endlocal
