.DEFAULT_GOAL := build 

clean:
	rm -f tableweaver
	rm -f tableweaver.linux
	rm -f tw-cli
	rm -f tw-cli.linux

build: clean
	templ generate
	go build -o tableweaver cmd/web/*.go
	go build -o tw-cli cmd/cli/*.go

linuxbuild: clean
	templ generate
	GOARCH=amd64 GOOS=linux go build -o tableweaver.linux cmd/web/*.go
	GOARCH=amd64 GOOS=linux go build -o tw-cli.linux cmd/cli/*.go

run:
	go run cmd/web/*.go