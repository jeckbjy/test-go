BUILD_TIME=`date +"%s"`
go build -ldflags "-X main.BuildTime=$BUILD_TIME" -o ../bin/inject