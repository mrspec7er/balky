#!/bin/bash

# Run Go tests
go test ./test/health_test.go

# Check the exit code of the tests
if [ $? -eq 0 ]; then
  exit 0  # Tests passed, exit with success
else
  exit 1  # Tests failed, exit with failure
fi