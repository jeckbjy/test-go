#!/usr/bin/env sh
set -e
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )
ROOT=$DIR/..

cd $DIR

# Install proto3 from source
#  brew install autoconf automake libtool
#  git clone https://github.com/google/protobuf
#  ./autogen.sh ; ./configure ; make ; make install
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples
rm -f *.go

echo 'build grpc'
protoc greeter.proto --go_out=plugins=grpc:.

echo 'copy to go-kit/pb'
GOKIT=$ROOT/go-kit/pb/
mkdir -p $GOKIT
cp *.pb.go $GOKIT

echo 'copy to gizmo/pb'
GIZMOS=$ROOT/gizmo/pb/
mkdir -p $GIZMOS
cp *.pb.go $GIZMOS

# https://github.com/micro/protoc-gen-micro
echo 'build micro'
rm -f *.pb.go
protoc greeter.proto --micro_out=. --go_out=.
MICRO=${ROOT}/go-micro/pb/
mkdir -p ${MICRO}
cp *.go ${MICRO}
rm -f *.go