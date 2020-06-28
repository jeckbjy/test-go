#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )
HOST=mywork
DEST=/var/www/torsniff/

# build
GOOS=linux GOARCH=amd64 go build -o torsniff .

# copy
rsync -avcuR --progress ./torsniff ${HOST}:${DEST}