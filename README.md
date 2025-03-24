# Go Coverage Demo

## Generating Coverage Reports for SonarCloud

This project includes scripts to generate test coverage reports compatible with SonarCloud.

### Prerequisites

Before running the coverage script, make sure you have installed the required tools:

```bash
go install github.com/axw/gocov/gocov@latest
go install github.com/AlekSi/gocov-xml@latest
```

### Running the Coverage Script

#### On Windows

To generate the coverage reports on Windows, run:

```cmd
scripts\generate-coverage.bat
```

#### On Linux/Mac

To generate the coverage reports on Linux or Mac, run:

```bash
# Make the script executable (only needed once)
chmod +x scripts/generate-coverage.sh

# Run the script
./scripts/generate-coverage.sh
```

This will:
1. Run all tests in the project with coverage enabled
2. Generate HTML coverage report for local inspection
3. Convert the Go coverage output to XML format compatible with SonarCloud
4. Place all reports in the `coverage` directory

### Using with SonarCloud

The generated XML file at `coverage/coverage.xml` can be used with SonarCloud. The `sonar-project.properties` file is already configured to use this file.

When running SonarCloud analysis, it will automatically pick up the coverage report.
