#!/usr/bin/env bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )
ROOT=$DIR/../
set -e

if [[ $# -lt 1 ]]; then
    echo "need service folder"
    exit 1
fi

BASE_DIR="srv"
if [[ $# -eq 2 ]]; then
    BASE_DIR=$2
fi

PROTO_DIR=$ROOT/src/$BASE_DIR/$1/proto

if [[ ! -d $PROTO_DIR ]]; then
    echo "cannot find path:$PROTO_DIR"
    exit 1
fi

cd $PROTO_DIR
echo "build proto:$1"
protoc --proto_path=../../.. --proto_path=. *.proto --micro_out=. --go_out=.

#workdir=`pwd`
#echo "build proto in dir:$workdir"
#protoc *.proto --micro_out=. --go_out=.
