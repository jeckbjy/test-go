#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )

BUILD_TIME=`date +"%s"`
go build -ldflags "-X main.BuildTime=$BUILD_TIME" -o ../bin/inject