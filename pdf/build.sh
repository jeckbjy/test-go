#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )
cd $DIR

go build -o $DIR/pdf $DIR/main.go
   