#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )
HOST=mywork
DEST=/var/www/dht/

# build
GOOS=linux GOARCH=amd64 go build -o spider .

# copy
rsync -avcuR --progress ./spider ${HOST}:${DEST}