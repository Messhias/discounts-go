#!/bin/sh
go test ./tests/... -coverprofile=test-results/coverage.out
go tool cover -html=test-results/coverage.out -o test-results/coverage.html
go test ./tests/... -v > test-results/test_results.txt
./dgoo