#!/usr/bin/env bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )

pushd $DIR/server
go build -o $DIR/bin/server .
popd

pushd $DIR/client
go build -o $DIR/bin/client .
popd