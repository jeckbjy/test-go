#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )
cd $DIR

curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.2.0 sh -