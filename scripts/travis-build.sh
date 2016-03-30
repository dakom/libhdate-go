#!/bin/sh
echo "running tests and building..."
go build ./ ./github.com/dakom/libhdate-go/libhdate
go test ./ ./github.com/dakom/libhdate-go/tests
