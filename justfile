
PGPING_VERSION := `git tag --sort=taggerdate | tail -1`

dev:
	air

build:
	go build -o ./build/pg-ping -ldflags "-s -w -X main.version={{PGPING_VERSION}}-dev" .

run:
	./build/pg-ping

start: build run

build-dist:
	go build -o ./build/pg-ping -ldflags "-s -w -X main.version={{PGPING_VERSION}}" .
