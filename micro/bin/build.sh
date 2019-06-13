#!/usr/bin/env bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )
ROOT=$DIR/..
BUILD=$DIR/build/
export GOPATH=$GOPATH:$ROOT

cd $ROOT/src

build() {
	for d in $(ls ./$1); do
		echo "building $1/$d"
		pushd ./$1/$d >/dev/null
		CGO_ENABLED=0 GOOS=linux go build -o $BUILD/$d -a -installsuffix cgo -ldflags '-w'
		popd >/dev/null
	done
}

mkdir -p $BUILD
build srv
build api
