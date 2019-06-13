.PHONY: proto build

proto:
	./proto.sh hotel api
	./proto.sh auth
	./proto.sh geo
	./proto.sh profile
	./proto.sh rate

build:
	./build.sh

run:
	docker-compose build
	docker-compose up
