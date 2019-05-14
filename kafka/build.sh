#!/usr/bin/env bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )

go build -o zapi     ./cmd/api/main.go
go build -o zworker  ./cmd/worker/main.go